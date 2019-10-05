package repository

import (
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"log"
	"testing"
)

func TestShouldFindAllRecommendedBooks(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	repo := RecommendedBookRepository{
		DynamoDBHandler: dynamoDbHandler,
	}
	recommendedBooks, err := repo.All()
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
	if len(recommendedBooks) != 5 {
		t.Fatalf("Fail expected: 5, got: %v", len(recommendedBooks))
	}
}

func TestShouldCreateRecommendedBook(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	repo := RecommendedBookRepository{
		DynamoDBHandler: dynamoDbHandler,
	}
	params := RecommendedBookCreateParams{
		Id:        6,
		Link:      "http://test.example.com/hoge",
		ImageUrl:  "http://test.example.com/hoge.png",
		ButtonUrl: "http://test.example.com/hoge.png",
	}
	err := repo.Create(params)
	log.Println(err)
	if err != nil {
		t.Fatalf("Cannot create recommended_book: %s", err)
	}
}
