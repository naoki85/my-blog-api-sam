package usecase

import (
	"github.com/naoki85/my-blog-api-sam/interface/database"
	"testing"
)

type MockUserRepository struct {
}

func (repo *MockUserRepository) Create(params database.UserCreateParams) (bool, error) {
	return true, nil
}

func TestShouldCreateNewUser(t *testing.T) {
	repo := new(MockUserRepository)
	interactor := UserInteractor{
		UserRepository: repo,
	}

	params := UserInteractorCreateParams{
		Email:    "test@example.com",
		Password: "hogehoge",
	}

	_, err := interactor.Create(params)
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
}
