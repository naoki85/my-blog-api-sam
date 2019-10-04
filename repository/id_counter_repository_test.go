package repository

import (
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"testing"
)

func TestShouldFindCountByIdentifier(t *testing.T) {
	dynamoDbHandler, _ := testSupport.NewDynamoDbHandler()
	repo := IdCounterRepository{
		DynamoDBHandler: dynamoDbHandler,
	}
	count, err := repo.FindCountByIdentifier("RecommendedBooks")
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
	if count != 5 {
		t.Fatalf("Expected count 5, but got: %d", count)
	}
}
