package controller

import (
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/interface/database"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
	"testing"
)

func TestShouldCreateUser(t *testing.T) {
	mockSqlHandler, _ := database.NewMockSqlHandler()
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
