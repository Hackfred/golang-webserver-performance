package grpc

import (
	"context"
	v1 "github.com/Hackfred/golang-webserver-performance/grpc/review/v1"
	"github.com/Hackfred/golang-webserver-performance/pkg/review/interface/grpc/hydration"
)

func (s *grpcApi) GetReviews(_ context.Context, request *v1.GetReviewsRequest) (*v1.GetReviewsResponse, error) {
	domainRequest, err := hydration.HydrateGetReviewsRequest(request)

	domainResponse, err := s.reviewAPI.GetReviews(domainRequest)
	if err != nil {
		return nil, err
	}

	response, err := hydration.HydrateGetReviewsResponse(domainResponse)

	return &response, nil
}

func (s *grpcApi) AddReview(_ context.Context, request *v1.AddReviewRequest) (*v1.AddReviewResponse, error) {
	domainRequest, err := hydration.HydrateAddReviewRequest(request)

	domainResponse, err := s.reviewAPI.AddReview(domainRequest)
	if err != nil {
		return nil, err
	}

	response, err := hydration.HydrateAddReviewResponse(domainResponse)

	return &response, nil
}
