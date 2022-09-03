package rest

import (
	"github.com/Hackfred/golang-webserver-performance/pkg/review/application"
)

type restApi struct {
	reviewAPI application.Application
}

func NewRestApi(
	reviewAPI application.Application,
) (*restApi, error) {
	return &restApi{
		reviewAPI: reviewAPI,
	}, nil
}
