package review

import (
	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/incoming"
	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/outgoing"
	reviewEnt "github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/review"
	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/repository/review"
	"github.com/rs/zerolog"
	"time"
)

type Service interface {
	GetReviewsForEntity(request incoming.GetReviewsRequest) (outgoing.GetReviewsResponse, error)
	AddReviewForEntity(request incoming.AddReviewRequest) (outgoing.AddReviewResponse, error)
}

type impl struct {
	reviewRepository review.Repository
	logger           *zerolog.Logger
}

func NewService(
	reviewRepository review.Repository,
	logger *zerolog.Logger,
) (*impl, error) {
	return &impl{
		reviewRepository: reviewRepository,
		logger:           logger,
	}, nil
}

func (i *impl) GetReviewsForEntity(
	request incoming.GetReviewsRequest,
) (outgoing.GetReviewsResponse, error) {
	summary, reviews, err := i.reviewRepository.GetByEntityID(request.ReviewedEntityID, request.PreloadedReviewsLimit)
	if err != nil {
		return outgoing.GetReviewsResponse{}, err
	}

	return outgoing.GetReviewsResponse{
		Summary: summary,
		Reviews: reviews,
	}, nil
}

func (i *impl) AddReviewForEntity(
	request incoming.AddReviewRequest,
) (outgoing.AddReviewResponse, error) {
	newReview := i.doSomeWork(reviewEnt.Review{
		AuthorID: request.AuthorID,
		Text:     request.Text,
		Stars:    request.Stars,
		Created:  time.Now(),
		Payload:  [10]uint64{1},
	}, 600)

	// fmt.Printf(" %d\n", unsafe.Sizeof(newReview))
	// fmt.Printf(" %v\n", newReview)

	err := i.reviewRepository.Add(request.ReviewedEntityID, newReview)
	if err != nil {
		return outgoing.AddReviewResponse{}, err
	}

	return outgoing.AddReviewResponse{}, nil
}

func (i *impl) doSomeWork(in reviewEnt.Review, repeat int) reviewEnt.Review {
	if repeat == 0 {
		return in
	}

	a := in
	a.Stars += 1

	return i.doSomeWork(a, repeat-1)
}
