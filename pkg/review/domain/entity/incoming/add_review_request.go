package incoming

import "github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/review"

type AddReviewRequest struct {
	Stars            uint8
	ReviewedEntityID review.ReviewedEntityID
	AuthorID         string
	Text             string
}
