package review

import (
	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/review"
	reviewVal "github.com/Hackfred/golang-webserver-performance/pkg/review/domain/value/review"
)

type Repository interface {
	Add(entityID review.ReviewedEntityID, review review.Review) error
	GetByEntityID(entityID review.ReviewedEntityID, preloadReviews uint32) (*reviewVal.Summary, []review.Review, error)
}
