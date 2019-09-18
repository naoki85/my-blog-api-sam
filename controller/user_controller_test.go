package controller

import (
	"encoding/json"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/infrastructure"
	"github.com/naoki85/my-blog-api-sam/repository"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
	"testing"
)

func TestShouldCreateUser(t *testing.T) {
	mockSqlHandler, _ := repository.NewMockSqlHandler()
	mockSqlHandler.Mock.ExpectExec("INSERT INTO users").
		WillReturnResult(sqlmock.NewResult(1, 1))
	controller := NewUserController(mockSqlHandler)

	params := usecase.UserInteractorCreateParams{
		Email:    "test@example.com",
		Password: "password",
	}

	_, status := controller.Create(params)
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}
}

func TestShouldUserLogin(t *testing.T) {
	sqlHandler, tearDown := SetupTest()
	defer tearDown()
	controller := NewUserController(sqlHandler)
	var result struct {
		AuthenticationToken string
	}
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

		res, status := controller.Login(successfulParams)
		if status != config.SuccessStatus {
			t.Fatalf("Should get 200 status, but got: %d", status)
		}
		err := json.Unmarshal(res, &result)
		if err != nil {
			t.Fatalf("Response could not pasred: %s", err.Error())
		}
		if result.AuthenticationToken == "" {
			t.Fatal("Could not get AuthenticationToken")
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

func SetupTest() (repository.SqlHandler, func()) {
	config.InitDbConf("")
	c := config.GetDbConf()
	sqlHandler, err := infrastructure.NewSqlHandler(c)
	if err != nil {
		panic(err.Error())
	}

	return sqlHandler, func() {
		_, _ = sqlHandler.Execute("DELETE FROM users")
	}
}
