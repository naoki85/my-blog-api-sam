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

type PostItem struct {
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

func (repo *PostRepository) Create(params PostCreateParams) (err error) {
	now := time.Now().Format("2006-01-02 15-04-05")

	item := PostItem{
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

func (repo *PostRepository) Update(params PostCreateParams) (err error) {
	now := time.Now().Format("2006-01-02 15-04-05")

	item := PostItem{
		Id:          params.Id,
		UserId:      params.UserId,
		Category:    params.Category,
		Title:       params.Title,
		Content:     params.Content,
		ImageUrl:    params.ImageUrl,
		Active:      params.Active,
		PublishedAt: params.PublishedAt,
		UpdatedAt:   now,
	}

	stmt := "set UserId = :userId, Category = :category, Title = :title, Content = :content, ImageUrl = :imageUrl,"
	stmt += "Active = :active, PublishedAt = :publishedAt, UpdatedAt = :updatedAt"

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":userId": {
				N: aws.String(strconv.Itoa(item.UserId)),
			},
			":category": {
				S: aws.String(item.Category),
			},
			":title": {
				S: aws.String(item.Title),
			},
			":content": {
				S: aws.String(item.Content),
			},
			":imageUrl": {
				S: aws.String(item.ImageUrl),
			},
			":active": {
				S: aws.String(item.Active),
			},
			":publishedAt": {
				S: aws.String(item.PublishedAt),
			},
			":updatedAt": {
				S: aws.String(item.UpdatedAt),
			},
		},
		TableName: aws.String(repo.tableName()),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(item.Id)),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String(stmt),
	}

	_, err = repo.DynamoDBHandler.UpdateItem(input)
	if err != nil {
		log.Printf("dynamodbErr: %s", err.Error())
		return
	}

	return nil
}
