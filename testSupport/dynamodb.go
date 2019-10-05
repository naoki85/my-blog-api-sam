package testSupport

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"strconv"
)

func NewDynamoDbHandler() (*dynamodb.DynamoDB, error) {
	dynamoSession, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("hogehoge", "fugafuga", ""),
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String("http://localhost:3307"),
	})
	if err != nil {
		log.Printf("%s", err.Error())
	}
	return dynamodb.New(dynamoSession), err
}

func SetupTestDynamoDb() (*dynamodb.DynamoDB, func()) {
	dynamoSession, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("hogehoge", "fugafuga", ""),
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String("http://localhost:3307"),
	})
	if err != nil {
		panic(err.Error())
	}
	dynamoDbHandler := dynamodb.New(dynamoSession)

	return dynamoDbHandler, func() {
		deleteInput := &dynamodb.DeleteItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"Id": {
					N: aws.String("6"),
				},
			},
			TableName: aws.String("RecommendedBooks"),
		}
		_, _ = dynamoDbHandler.DeleteItem(deleteInput)

		updateInput := &dynamodb.UpdateItemInput{
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":m": {
					N: aws.String(strconv.Itoa(5)),
				},
			},
			TableName: aws.String("IdCounter"),
			Key: map[string]*dynamodb.AttributeValue{
				"Identifier": {
					S: aws.String("RecommendedBooks"),
				},
			},
			UpdateExpression: aws.String("set MaxId = :m"),
		}
		_, _ = dynamoDbHandler.UpdateItem(updateInput)
	}
}
