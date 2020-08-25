package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest the lambda request handler
func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name, ok := req.QueryStringParameters["name"]
	if !ok {
		res := events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}
		return res, nil
	}

	var oath string
	switch name {
	case "Aragorn":
		oath = "Take my sword"
	case "Legolas":
		oath = "And my bow"
	case "Gimli":
		oath = "And my axe"
	}

	res := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "text/plain; charset=utf-8"},
		Body:       oath,
	}
	return res, nil
}

func main() {
	lambda.Start(HandleRequest)
}
