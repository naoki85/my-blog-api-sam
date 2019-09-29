package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/naoki85/my-blog-api-sam/config"
	"log"
)

type DynamoDbHandler struct {
	Conn *dynamodb.DynamoDB
}

func NewDynamoDbHandler(c *config.Config) (*DynamoDbHandler, error) {
	dynamoSession, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("hogehoge", "fugafuga", ""),
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String(c.DynamoDbEndpoint),
	})
	if err != nil {
		log.Printf("%s", err.Error())
	}

	DynamoDbHandler := new(DynamoDbHandler)
	DynamoDbHandler.Conn = dynamodb.New(dynamoSession)
	return DynamoDbHandler, nil
}
