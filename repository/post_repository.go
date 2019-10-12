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
