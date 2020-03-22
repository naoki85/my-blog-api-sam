package main

import (
	"fmt"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestCategoriesHandler(t *testing.T) {
	t.Run("categories", func(t *testing.T) {
		_, err := handler(events.APIGatewayProxyRequest{Path: "/categories"})
		if err != nil {
			t.Fatal("Everything should be ok")
		}
	})
}

func TestCreateUpdateDeleteHandler(t *testing.T) {
	_, tearDown := testSupport.SetupTestDynamoDb()
	defer tearDown()
	authToken := testLogin()
	res, _ := handler(events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Path:       "/categories",
		Headers:    map[string]string{"Authorization": fmt.Sprintf("Bearer %s", authToken)},
		Body:       `{"identifier":"test","jpName":"Test","Color":"#000000"}`,
	})
	if res.StatusCode != config.SuccessStatus {
		t.Fatalf("Expected status: 200, but got %v", res.StatusCode)
	}

	res, _ = handler(events.APIGatewayProxyRequest{
		HTTPMethod:     "PUT",
		Path:           "/categories/test",
		PathParameters: map[string]string{"identifier": "test"},
		Headers:        map[string]string{"Authorization": fmt.Sprintf("Bearer %s", authToken)},
		Body:           `{"identifier":"test","jpName":"Test2","Color":"#000001"}`,
	})
	if res.StatusCode != config.SuccessStatus {
		t.Fatalf("Expected status: 200, but got %v", res.StatusCode)
	}

	res, _ = handler(events.APIGatewayProxyRequest{
		HTTPMethod:     "DELETE",
		Path:           "/categories/test",
		PathParameters: map[string]string{"identifier": "test"},
		Headers:        map[string]string{"Authorization": fmt.Sprintf("Bearer %s", authToken)},
	})
	if res.StatusCode != config.SuccessStatus {
		t.Fatalf("Expected status: 200, but got %v", res.StatusCode)
	}
}

func TestCategoryHandler(t *testing.T) {
	res, _ := handler(events.APIGatewayProxyRequest{
		Path:           "/categories/aws",
		Headers:        map[string]string{"content-type": "application/json"},
		PathParameters: map[string]string{"identifier": "aws"},
	})
	if res.StatusCode != config.SuccessStatus {
		t.Fatalf("Expected status: 200, but got %v", res.StatusCode)
	}
}
