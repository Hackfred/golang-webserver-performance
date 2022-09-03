package grpc

import (
	v1 "github.com/Hackfred/golang-webserver-performance/grpc/review/v1"
	"github.com/Hackfred/golang-webserver-performance/pkg/review/application"
)

type grpcApi struct {
	v1.UnimplementedReviewServiceServer
	reviewAPI application.Application
}

func NewGrpcApi(
	reviewAPI application.Application,
) (*grpcApi, error) {
	return &grpcApi{
		reviewAPI: reviewAPI,
	}, nil
}
