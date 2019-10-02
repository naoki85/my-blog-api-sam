package repository

import (
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
	"testing"
)

func TestShouldFindAllRecommendedBooks(t *testing.T) {
	dynamoDbHandler, _ := testSupport.NewDynamoDbHandler()
	mockSqlHandler, _ := NewMockSqlHandler()
	repo := RecommendedBookRepository{
		SqlHandler:      mockSqlHandler,
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
	mockSqlHandler, _ := NewMockSqlHandler()
	mockSqlHandler.Mock.ExpectExec("INSERT INTO recommended_books").
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := RecommendedBookRepository{
		SqlHandler: mockSqlHandler,
	}
	params := RecommendedBookCreateParams{
		Link:      "http://test.example.com/hoge",
		ImageUrl:  "http://test.example.com/hoge.png",
		ButtonUrl: "http://test.example.com/hoge.png",
	}
	err := repo.Create(params)
	if err != nil {
		t.Fatalf("Cannot create recommended_book: %s", err)
	}
	if err := mockSqlHandler.Mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
