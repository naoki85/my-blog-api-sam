package controller

import (
	"encoding/json"
	"fmt"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"strings"
	"testing"
)

func TestShouldGetPostsForIndex(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()

	controller := NewPostController(dynamoDbHandler)
	posts, status := controller.Index(1, false)
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}
	var res struct {
		TotalPage int
		Posts     model.Posts
	}
	err := json.Unmarshal(posts, &res)
	if err != nil {
		t.Fatalf("Response could not pasred: %s", err.Error())
	}
	if res.TotalPage != 1 {
		t.Fatalf("Fail expected length: 1, got: %v", res)
	}
}

func TestShouldGetPostsForShow(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()

	t.Run("format json", func(t *testing.T) {
		controller := NewPostController(dynamoDbHandler)
		post, status := controller.Show(1, "json")
		if status != config.SuccessStatus {
			t.Fatalf("Should get 200 status, but got: %d", status)
		}
		var res model.Post
		err := json.Unmarshal(post, &res)
		if err != nil {
			t.Fatalf("Response could not pasred: %s", err.Error())
		}
		if res.Id != 1 {
			t.Fatalf("Fail expected length: 1, got: %v", res)
		}
	})

	t.Run("format html", func(t *testing.T) {
		controller := NewPostController(dynamoDbHandler)
		res, status := controller.Show(1, "html")
		if status != config.SuccessStatus {
			t.Fatalf("Should get 200 status, but got: %d", status)
		}
		res2Html := fmt.Sprintf("%s", res)

		if strings.Contains(res2Html, `<meta property="og:title" content="Test">`) == false {
			t.Fatalf("Not match expected strings: %s", res2Html)
		}
	})
}

func TestShouldCreatePost(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	controller := NewPostController(dynamoDbHandler)

	params := usecase.PostInteractorCreateParams{
		Category:    "aws",
		Title:       "Test title",
		Content:     "Test content",
		ImageUrl:    "test.com",
		Active:      "published",
		PublishedAt: "2019-10-01 00:00:00",
	}

	_, status := controller.Create(params)
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}

	params.Id = 2
	_, status = controller.Update(params)
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}
}
