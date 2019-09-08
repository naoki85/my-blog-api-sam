package database

import (
	"testing"
)

func TestShouldPostsIndex(t *testing.T) {
	mockSqlHandler, _ := NewMockSqlHandler()
	mockSqlHandler.ResistMockForPostsIndex("^SELECT (.+) FROM posts .*", []string{"id", "post_category_id", "title", "image_file_name", "published_at"})
	repo := PostRepository{
		SqlHandler: mockSqlHandler,
	}
	posts, err := repo.Index(1)
	if err != nil {
		t.Fatalf("Cannot get posts: %s", err)
	}
	if len(posts) != 5 {
		t.Fatalf("Fail expected: 5, got: %v", len(posts))
	}
	if posts[0].Title != "test title 1" {
		t.Fatalf("Fail expected: test title 1, got: %v", posts[0].Title)
	}
}

func TestShouldFindPostById(t *testing.T) {
	mockSqlHandler, _ := NewMockSqlHandler()
	mockSqlHandler.ResistMockForPost("^SELECT (.+) FROM posts .*", []string{"id", "post_category_id", "title", "content", "image_file_name", "published_at"})
	repo := PostRepository{
		SqlHandler: mockSqlHandler,
	}
	post, err := repo.FindById(1)
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
	if post.Title != "test title 1" {
		t.Fatalf("Fail expected: test title 1, got: %v", post.Title)
	}
}

func TestShouldGetPostsCount(t *testing.T) {
	mockSqlHandler, _ := NewMockSqlHandler()
	mockSqlHandler.ResistMockForPostCount("^SELECT (.+) FROM posts .*", []string{"count"})
	repo := PostRepository{
		SqlHandler: mockSqlHandler,
	}

	count, err := repo.GetPostsCount()
	if err != nil {
		t.Fatalf("Cannot get post: %s", err)
	}

	expected := 68
	if count != expected {
		t.Fatalf("Fail expected: %v, got: %v", expected, count)
	}
}
