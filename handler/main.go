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
	return handleError(404), nil
}

func main() {
	lambda.Start(handler)
}

func recommendedBooks(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	config.InitDbConf("")
	c := config.GetDbConf()
	sqlHandler, _ := infrastructure.NewSqlHandler(c)
	controller := controller.NewRecommendedBookController(sqlHandler)
	recommendedBooks, status := controller.Index()

	if status != config.SuccessStatus {
		return handleError(status), nil
	}

	return apiResponse(fmt.Sprintf("%s", recommendedBooks), status), nil
}

func posts(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	page, err := strconv.Atoi(request.QueryStringParameters["page"])
	if err != nil {
		return handleError(400), nil
	}

	config.InitDbConf("")
	c := config.GetDbConf()
	sqlHandler, _ := infrastructure.NewSqlHandler(c)
	controller := controller.NewPostController(sqlHandler)
	resp, status := controller.Index(page)
	if status != config.SuccessStatus {
		return handleError(status), nil
	}

	return apiResponse(fmt.Sprintf("%s", resp), status), nil
}

func post(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	postId, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return handleError(400), nil
	}

	config.InitDbConf("")
	c := config.GetDbConf()
	sqlHandler, _ := infrastructure.NewSqlHandler(c)
	controller := controller.NewPostController(sqlHandler)

	post, status := controller.Show(postId)
	if status != config.SuccessStatus {
		return handleError(status), nil
	}

	return apiResponse(fmt.Sprintf("%s", post), status), nil
}

func handleError(status int) events.APIGatewayProxyResponse {
	var message string
	switch status {
	case config.InvalidParameterStatus:
		message = "Invalid Parameter"
	case config.NotFoundStatus:
		message = "Not Found"
	default:
		message = "Internal Server Error"
	}

	return apiResponse(message, status)
}

func apiResponse(message string, status int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       message,
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: status,
	}
}
