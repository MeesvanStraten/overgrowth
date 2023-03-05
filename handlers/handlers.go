package handlers

import (
	"awesomeProject/structs"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"log"
)

const tableName = "Exercises"

func CreateExercise(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse) {
	ApiResponse := events.APIGatewayProxyResponse{}
	svc := GetDBClient()
	var exercise structs.Exercise

	if err := json.Unmarshal([]byte(request.Body), &exercise); err != nil {
		ApiResponse.StatusCode = 500
		return ApiResponse
	}

	av, err := dynamodbattribute.MarshalMap(exercise)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
	ApiResponse.StatusCode = 201
	return ApiResponse
}

func GetDBClient() (dyna *dynamodb.DynamoDB) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return dynamodb.New(sess)
}

func GetExercises() (response events.APIGatewayProxyResponse) {
	ApiResponse := events.APIGatewayProxyResponse{}

	svc := GetDBClient()

	proj := expression.NamesList(expression.Name("Title"), expression.Name("Description"))

	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
		ApiResponse.StatusCode = 500
	}

	ApiResponse.StatusCode = 200
	ApiResponse.Body = result.GoString()
	return ApiResponse
}
