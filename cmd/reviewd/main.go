package main

import (
	"net"
	"os"
	"runtime/trace"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	grpcV1 "github.com/Hackfred/golang-webserver-performance/grpc/review/v1"
	restV1 "github.com/Hackfred/golang-webserver-performance/openapi"
	"github.com/Hackfred/golang-webserver-performance/pkg/review/application"
	reviewSrv "github.com/Hackfred/golang-webserver-performance/pkg/review/domain/service/review"
	reviewRepo "github.com/Hackfred/golang-webserver-performance/pkg/review/infrastructure/persistence/memory/review"
	grpcApi "github.com/Hackfred/golang-webserver-performance/pkg/review/interface/grpc"
	"github.com/Hackfred/golang-webserver-performance/pkg/review/interface/rest"
)

func main() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	f, err := os.Create("trace.out")
	if err != nil {
		logger.Fatal().Msgf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			logger.Fatal().Msgf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		logger.Fatal().Msgf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	shutdownWg := sync.WaitGroup{}
	shutdownWg.Add(1)

	reviewRepository, _ := reviewRepo.NewRepository()
	reviewService, _ := reviewSrv.NewService(reviewRepository, &logger)
	reviewApi, _ := application.NewApplication(reviewService)

	startGrpcServer(reviewApi, "localhost:8080")
	startRestServer(reviewApi, "localhost:8081")

	go func() {
		time.Sleep(time.Second * 10)
		shutdownWg.Done()
	}()

	shutdownWg.Wait()
}

func startGrpcServer(reviewApi application.Application, address string) {
	grpcServer := grpc.NewServer()

	reviewServiceServer, err := grpcApi.NewGrpcApi(reviewApi)
	if err != nil {
		os.Exit(1)
	}

	go func() {
		grpcV1.RegisterReviewServiceServer(grpcServer, reviewServiceServer)
		listener, err := net.Listen("tcp4", address)
		err = grpcServer.Serve(listener)
		if err != nil {
			os.Exit(1)
		}
	}()
}

func startRestServer(reviewApi application.Application, address string) {
	restServer := echo.New()

	reviewServiceServer, err := rest.NewRestApi(reviewApi)
	if err != nil {
		os.Exit(1)
	}

	// restServer.Use(middleware.Logger())

	go func() {
		restV1.RegisterHandlers(restServer, reviewServiceServer)
		restServer.Logger.Fatal(restServer.Start(address))
	}()
}
