package hydration

import (
	v1 "github.com/Hackfred/golang-webserver-performance/grpc/review/v1"
	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/incoming"
	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/outgoing"
	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/review"
)

const dateFormat = "2006-02-01"

func HydrateGetReviewsRequest(in *v1.GetReviewsRequest) (incoming.GetReviewsRequest, error) {
	return incoming.GetReviewsRequest{
		ReviewedEntityID:      review.ReviewedEntityID(in.ReviewedEntityId),
		PreloadedReviewsLimit: in.PreloadedReviewsLimit,
	}, nil
}

func HydrateGetReviewsResponse(in outgoing.GetReviewsResponse) (v1.GetReviewsResponse, error) {
	out := v1.GetReviewsResponse{}

	if in.Summary != nil {
		out.ReviewSummary = &v1.ReviewSummary{
			Count:        in.Summary.Count,
			AverageStars: in.Summary.AverageStars,
		}
	}

	if in.Reviews != nil {
		out.Reviews = make([]*v1.Review, len(in.Reviews))
		for c, r := range in.Reviews {
			out.Reviews[c] = hydrateReview(r)
		}
	}

	return out, nil
}

func hydrateReview(in review.Review) *v1.Review {
	out := &v1.Review{
		Id:       in.ID,
		AuthorId: in.AuthorID,
		Text:     in.AuthorID,
		Stars:    uint32(in.Stars),
		Created:  in.Created.Format(dateFormat),
	}

	if in.Updated != nil {
		out.Updated = in.Updated.Format(dateFormat)
	}

	return out
}
