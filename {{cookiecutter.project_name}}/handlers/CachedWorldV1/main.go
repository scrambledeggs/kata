package main

import (
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/scrambledeggs/booky-go-common/logs"

	r "github.com/scrambledeggs/booky-go-common/apigatewayresponse"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fileData, err := os.ReadFile("./cache_thingy.txt") // Replace with actual file path
	if err != nil {
		return r.SingleErrorResponse(http.StatusInternalServerError, r.ErrorResponseBody{
			Message: "Failed to read file: " + err.Error(),
			Code:    "INTERNAL_SERVER_ERROR",
		})
	}

	greeting := string(fileData)

	logs.Print("Greet", greeting)
	return r.SingleSuccessResponse(http.StatusOK, greeting)
}

func main() {
	lambda.Start(handler)
}
