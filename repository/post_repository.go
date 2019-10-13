package repository

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/naoki85/my-blog-api-sam/model"
	"log"
	"os"
	"strconv"
	"time"
)

type PostRepository struct {
	DynamoDBHandler *dynamodb.DynamoDB
}

func (repo *PostRepository) tableName() (tableName string) {
	if tableName = os.Getenv("POSTS_TABLE_NAME"); tableName != "" {
		return tableName
	} else {
		return "Posts"
	}
}

type PostCreateParams struct {
	Id          int
	UserId      int
	Category    string
	Title       string
	Content     string
	ImageUrl    string
	Active      string
	PublishedAt string
}

func (repo *PostRepository) All() (posts model.Posts, count int, err error) {
	result, err := repo.DynamoDBHandler.Scan(&dynamodb.ScanInput{
		TableName:            aws.String(repo.tableName()),
		ProjectionExpression: aws.String("Id, Category, Title, Content, ImageUrl, PublishedAt"),
	})

	if err != nil {
		log.Printf("dynamodbErr: %s", err.Error())
		return
	}

	for _, i := range result.Items {
		p := model.Post{}
		err = dynamodbattribute.UnmarshalMap(i, &p)
		if err != nil {
			log.Println("Got error unmarshalling:")
			log.Println(err.Error())
			return
		}
		posts = append(posts, p)
	}

	var postsCount struct {
		Count int
	}
	res2json, err := json.Marshal(result)
	if err != nil {
		log.Println("Got error marshalling:")
		log.Println(err.Error())
	}
	err = json.Unmarshal(res2json, &postsCount)
	if err != nil {
		log.Println("Got error unmarshalling:")
		log.Println(err.Error())
	}
	count = postsCount.Count

	return
}

func (repo *PostRepository) FindById(id int) (post model.Post, err error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(id)),
			},
		},
		TableName: aws.String(repo.tableName()),
	}

	result, err := repo.DynamoDBHandler.GetItem(input)
	if err != nil {
		log.Printf("dynamodbErr: %s", err.Error())
		return
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &post)
	if err != nil {
		log.Printf("Could not unmarshal map: %s", err.Error())
		return
	}
	return
}

func (repo *PostRepository) Create(params PostCreateParams) (err error) {
	type Item struct {
		Id          int
		UserId      int
		Category    string
		Title       string
		Content     string
		ImageUrl    string
		Active      string
		PublishedAt string
		CreatedAt   string
		UpdatedAt   string
	}
	now := time.Now().Format("2006-01-02 15-04-05")

	item := Item{
		Id:          params.Id,
		UserId:      params.UserId,
		Category:    params.Category,
		Title:       params.Title,
		Content:     params.Content,
		ImageUrl:    params.ImageUrl,
		Active:      params.Active,
		PublishedAt: params.PublishedAt,
		CreatedAt:   now,
		UpdatedAt:   now,
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
