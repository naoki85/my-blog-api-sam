package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/naoki85/my-blog-api-sam/model"
	"log"
	"os"
)

type IdCounterRepository struct {
	DynamoDBHandler *dynamodb.DynamoDB
}

func (repo *IdCounterRepository) tableName() (tableName string) {
	if tableName = os.Getenv("ID_COUNTER_TABLE_NAME"); tableName != "" {
		return tableName
	} else {
		return "IdCounter"
	}
}

func (repo *IdCounterRepository) TestShouldFindCountByIdentifier(identifier string) (count int, err error) {
	result, err := repo.DynamoDBHandler.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repo.tableName()),
		Key: map[string]*dynamodb.AttributeValue{
			"Identifier": {
				S: aws.String(identifier),
			},
		},
	})
	if err != nil {
		log.Printf("dynamodbErr: %s", err.Error())
		return
	}

	item := model.IdCounter{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		log.Printf("Could not unmarshal map: %s", err.Error())
		return
	}

	return item.MaxId, nil
}
