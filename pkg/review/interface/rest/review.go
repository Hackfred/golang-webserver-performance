package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"

	v1 "github.com/Hackfred/golang-webserver-performance/openapi"
	"github.com/Hackfred/golang-webserver-performance/pkg/review/interface/rest/hydration"
)

func (r *restApi) GetReviews(ctx echo.Context, entityId string, params v1.GetReviewsParams) error {
	domainRequest, err := hydration.HydrateGetReviewsRequest(entityId, params)
	if err != nil {
		return err
	}

	domainResponse, err := r.reviewAPI.GetReviews(domainRequest)
	if err != nil {
		return err
	}

	response, err := hydration.HydrateGetReviewsResponse(domainResponse)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}

func (r *restApi) AddReviews(ctx echo.Context, entityId string) error {
	params := v1.AddReviewRequest{}
	err := ctx.Bind(&params)
	if err != nil {
		return err
	}

	domainRequest, err := hydration.HydrateAddReviewRequest(entityId, params)
	if err != nil {
		return err
	}

	domainResponse, err := r.reviewAPI.AddReview(domainRequest)
	if err != nil {
		return err
	}

	response, err := hydration.HydrateAddReviewResponse(domainResponse)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}
