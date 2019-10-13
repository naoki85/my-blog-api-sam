package repository

import (
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"testing"
)

func TestShouldAllPosts(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	repo := PostRepository{
		DynamoDBHandler: dynamoDbHandler,
	}
	posts, count, err := repo.All()
	if err != nil {
		t.Fatalf("Cannot get posts: %s", err)
	}
	if len(posts) != 1 {
		t.Fatalf("Fail expected: 1, got: %v", len(posts))
	}
	if posts[0].Title != "Test" {
		t.Fatalf("Fail expected: Test, got: %v", posts[0].Title)
	}
	if count != 1 {
		t.Fatalf("Fail expected: 1, got: %d", count)
	}
}

func TestShouldFindPostById(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	repo := PostRepository{
		DynamoDBHandler: dynamoDbHandler,
	}
	post, err := repo.FindById(1)
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
	if post.Title != "Test" {
		t.Fatalf("Fail expected: Test, got: %v", post.Title)
	}
}

func TestShouldCreatePost(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	repo := PostRepository{
		DynamoDBHandler: dynamoDbHandler,
	}
	params := PostCreateParams{
		Id:          2,
		UserId:      1,
		Category:    "aws",
		Title:       "Test title",
		Content:     "Test content",
		Active:      "published",
		PublishedAt: "2019-10-01 00:00:00",
	}
	err := repo.Create(params)
	if err != nil {
		t.Fatalf("Cannot create recommended_book: %s", err)
	}
}
