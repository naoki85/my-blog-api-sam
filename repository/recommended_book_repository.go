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

type RecommendedBookRepository struct {
	DynamoDBHandler *dynamodb.DynamoDB
}

type RecommendedBookCreateParams struct {
	Id        int
	Link      string
	ImageUrl  string
	ButtonUrl string
}

func (repo *RecommendedBookRepository) tableName() (tableName string) {
	if tableName = os.Getenv("RECOMMENDED_BOOKS_TABLE_NAME"); tableName != "" {
		return tableName
	} else {
		return "RecommendedBooks"
	}
}

func (repo *RecommendedBookRepository) All() (recommendedBooks model.RecommendedBooks, err error) {
	result, err := repo.DynamoDBHandler.Scan(&dynamodb.ScanInput{
		TableName:            aws.String(repo.tableName()),
		ProjectionExpression: aws.String("Id, Link, ImageUrl, ButtonUrl"),
	})

	if err != nil {
		log.Printf("dynamodbErr: %s", err.Error())
	}

	for _, i := range result.Items {
		book := model.RecommendedBook{}
		err = dynamodbattribute.UnmarshalMap(i, &book)
		if err != nil {
			log.Println("Got error unmarshalling:")
			log.Println(err.Error())
			return
		}
		recommendedBooks = append(recommendedBooks, book)
	}

	return
}

func (repo *RecommendedBookRepository) Create(params RecommendedBookCreateParams) (err error) {
	type Item struct {
		Id        int
		Link      string
		ImageUrl  string
		ButtonUrl string
		CreatedAt string
		UpdatedAt string
	}
	now := time.Now().Format("2006-01-02 15-04-05")

	item := Item{
		Id:        params.Id,
		Link:      params.Link,
		ImageUrl:  params.ImageUrl,
		ButtonUrl: params.ButtonUrl,
		CreatedAt: now,
		UpdatedAt: now,
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
