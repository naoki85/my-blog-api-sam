package controller

import (
	"encoding/json"
	"fmt"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"strings"
	"testing"
)

func TestShouldGetPostsForIndex(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()

	controller := NewPostController(dynamoDbHandler)
	posts, status := controller.Index(1)
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
