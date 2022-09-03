package incoming

import "github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/review"

type GetReviewsRequest struct {
	ReviewedEntityID      review.ReviewedEntityID
	PreloadedReviewsLimit uint32
}
