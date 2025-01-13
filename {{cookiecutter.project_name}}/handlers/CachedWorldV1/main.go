package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/scrambledeggs/booky-go-common/logs"

	r "github.com/scrambledeggs/booky-go-common/apigatewayresponse"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	randomNumber := rand.Intn(100) + 1
	randomNumberStr := strconv.Itoa(randomNumber)

	greeting := "Random number saved: " + randomNumberStr
	logs.Print("Greet", greeting)
	return r.SingleSuccessResponse(http.StatusOK, greeting)
}

func main() {
	lambda.Start(handler)
}
