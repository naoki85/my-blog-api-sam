package usecase

import (
	"github.com/naoki85/my-blog-api-sam/repository"
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"testing"
)

type MockSesHandler struct{}

func (sesHandler *MockSesHandler) SendMail(to string, title string, body string) error {
	return nil
}

func TestShouldCreateUserAndLogin(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	interactor := UserInteractor{
		UserRepository: &repository.UserRepository{
			DynamoDBHandler: dynamoDbHandler,
		},
		IdCounterRepository: &repository.IdCounterRepository{
			DynamoDBHandler: dynamoDbHandler,
		},
		SesHandler: &MockSesHandler{},
	}

	params := UserInteractorCreateParams{
		Email:    "test@example.com",
		Password: "hogehoge",
	}

	err := interactor.Create(params)
	if err != nil {
		t.Fatalf("Could not create user: %s", err.Error())
	}
	user, err := interactor.Login(params)
	if err != nil {
		t.Fatalf("Could not login: %s", err.Error())
	}
	if user.AuthenticationToken == "" {
		t.Fatal("Should set authorization_token to user")
	}
}

func TestShouldCheckAuthenticationToken(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	interactor := UserInteractor{
		UserRepository: &repository.UserRepository{
			DynamoDBHandler: dynamoDbHandler,
		},
		IdCounterRepository: &repository.IdCounterRepository{
			DynamoDBHandler: dynamoDbHandler,
		},
	}

	t.Run("find logged in user", func(t *testing.T) {
		params := UserInteractorCreateParams{
			Email:    "test@example.com",
			Password: "hogehoge",
		}

		err := interactor.Create(params)
		if err != nil {
			t.Fatalf("Could not create user: %s", err.Error())
		}
		user, err := interactor.Login(params)
		if err != nil {
			t.Fatalf("Could not login: %s", err.Error())
		}
		if user.AuthenticationToken == "" {
			t.Fatal("Should set authorization_token to user")
		}
		user2, err := interactor.CheckAuthenticationToken(user.AuthenticationToken)
		if err != nil {
			t.Fatalf("Could not find user: %s", err.Error())
		}
		if user2.Id == 0 {
			t.Fatalf("Could not find user: %d", user2.Id)
		}
	})

	t.Run("could not find logged in user", func(t *testing.T) {
		_, err := interactor.CheckAuthenticationToken("hogehoge")
		if err == nil {
			t.Fatal("Wrong result")
		}
	})
}

func TestShouldUserLogout(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	interactor := UserInteractor{
		UserRepository: &repository.UserRepository{
			DynamoDBHandler: dynamoDbHandler,
		},
		IdCounterRepository: &repository.IdCounterRepository{
			DynamoDBHandler: dynamoDbHandler,
		},
	}

	params := UserInteractorCreateParams{
		Email:    "test@example.com",
		Password: "hogehoge",
	}

	err := interactor.Create(params)
	if err != nil {
		t.Fatalf("Could not create user: %s", err.Error())
	}
	user, err := interactor.Login(params)
	if err != nil {
		t.Fatalf("Could not login: %s", err.Error())
	}
	if user.AuthenticationToken == "" {
		t.Fatal("Should set authorization_token to user")
	}

	err = interactor.Logout(user.AuthenticationToken)
	if err != nil {
		t.Fatalf("Could not logout: %s", err.Error())
	}
}
