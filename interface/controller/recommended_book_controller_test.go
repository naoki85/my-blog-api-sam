package controller

import (
	"encoding/json"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/interface/database"
	"github.com/naoki85/my-blog-api-sam/model"
	"testing"
)

func TestShouldFindAllRecommendedBooks(t *testing.T) {
	mockSqlHandler, _ := database.NewMockSqlHandler()
	mockSqlHandler.ResistMock("^SELECT (.+) FROM recommended_books .*", []string{"id", "link", "image_url", "button_url"})
	controller := NewRecommendedBookController(mockSqlHandler)
	recommendedBooks, status := controller.Index()
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}
	var res struct {
		RecommendedBooks model.RecommendedBooks
	}
	err := json.Unmarshal(recommendedBooks, &res)
	if err != nil {
		t.Fatalf("Response could not pasred: %s", err.Error())
	}
	if len(res.RecommendedBooks) != 4 {
		t.Fatalf("Fail expected length: 4, got: %v", res)
	}
}
