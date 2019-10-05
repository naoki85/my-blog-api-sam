package repository

import (
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"testing"
)

func TestShouldFindCountByIdentifier(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	repo := IdCounterRepository{
		DynamoDBHandler: dynamoDbHandler,
	}
	count, err := repo.FindMaxIdByIdentifier("RecommendedBooks")
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
	if count != 5 {
		t.Fatalf("Expected count 5, but got: %d", count)
	}
}

func TestShouldUpdateMaxIdByIdentifier(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	repo := IdCounterRepository{
		DynamoDBHandler: dynamoDbHandler,
	}
	count, err := repo.UpdateMaxIdByIdentifier("RecommendedBooks", 6)
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
	if count != 6 {
		t.Fatalf("Expected count 6, but got: %d", count)
	}
}
