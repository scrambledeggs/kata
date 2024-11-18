package main

import (
	"net/http"

	"foobar/internal/helpers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	r "github.com/scrambledeggs/booky-go-common/apigatewayresponse"
	"github.com/scrambledeggs/booky-go-common/logs"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	paths := []string{
		"/v1/cached-world",
	}

	out, err := helpers.InvalidateCache(paths)
	if err != nil {
		logs.Error("InvalidateCache", err.Error())
		return r.SingleErrorResponse(http.StatusInternalServerError, r.ErrorResponseBody{
			Message: err.Error(),
			Code:    "INTERNAL_SERVER_ERROR",
		})
	}

	logs.Print("Successfully invalidated cache", out)
	return r.SingleSuccessResponse(http.StatusOK, "ok")
}

func main() {
	lambda.Start(handler)
}
