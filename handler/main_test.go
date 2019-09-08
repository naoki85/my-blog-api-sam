package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Run("/posts", func(t *testing.T) {
		_, err := handler(events.APIGatewayProxyRequest{Path: "/posts", QueryStringParameters: map[string]string{"page": "1"}})
		if err != nil {
			t.Fatal("Everything should be ok")
		}
	})

	t.Run("/posts/66", func(t *testing.T) {
		_, err := handler(events.APIGatewayProxyRequest{Path: "/posts/66", PathParameters: map[string]string{"id": "1"}})
		if err != nil {
			t.Fatal("Everything should be ok")
		}
	})

	t.Run("/recommended_books", func(t *testing.T) {
		_, err := handler(events.APIGatewayProxyRequest{Path: "/recommended_books"})
		if err != nil {
			t.Fatal("Everything should be ok")
		}
	})
}
