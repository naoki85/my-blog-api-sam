package testSupport

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
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
