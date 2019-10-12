package repository

import (
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"testing"
)

func TestShouldFindUserByEmail(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	repo := UserRepository{
		DynamoDBHandler: dynamoDbHandler,
	}
	user, err := repo.FindByEmail("hoge@example.com")
	if err != nil {
		t.Fatalf("Could not find user: %s", err.Error())
	}
	if user.Email != "hoge@example.com" {
		t.Fatalf("Fail expected: hoge@example.com, got: %s", user.Email)
	}
}

func TestShouldFindUserByAuthenticationToken(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	repo := UserRepository{
		DynamoDBHandler: dynamoDbHandler,
	}
	user, err := repo.FindByAuthenticationToken("Q1LZlKFt2h0000001vER4TjyFo7")
	if err != nil {
		t.Fatalf("Could not find user: %s", err.Error())
	}
	if user.Id != 3 {
		t.Fatalf("Fail expected id: 1, got: %d", user.Id)
	}
}

func TestShouldUpdateAttribute(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	repo := UserRepository{
		DynamoDBHandler: dynamoDbHandler,
	}
	err := repo.UpdateAttribute("hoge@example.com", "AuthenticationToken", "hogehoge")
	if err != nil {
		t.Fatalf("Could not update: %s", err.Error())
	}
}

func TestShouldCreateUser(t *testing.T) {
	dynamoDbHandler, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	repo := UserRepository{
		DynamoDBHandler: dynamoDbHandler,
	}
	params := UserCreateParams{
		Id:       2,
		Email:    "test@example.com",
		Password: "password",
	}
	err := repo.Create(params)
	if err != nil {
		t.Fatalf("Cannot create user: %s", err)
	}
}
