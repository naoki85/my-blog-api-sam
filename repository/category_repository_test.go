package repository

import (
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"testing"
)

func TestShouldAllCategories(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	repo := CategoryRepository{
		DynamoDBHandler: dynamoDbHandler,
	}
	categories, err := repo.All()
	if err != nil {
		t.Fatalf("Cannot get categories: %s", err)
	}
	if len(categories) != 10 {
		t.Fatalf("Fail expected: 10, got: %v", len(categories))
	}
}

func TestShouldFindCategoryByIdentifier(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	repo := CategoryRepository{
		DynamoDBHandler: dynamoDbHandler,
	}
	category, err := repo.FindByIdentifier("aws")
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
	if category.JpName != "AWS" {
		t.Fatalf("Fail expected: AWS, got: %v", category.JpName)
	}
}

func TestFromCreatingCategoryToDeletingCategoryThroughUpdate(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	repo := CategoryRepository{
		DynamoDBHandler: dynamoDbHandler,
	}
	identifier := "book"
	params := CategoryCreateParams{
		Identifier: "book",
		JpName:     "書評",
		Color:      "#f8f9fa",
	}
	err := repo.Create(params)
	if err != nil {
		t.Fatalf("fail to create category: %s", err)
	}

	name := "機械学習"
	color := "#9c27b0"

	params = CategoryCreateParams{
		Identifier: identifier,
		JpName:     name,
		Color:      color,
	}
	err = repo.Update(params)
	if err != nil {
		t.Fatalf("fail to update category: %s", err)
	}

	category, _ := repo.FindByIdentifier(identifier)
	if category.JpName != name {
		t.Fatalf("Expacted: %s, but got %s", name, category.JpName)
	}
	if category.Color != color {
		t.Fatalf("Expacted: %s, but got %s", color, category.Color)
	}

	err = repo.Delete(identifier)
	if err != nil {
		t.Fatalf("fail to delete post: %s", err)
	}
}
