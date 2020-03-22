package controller

import (
	"encoding/json"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"testing"
)

func TestShouldGetCategoriesForIndex(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()

	controller := NewCategoryController(dynamoDbHandler)
	categories, status := controller.Index()
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}
	var res struct {
		Categories model.Categories
	}
	err := json.Unmarshal(categories, &res)
	if err != nil {
		t.Fatalf("Response could not pasred: %s", err.Error())
	}
	if len(res.Categories) != 10 {
		t.Fatalf("Fail expected length: 1, got: %v", res)
	}
}

func TestShouldGetCategoryForShow(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()

	controller := NewCategoryController(dynamoDbHandler)
	category, status := controller.Show("aws")
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}
	var res model.Category
	err := json.Unmarshal(category, &res)
	if err != nil {
		t.Fatalf("Response could not pasred: %s", err.Error())
	}
	if res.Identifier != "aws" {
		t.Fatalf("Fail expected identifier aws, got: %v", res)
	}
}

func TestFromCreateToDeleteThroughUpdate(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	controller := NewCategoryController(dynamoDbHandler)

	params := usecase.CategoryCreateParams{
		Identifier: "test",
		JpName:     "Test",
		Color:      "#000000",
	}

	_, status := controller.Create(params)
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}

	_, status = controller.Update(params)
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}
	_, status = controller.Delete("test")
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}
}
