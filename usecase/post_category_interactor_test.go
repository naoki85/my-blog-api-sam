package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"testing"
)

type MockPostCategoryRepository struct {
}

func (repo *MockPostCategoryRepository) FindById(int) (model.PostCategory, error) {
	postCategory := model.PostCategory{
		Id:    1,
		Name:  "AWS",
		Color: "#ffffff",
	}
	return postCategory, nil
}

func TestShouldFindPostCategoryById(t *testing.T) {
	repo := new(MockPostCategoryRepository)
	interactor := PostCategoryInteractor{
		PostCategoryRepository: repo,
	}
	postCategory, err := interactor.FindById(1)
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
	if postCategory.Name != "AWS" {
		t.Fatalf("Fail expected id: 1, got: %v", postCategory.Name)
	}
}
