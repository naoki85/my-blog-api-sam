package main

import (
	"fmt"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/infrastructure"
	"github.com/naoki85/my-blog-api-sam/interface/controller"
	"regexp"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.Path == "/recommended_books" {
		return recommendedBooks(request)
	} else if request.Path == "/posts" {
		return posts(request)
	} else if regexp.MustCompile(`^/posts/(\d+)`).MatchString(request.Path) {
		return post(request)
	}
	return notFound()
}

func main() {
	lambda.Start(handler)
}

func recommendedBooks(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sqlHandler, _ := infrastructure.NewSqlHandler()
	controller := controller.NewRecommendedBookController(sqlHandler)
	recommendedBooks, status := controller.Index()

	if status != config.SuccessStatus {
		return notFound()
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%s", recommendedBooks),
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: status,
	}, nil
}

func posts(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	page, err := strconv.Atoi(request.QueryStringParameters["page"])
	if err != nil {
		return invalidParameter()
	}

	sqlHandler, _ := infrastructure.NewSqlHandler()
	controller := controller.NewPostController(sqlHandler)
	resp, status := controller.Index(page)
	if status != config.SuccessStatus {
		return notFound()
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%s", resp),
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: status,
	}, nil
}

func post(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	postId, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return invalidParameter()
	}

	sqlHandler, _ := infrastructure.NewSqlHandler()
	controller := controller.NewPostController(sqlHandler)

	post, status := controller.Show(postId)
	if status != config.SuccessStatus {
		return notFound()
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%s", post),
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}

func invalidParameter() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprint("Invalid Parameter"),
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: config.InvalidParameterStatus,
	}, nil
}

func notFound() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprint("Not Found"),
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: config.NotFoundStatus,
	}, nil
}

func internalServerError() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprint("Internal Server Error"),
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: config.InternalServerErrorStatus,
	}, nil
}
