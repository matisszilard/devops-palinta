package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	l "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// HandleRequest the lambda request handler
func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name, ok := req.QueryStringParameters["name"]
	if !ok {
		return sendBadRequest()
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambda.New(sess, &aws.Config{Region: aws.String("eu-central-1")})

	heroRequest := events.APIGatewayProxyRequest{}

	heroRequest.QueryStringParameters = make(map[string]string)

	heroRequest.QueryStringParameters["name"] = name
	payload, err := json.Marshal(heroRequest)
	if err != nil {
		fmt.Println("Error marshalling MyGetItemsFunction request")
		return sendBadRequest()
	}

	result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String("mszg-gondol-hero"), Payload: payload})
	if err != nil {
		fmt.Println("Error calling MyGetItemsFunction")
		return sendBadRequest()
	}

	var resp events.APIGatewayProxyResponse
	err = json.Unmarshal(result.Payload, &resp)
	if err != nil {
		fmt.Println("Error unmarshalling MyGetItemsFunction response")
		return sendBadRequest()
	}

	fmt.Printf("Response: %+v\n", resp)

	// If the status code is NOT 200, the call failed
	if resp.StatusCode != 200 {
		fmt.Println("Error getting items, StatusCode: " + strconv.Itoa(resp.StatusCode))
		return sendBadRequest()
	}

	res := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":                 "text/plain; charset=utf-8",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
			"Access-Control-Allow-Methods": "OPTIONS,GET"},
		Body: fmt.Sprintf("I am %s! This is my oath: %s\n", name, resp.Body),
	}
	return res, nil
}

func sendBadRequest() (events.APIGatewayProxyResponse, error) {
	res := events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
	}
	return res, nil
}

func main() {
	l.Start(HandleRequest)
}
