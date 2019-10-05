package repository

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/naoki85/my-blog-api-sam/model"
	"log"
	"os"
	"time"
)

type UserRepository struct {
	DynamoDBHandler *dynamodb.DynamoDB
}

type UserCreateParams struct {
	Id       int
	Email    string
	Password string
}

func (repo *UserRepository) tableName() (tableName string) {
	if tableName = os.Getenv("USERS_TABLE_NAME"); tableName != "" {
		return tableName
	} else {
		return "Users"
	}
}

func (repo *UserRepository) FindByEmail(value string) (user model.User, err error) {
	result, err := repo.DynamoDBHandler.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repo.tableName()),
		Key: map[string]*dynamodb.AttributeValue{
			"Email": {
				S: aws.String(value),
			},
		},
	})
	if err != nil {
		log.Printf("dynamodbErr: %s", err.Error())
		return
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		log.Printf("Could not unmarshal map: %s", err.Error())
		return
	}
	return user, err
}

func (repo *UserRepository) FindByAuthenticationToken(value string) (user model.User, err error) {
	result, err := repo.DynamoDBHandler.Query(&dynamodb.QueryInput{
		TableName:              aws.String(repo.tableName()),
		IndexName:              aws.String("AuthenticationTokenIndex"),
		KeyConditionExpression: aws.String("AuthenticationToken = :token"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":token": {
				S: aws.String(value),
			},
		},
	})
	if err != nil {
		log.Printf("dynamodbErr: %s", err.Error())
		return
	}
	if len(result.Items) != 1 {
		log.Printf("invalid response from dynamo db: %d", len(result.Items))
		return user, errors.New("invalid response from dynamodb")
	}

	err = dynamodbattribute.UnmarshalMap(result.Items[0], &user)
	if err != nil {
		log.Printf("Could not unmarshal map: %s", err.Error())
		return
	}
	return user, err
}

func (repo *UserRepository) UpdateAttribute(email string, field string, param string) (bool, error) {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#f": aws.String(field),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":m": {
				S: aws.String(param),
			},
		},
		TableName: aws.String(repo.tableName()),
		Key: map[string]*dynamodb.AttributeValue{
			"Email": {
				S: aws.String(email),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set #f = :m"),
	}
	_, err := repo.DynamoDBHandler.UpdateItem(input)
	if err != nil {
		log.Printf("dynamodbErr: %s", err.Error())
		return false, err
	}

	return true, nil
}

func (repo *UserRepository) Create(params UserCreateParams) (err error) {
	type Item struct {
		Id                int
		Email             string
		EncryptedPassword string
		CreatedAt         string
		UpdatedAt         string
	}
	now := time.Now().Format("2006-01-02 15-04-05")

	item := Item{
		Id:                params.Id,
		Email:             params.Email,
		EncryptedPassword: params.Password,
		CreatedAt:         now,
		UpdatedAt:         now,
	}
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Println("Got error marshalling new user item:")
		log.Println(err.Error())
		return
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(repo.tableName()),
	}

	_, err = repo.DynamoDBHandler.PutItem(input)
	if err != nil {
		log.Println("Got error calling PutItem:")
		log.Println(err.Error())
		return
	}

	return
}
