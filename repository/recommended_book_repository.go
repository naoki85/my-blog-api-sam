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
	SqlHandler
	DynamoDBHandler *dynamodb.DynamoDB
}

type RecommendedBookCreateParams struct {
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

func (repo *RecommendedBookRepository) Create(params RecommendedBookCreateParams) error {
	query := "INSERT INTO recommended_books (link, image_url, button_url, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	now := time.Now().Format("2006-01-02 15-04-05")
	_, err := repo.SqlHandler.Execute(query, params.Link, params.ImageUrl, params.ButtonUrl, now, now)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	return nil
}
