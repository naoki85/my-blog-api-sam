package controller

import (
	"fmt"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"testing"
)

var mockSesHandler = testSupport.SetupMockSesHandler()

func TestShouldCreateUser(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	controller := NewUserController(dynamoDbHandler, &mockSesHandler)

	params := usecase.UserInteractorCreateParams{
		Email:    "test@example.com",
		Password: "password",
	}

	_, status := controller.Create(params)
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}
}

func TestShouldUserLoginAndLogout(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	controller := NewUserController(dynamoDbHandler, &mockSesHandler)
	params := usecase.UserInteractorCreateParams{
		Email:    "test@example.com",
		Password: "password",
	}
	_, status := controller.Create(params)
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}

	t.Run("Successful Request", func(t *testing.T) {
		successfulParams := usecase.UserInteractorCreateParams{
			Email:    "test@example.com",
			Password: "password",
		}

		token, status := controller.Login(successfulParams)
		if status != config.SuccessStatus {
			t.Fatalf("Should get 200 status, but got: %d", status)
		}
		_, status = controller.Logout(fmt.Sprintf("%s", token))
		if status != config.SuccessStatus {
			t.Fatalf("Should get 200 status, but got: %d", status)
		}
	})

	t.Run("Login Failure", func(t *testing.T) {
		failureParams := usecase.UserInteractorCreateParams{
			Email:    "test@example.com",
			Password: "fugafuga",
		}
		_, status := controller.Login(failureParams)
		if status != config.InvalidParameterStatus {
			t.Fatalf("Should get 401 status, but got: %d", status)
		}
	})
}

func TestCheckLoginStatus(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	controller := NewUserController(dynamoDbHandler, &mockSesHandler)
	params := usecase.UserInteractorCreateParams{
		Email:    "test@example.com",
		Password: "password",
	}
	_, status := controller.Create(params)
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}

	t.Run("Successful Request", func(t *testing.T) {
		successfulParams := usecase.UserInteractorCreateParams{
			Email:    "test@example.com",
			Password: "password",
		}

		token, status := controller.Login(successfulParams)
		if status != config.SuccessStatus {
			t.Fatalf("Should get 200 status, but got: %d", status)
		}
		_, status = controller.LoginStatus(fmt.Sprintf("%s", token))
		if status != config.SuccessStatus {
			t.Fatalf("Should get 200 status, but got: %d", status)
		}
	})

	t.Run("Login Failure", func(t *testing.T) {
		_, status = controller.LoginStatus("hogehoge")
		if status != config.InvalidParameterStatus {
			t.Fatalf("Should get 401 status, but got: %d", status)
		}
	})
}
