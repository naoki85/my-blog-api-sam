package main

import (
	"fmt"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/infrastructure"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Run("posts", func(t *testing.T) {
		_, err := handler(events.APIGatewayProxyRequest{Path: "/posts", QueryStringParameters: map[string]string{"page": "1"}})
		if err != nil {
			t.Fatal("Everything should be ok")
		}
	})

	t.Run("posts/1", func(t *testing.T) {
		_, err := handler(events.APIGatewayProxyRequest{Path: "/posts/66", PathParameters: map[string]string{"id": "1"}})
		if err != nil {
			t.Fatal("Everything should be ok")
		}
	})

	t.Run("recommended_books", func(t *testing.T) {
		_, err := handler(events.APIGatewayProxyRequest{Path: "/recommended_books"})
		if err != nil {
			t.Fatal("Everything should be ok")
		}
	})
}

func TestCreateRecommendedBookHandler(t *testing.T) {
	t.Run("Successful Request", func(t *testing.T) {
		res, _ := createRecommendedBook(events.APIGatewayProxyRequest{
			HTTPMethod: "POST",
			Path:       "/recommended_books",
			Body:       `{"link":"http://test.example.com","image_url":"http://test.example.com","button_url":"http://test.example.com"}`,
		})
		if res.StatusCode != config.SuccessStatus {
			t.Fatalf("Expected status: 200, but got %v", res.StatusCode)
		}
	})
}

func TestCreateUserHandler(t *testing.T) {
	t.Run("Successful Request", func(t *testing.T) {
		res, _ := createUser(events.APIGatewayProxyRequest{
			HTTPMethod: "POST",
			Path:       "/users",
			Body:       `{"email":"hoge@example.com","password":"hogehoge"}`,
		})
		if res.StatusCode != config.SuccessStatus {
			t.Fatalf("Expected status: 200, but got %v", res.StatusCode)
		}
	})
}

func TestLoginAndLogoutHandler(t *testing.T) {
	_, teardown := SetupTest()
	defer teardown()
	_, _ = createUser(events.APIGatewayProxyRequest{
		Body: `{"email":"hoge@example.com","password":"hogehoge"}`,
	})

	t.Run("Successful Request", func(t *testing.T) {
		res, _ := login(events.APIGatewayProxyRequest{
			HTTPMethod: "POST",
			Path:       "/login",
			Body:       `{"email":"hoge@example.com","password":"hogehoge"}`,
		})
		if res.StatusCode != config.SuccessStatus {
			t.Fatalf("Expected status: 200, but got %v", res.StatusCode)
		}

		logoutResult, _ := logout(events.APIGatewayProxyRequest{
			HTTPMethod: "DELETE",
			Path:       "/logout",
			Headers:    map[string]string{"Authorization": fmt.Sprintf("Bearer %s", res.Body)},
		})
		if logoutResult.StatusCode != config.SuccessStatus {
			t.Fatalf("Expected status: 200, but got %v", res.StatusCode)
		}
	})

	t.Run("Invalid Request", func(t *testing.T) {
		res, _ := login(events.APIGatewayProxyRequest{
			HTTPMethod: "POST",
			Path:       "/login",
			Body:       `{"email":"hoge@example.com","password":"fugafuga"}`,
		})
		if res.StatusCode != config.InvalidParameterStatus {
			t.Fatalf("Expected status: 401, but got %v", res.StatusCode)
		}
	})
}

func TestHealthHandler(t *testing.T) {
	_, teardown := SetupTest()
	defer teardown()

	t.Run("Successful Request", func(t *testing.T) {
		res, _ := health(events.APIGatewayProxyRequest{})
		if res.StatusCode != config.SuccessStatus {
			t.Fatalf("Expected status: 200, but got %v", res.StatusCode)
		}
	})
}

func SetupTest() (bool, func()) {
	config.InitDbConf("")
	c := config.GetDbConf()
	sqlHandler, err := infrastructure.NewSqlHandler(c)
	if err != nil {
		panic(err.Error())
	}

	return true, func() {
		_, _ = sqlHandler.Execute("DELETE FROM users")
	}
}
