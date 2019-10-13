package testSupport

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
	"strconv"
)

func SetupTestDynamoDb() (*dynamodb.DynamoDB, func()) {
	endpoint := os.Getenv("DYNAMODB_ENDPOINT")
	if endpoint == "" {
		endpoint = "http://localhost:3307"
	}
	dynamoSession, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("hogehoge", "fugafuga", ""),
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String(endpoint),
	})
	if err != nil {
		panic(err.Error())
	}
	dynamoDbHandler := dynamodb.New(dynamoSession)

	return dynamoDbHandler, func() {
		deletePostInput := &dynamodb.DeleteItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"Id": {
					N: aws.String("2"),
				},
			},
			TableName: aws.String("Posts"),
		}
		_, _ = dynamoDbHandler.DeleteItem(deletePostInput)

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

		updateInput = &dynamodb.UpdateItemInput{
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":m": {
					N: aws.String(strconv.Itoa(1)),
				},
			},
			TableName: aws.String("IdCounter"),
			Key: map[string]*dynamodb.AttributeValue{
				"Identifier": {
					S: aws.String("Users"),
				},
			},
			UpdateExpression: aws.String("set MaxId = :m"),
		}
		_, _ = dynamoDbHandler.UpdateItem(updateInput)

		userDeleteInput := &dynamodb.DeleteItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"Email": {
					S: aws.String("test@example.com"),
				},
			},
			TableName: aws.String("Users"),
		}
		_, _ = dynamoDbHandler.DeleteItem(userDeleteInput)

		userUpdateInput := &dynamodb.UpdateItemInput{
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":a": {
					S: aws.String("Q1LZlKFt2h0000001vER4TjyFo7"),
				},
			},
			TableName: aws.String("Users"),
			Key: map[string]*dynamodb.AttributeValue{
				"Email": {
					S: aws.String("hoge@example.com"),
				},
			},
			UpdateExpression: aws.String("set AuthenticationToken = :a"),
		}
		_, _ = dynamoDbHandler.UpdateItem(userUpdateInput)
	}
}
