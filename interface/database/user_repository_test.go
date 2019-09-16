package database

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
	"testing"
)

func TestShouldCreateUser(t *testing.T) {
	mockSqlHandler, _ := NewMockSqlHandler()
	mockSqlHandler.Mock.ExpectExec("INSERT INTO users").
		WithArgs("test@example.com", "password").
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := UserRepository{
		SqlHandler: mockSqlHandler,
	}
	params := UserCreateParams{
		Email:    "test@example.com",
		Password: "password",
	}
	_, err := repo.Create(params)
	if err != nil {
		t.Fatalf("Cannot create user: %s", err)
	}
	if err := mockSqlHandler.Mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
