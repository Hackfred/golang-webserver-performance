package application

import (
	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/incoming"
	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/outgoing"
	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/service/review"
)

type Application interface {
	GetReviews(request incoming.GetReviewsRequest) (outgoing.GetReviewsResponse, error)
	AddReview(request incoming.AddReviewRequest) (outgoing.AddReviewResponse, error)
}

type impl struct {
	reviewService review.Service
}

func NewApplication(
	reviewService review.Service,
) (*impl, error) {
	return &impl{
		reviewService: reviewService,
	}, nil
}

func (i *impl) GetReviews(request incoming.GetReviewsRequest) (outgoing.GetReviewsResponse, error) {
	return i.reviewService.GetReviewsForEntity(request)
}

func (i *impl) AddReview(request incoming.AddReviewRequest) (outgoing.AddReviewResponse, error) {
	return i.reviewService.AddReviewForEntity(request)
}
