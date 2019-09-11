package main

import (
	"github.com/naoki85/my-blog-api-sam/config"
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

	t.Run("posts/66", func(t *testing.T) {
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

func TestHealthHandler(t *testing.T) {
	t.Run("Successful Request", func(t *testing.T) {
		res, _ := health(events.APIGatewayProxyRequest{})
		if res.StatusCode != config.SuccessStatus {
			t.Fatalf("Expected status: 200, but got %v", res.StatusCode)
		}
	})
}
