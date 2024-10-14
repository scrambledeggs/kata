package main

import (
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	gw "github.com/scrambledeggs/booky-go-common/apigatewayresponse"
	"github.com/scrambledeggs/booky-go-common/logs"
	"gopkg.in/yaml.v3"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var endpoints []string

	rawEndpoints, err := os.ReadFile("./endpoints.yml")
	if err != nil {
		logs.Error("ReadFile", err.Error())
		return gw.SingleErrorResponse(http.StatusInternalServerError, gw.ErrorResponseBody{Code: "API_ERROR", Message: err.Error()})
	}

	err = yaml.Unmarshal([]byte(rawEndpoints), &endpoints)
	if err != nil {
		logs.Error("Unmarshal Endpoints", err.Error())
		return gw.SingleErrorResponse(http.StatusInternalServerError, gw.ErrorResponseBody{Code: "API_ERROR", Message: err.Error()})
	}

	return gw.SingleSuccessResponse(http.StatusOK, endpoints)
}

func main() {
	lambda.Start(handler)
}
