package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/model"
	"log"
	"time"
)

type RecommendedBookRepository struct {
	SqlHandler
}

type RecommendedBookCreateParams struct {
	Link      string
	ImageUrl  string
	ButtonUrl string
}

func (repo *RecommendedBookRepository) All(limit int) (recommendedBooks model.RecommendedBooks, err error) {
	config.InitDbConf("")
	c := config.GetDbConf()
	dynamoSession, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("hogehoge", "fugafuga", ""),
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String(c.DynamoDbEndpoint),
	})
	if err != nil {
		panic(err)
	}

	db := dynamodb.New(dynamoSession)
	response, err2 := db.Scan(&dynamodb.ScanInput{
		TableName:            aws.String("RecommendedBooks"),
		ProjectionExpression: aws.String("Id, Link, ImageUrl, ButtonUrl"),
	})
	log.Printf("dynamodb: %v", response)
	if err2 != nil {
		log.Printf("dynamodbErr: %v", err2)
	}

	query := "SELECT id, link, image_url, button_url FROM recommended_books"
	query = query + " ORDER BY id DESC LIMIT ?"
	rows, err := repo.SqlHandler.Query(query, limit)
	if err != nil {
		log.Printf("%s", err.Error())
		return recommendedBooks, err
	}
	defer rows.Close()

	for rows.Next() {
		r := model.RecommendedBook{}
		if err := rows.Scan(&r.Id, &r.Link, &r.ImageUrl, &r.ButtonUrl); err != nil {
			log.Printf("%s", err.Error())
			return recommendedBooks, err
		}

		recommendedBooks = append(recommendedBooks, r)
	}
	return
}

func (repo *RecommendedBookRepository) Create(params RecommendedBookCreateParams) error {
	query := "INSERT INTO recommended_books (link, image_url, button_url, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	now := time.Now().Format("2006-01-02 03-04-05")
	_, err := repo.SqlHandler.Execute(query, params.Link, params.ImageUrl, params.ButtonUrl, now, now)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	return nil
}
