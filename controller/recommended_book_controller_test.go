package controller

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"testing"
)

func TestShouldFindAllRecommendedBooks(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	controller := NewRecommendedBookController(dynamoDbHandler)

	recommendedBooks, status := controller.Index()
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}
	var res struct {
		RecommendedBooks model.RecommendedBooks
	}
	err := json.Unmarshal(recommendedBooks, &res)
	if err != nil {
		t.Fatalf("Response could not pasred: %s", err.Error())
	}
	if len(res.RecommendedBooks) != 4 {
		t.Fatalf("Fail expected length: 4, got: %v", len(res.RecommendedBooks))
	}
}

func TestShouldCreateRecommendedBook(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	controller := NewRecommendedBookController(dynamoDbHandler)

	params := usecase.RecommendedBookInteractorCreateParams{
		Link:      "http://test.example.com/hoge",
		ImageUrl:  "http://test.example.com/hoge.png",
		ButtonUrl: "http://test.example.com/hoge.png",
	}

	_, status := controller.Create(params)
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String("6"),
			},
		},
		TableName: aws.String("RecommendedBooks"),
	}

	_, err := dynamoDbHandler.DeleteItem(input)
	if err != nil {
		t.Fatal("Got error calling DeleteItem")
		t.Fatal(err.Error())
	}
}
