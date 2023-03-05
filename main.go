package main

import (
	"awesomeProject/handlers"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ApiResponse := events.APIGatewayProxyResponse{}

	switch request.HTTPMethod {
	case "GET":
		return handlers.GetExercises(), nil
	case "POST":
		return handlers.CreateExercise(request), nil
	default:
		ApiResponse.StatusCode = 404
		return ApiResponse, nil
	}
}

func main() {
	lambda.Start(HandleRequest)
}
