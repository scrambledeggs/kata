package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/scrambledeggs/booky-go-common/logs"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP

	if sourceIP == "" {
		greeting = "Hello, world!\n"
	} else {
		greeting = fmt.Sprintf("Hello, %s! In %s env\n", sourceIP, os.Getenv("APP_ENV"))
	}

	logs.Print("Greet", greeting)

	return events.APIGatewayProxyResponse{
		Body:       greeting,
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(handler)
}
