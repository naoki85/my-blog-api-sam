package usecase

import (
	"testing"
)

type MockUserRepository struct {
}

func (repo *MockUserRepository) Create(params UserCreateParams) (bool, error) {
	return true, nil
}

func TestShouldCreateNewUser(t *testing.T) {
	repo := new(MockUserRepository)
	interactor := UserInteractor{
		UserRepository: repo,
	}

	params := UserInteractorCreateParams{
		email:    "test@example.com",
		password: "hogehoge",
	}

	_, err := interactor.Create(params)
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
}
