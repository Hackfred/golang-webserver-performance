package review

import (
	"sync"
	"time"

	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/review"
	reviewVal "github.com/Hackfred/golang-webserver-performance/pkg/review/domain/value/review"
)

type impl struct {
	reviews     map[review.ReviewedEntityID][]review.Review
	reviewsLock sync.Mutex
}

func NewRepository() (*impl, error) {
	return &impl{
		reviews:     map[review.ReviewedEntityID][]review.Review{},
		reviewsLock: sync.Mutex{},
	}, nil
}

func (i *impl) Add(entityID review.ReviewedEntityID, reviewToAdd review.Review) error {
	time.Sleep(time.Millisecond * 100)
	i.reviewsLock.Lock()
	if existingReviews, exists := i.reviews[entityID]; exists {
		i.reviews[entityID] = append(existingReviews, reviewToAdd)
	} else {
		i.reviews[entityID] = []review.Review{reviewToAdd}
	}
	i.reviewsLock.Unlock()

	return nil
}

func (i *impl) GetByEntityID(entityID review.ReviewedEntityID, preloadReviews uint32) (*reviewVal.Summary, []review.Review, error) {
	summary := &reviewVal.Summary{
		Count:        0,
		AverageStars: 0,
	}

	var reviews []review.Review
	existingReviews, exists := i.reviews[entityID]
	if exists {
		sum := uint32(0)
		for _, r := range existingReviews {
			sum += uint32(r.Stars)
		}

		summary.Count = uint32(len(existingReviews))
		summary.AverageStars = float32(sum) / float32(summary.Count)

		limit := int(preloadReviews)
		for c, r := range existingReviews {
			if c >= limit {
				break
			}

			reviews = append(reviews, r)
		}
	}

	return summary, reviews, nil
}
