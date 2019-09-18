package repository

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
	"testing"
)

func TestShouldFindUserByEmail(t *testing.T) {
	rows := sqlmock.NewRows([]string{"id", "email", "encrypted_password"}).
		AddRow(1, "hoge@example.com", "encrypted_password")
	mockSqlHandler, _ := NewMockSqlHandler()
	mockSqlHandler.Mock.ExpectQuery("^SELECT (.+) FROM users .*").WillReturnRows(rows)
	repo := UserRepository{
		SqlHandler: mockSqlHandler,
	}
	user, err := repo.FindByEmail("hoge@example.com")
	if err != nil {
		t.Fatalf("Could not find user: %s", err.Error())
	}
	if user.Email != "hoge@example.com" {
		t.Fatalf("Fail expected: hoge@example.com, got: %s", user.Email)
	}
}

func TestShouldUpdateAttribute(t *testing.T) {
	mockSqlHandler, _ := NewMockSqlHandler()
	mockSqlHandler.Mock.ExpectExec("UPDATE users SET").
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := UserRepository{
		SqlHandler: mockSqlHandler,
	}
	_, err := repo.UpdateAttribute(1, "authorization_token", "hogehoge")
	if err != nil {
		t.Fatalf("Could not update: %s", err.Error())
	}
}

func TestShouldCreateUser(t *testing.T) {
	mockSqlHandler, _ := NewMockSqlHandler()
	mockSqlHandler.Mock.ExpectExec("INSERT INTO users").
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
