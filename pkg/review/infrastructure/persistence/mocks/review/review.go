package review

import (
	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/review"
	reviewVal "github.com/Hackfred/golang-webserver-performance/pkg/review/domain/value/review"
	"time"
)

type impl struct {
}

func NewRepository() (*impl, error) {
	return &impl{}, nil
}

func (i *impl) Add(entityID review.ReviewedEntityID, review *review.Review) error {
	time.Sleep(time.Millisecond * 100)

	return nil
}

func (i *impl) GetByEntityID(entityID review.ReviewedEntityID, preloadReviews uint32) (*reviewVal.Summary, []*review.Review, error) {
	time.Sleep(time.Millisecond * 100)

	summary := &reviewVal.Summary{
		Count:        0,
		AverageStars: 0,
	}

	reviews := []*review.Review{
		{},
	}

	return summary, reviews, nil
}
