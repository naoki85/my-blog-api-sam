package controller

import (
	"encoding/json"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
	"testing"
)

func TestShouldGetPostsForIndex(t *testing.T) {
	// TODO: 複数クエリをモックにするならもはや DB 用意した方が良さそう
	mockSqlHandler, _ := repository.NewMockSqlHandler()
	mockSqlHandler.ResistMockForPostsIndex("^SELECT (.+) FROM posts .*", []string{"id", "post_category_id", "title", "image_file_name", "published_at"})
	mockSqlHandler.ResistMockForPostCategory("^SELECT (.+) FROM post_categories (.+)", []string{"id", "name", "color"})
	mockSqlHandler.ResistMockForPostCount("^SELECT COUNT(.+) FROM posts .*", []string{"count"})
	controller := NewPostController(mockSqlHandler)
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
	if res.TotalPage != 7 {
		t.Fatalf("Fail expected length: 4, got: %v", res)
	}
}

func TestShouldGetPostsForShow(t *testing.T) {
	mockSqlHandler, _ := repository.NewMockSqlHandler()
	mockSqlHandler.ResistMockForPost("^SELECT (.+) FROM posts .*", []string{"id", "post_category_id", "title", "content", "image_file_name", "published_at"})
	controller := NewPostController(mockSqlHandler)
	post, status := controller.Show(1)
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
}
