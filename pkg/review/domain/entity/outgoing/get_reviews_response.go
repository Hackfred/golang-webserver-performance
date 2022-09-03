package outgoing

import (
	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/review"
	reviewVal "github.com/Hackfred/golang-webserver-performance/pkg/review/domain/value/review"
)

type GetReviewsResponse struct {
	Summary *reviewVal.Summary
	Reviews []review.Review
}
