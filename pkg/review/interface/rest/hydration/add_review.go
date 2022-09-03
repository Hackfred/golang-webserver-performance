package hydration

import (
	v1 "github.com/Hackfred/golang-webserver-performance/openapi"
	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/incoming"
	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/outgoing"
	"github.com/Hackfred/golang-webserver-performance/pkg/review/domain/entity/review"
)

func HydrateAddReviewRequest(entityId string, in v1.AddReviewRequest) (incoming.AddReviewRequest, error) {
	return incoming.AddReviewRequest{
		Stars:            in.Stars,
		ReviewedEntityID: review.ReviewedEntityID(entityId),
		AuthorID:         in.AuthorId,
		Text:             in.Text,
	}, nil
}

func HydrateAddReviewResponse(in outgoing.AddReviewResponse) (v1.AddReviewResponse, error) {
	return v1.AddReviewResponse{}, nil
}
