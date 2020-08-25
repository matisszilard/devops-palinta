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
		oath = "If by my life or death I can protect you I will You have my sword."
	case "Legolas":
		oath = "And you have my bow"
	case "Gimli":
		oath = "And my axe"
	case "Boromir":
		oath = "You carry the fates of us all little one If this is indeed the will of the Council then Gondor will see it done."
	case "Gandalf":
		oath = "I will help you bear this burden Frodo Baggins as long as it is yours to bear."
	case "Thanos":
		oath = "Oops! Wrong universe!"
	default:
		oath = "I don't know who you are. I don't know what you want. If you are looking for the CI/CD. I can tell you I don't have it, but what I do have are a very particular set of skills. Skills I have acquired over a very long career (1 month). Skills that make me a nightmare for people like you. If you let the devops rule now that'll be the end of it. I will not look for you, I will not pursue you, but if you don't, I will look for you, I will find you and I will devopsify you."
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
