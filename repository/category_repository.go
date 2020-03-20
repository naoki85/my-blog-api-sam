package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/naoki85/my-blog-api-sam/model"
	"log"
	"os"
	"time"
)

type CategoryRepository struct {
	DynamoDBHandler *dynamodb.DynamoDB
}

func (repo *CategoryRepository) tableName() (tableName string) {
	if tableName = os.Getenv("CATEGORIES_TABLE_NAME"); tableName != "" {
		return tableName
	} else {
		return "Categories"
	}
}

type CategoryCreateParams struct {
	Identifier string
	JpName     string
	Color      string
}

func (repo *CategoryRepository) All() (categories model.Categories, err error) {
	result, err := repo.DynamoDBHandler.Scan(&dynamodb.ScanInput{
		TableName:            aws.String(repo.tableName()),
		ProjectionExpression: aws.String("Identifier, JpName, Color"),
	})

	if err != nil {
		log.Printf("dynamodbErr: %s", err.Error())
		return
	}

	for _, i := range result.Items {
		c := model.Category{}
		err = dynamodbattribute.UnmarshalMap(i, &c)
		if err != nil {
			log.Println("Got error unmarshalling:")
			log.Println(err.Error())
			return
		}
		categories = append(categories, c)
	}

	return
}

func (repo *CategoryRepository) FindByIdentifier(identifier string) (category model.Category, err error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Identifier": {
				S: aws.String(identifier),
			},
		},
		TableName: aws.String(repo.tableName()),
	}

	result, err := repo.DynamoDBHandler.GetItem(input)
	if err != nil {
		log.Printf("dynamodbErr: %s", err.Error())
		return
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &category)
	if err != nil {
		log.Printf("Could not unmarshal map: %s", err.Error())
		return
	}
	return
}

type CategoryItem struct {
	Identifier string
	JpName     string
	Color      string
	CreatedAt  string
	UpdatedAt  string
}

func (repo *CategoryRepository) Create(params CategoryCreateParams) (err error) {
	now := time.Now().Format("2006-01-02 15-04-05")

	item := CategoryItem{
		Identifier: params.Identifier,
		JpName:     params.JpName,
		Color:      params.Color,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Println("Got error marshalling new movie item:")
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

	return nil
}

func (repo *CategoryRepository) Update(params CategoryCreateParams) (err error) {
	now := time.Now().Format("2006-01-02 15-04-05")

	item := CategoryItem{
		Identifier: params.Identifier,
		JpName:     params.JpName,
		Color:      params.Color,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	stmt := "set JpName = :name, Color = :color, UpdatedAt = :updatedAt"

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name": {
				S: aws.String(item.JpName),
			},
			":color": {
				S: aws.String(item.Color),
			},
			":updatedAt": {
				S: aws.String(item.UpdatedAt),
			},
		},
		TableName: aws.String(repo.tableName()),
		Key: map[string]*dynamodb.AttributeValue{
			"Identifier": {
				S: aws.String(item.Identifier),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String(stmt),
	}

	_, err = repo.DynamoDBHandler.UpdateItem(input)
	if err != nil {
		log.Printf("dynamodbErr: %s", err.Error())
	}

	return
}

func (repo *CategoryRepository) Delete(identifier string) (err error) {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(repo.tableName()),
		Key: map[string]*dynamodb.AttributeValue{
			"Identifier": {
				S: aws.String(identifier),
			},
		},
	}

	_, err = repo.DynamoDBHandler.DeleteItem(input)
	if err != nil {
		log.Printf("dynamodbErr: %s", err.Error())
	}

	return
}
