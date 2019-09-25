package main

import (
	"encoding/json"
	"fmt"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/controller"
	"github.com/naoki85/my-blog-api-sam/infrastructure"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"regexp"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.Path == "/health" {
		return health(request)
	} else if request.Path == "/recommended_books" {
		if request.HTTPMethod == "POST" {
			return requireLogin(createRecommendedBook, request)
		} else {
			return recommendedBooks(request)
		}
	} else if request.Path == "/posts" {
		return posts(request)
	} else if regexp.MustCompile(`^/posts/(\d+)`).MatchString(request.Path) {
		return post(request)
	} else if request.HTTPMethod == "POST" && request.Path == "/users" {
		return createUser(request)
	} else if request.HTTPMethod == "POST" && request.Path == "/login" {
		return login(request)
	} else if request.HTTPMethod == "DELETE" && request.Path == "/logout" {
		return logout(request)
	}
	return handleError(404), nil
}

func main() {
	lambda.Start(handler)
}

func requireLogin(f func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error),
	request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	header := request.Headers["Authorization"]
	authenticationToken := strings.Split(header, " ")[1]
	config.InitDbConf("")
	c := config.GetDbConf()
	sqlHandler, _ := infrastructure.NewSqlHandler(c)
	userController := controller.NewUserController(sqlHandler)
	_, status := userController.LoginStatus(authenticationToken)
	if status != config.SuccessStatus {
		return handleError(config.UnauthorizedStatus), nil
	}
	return f(request)
}

func recommendedBooks(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	config.InitDbConf("")
	c := config.GetDbConf()
	sqlHandler, _ := infrastructure.NewSqlHandler(c)
	testController := controller.NewRecommendedBookController(sqlHandler)
	recommendedBooks, status := testController.Index()

	if status != config.SuccessStatus {
		return handleError(status), nil
	}

	return apiResponse(fmt.Sprintf("%s", recommendedBooks), status), nil
}

func createRecommendedBook(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var params usecase.RecommendedBookInteractorCreateParams
	requestBody := []byte(request.Body)
	err := json.Unmarshal(requestBody, &params)
	if err != nil {
		return handleError(400), nil
	}

	config.InitDbConf("")
	c := config.GetDbConf()
	sqlHandler, _ := infrastructure.NewSqlHandler(c)
	testController := controller.NewRecommendedBookController(sqlHandler)
	recommendedBooks, status := testController.Create(params)

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
	testController := controller.NewPostController(sqlHandler)
	resp, status := testController.Index(page)
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
	testController := controller.NewPostController(sqlHandler)

	post, status := testController.Show(postId)
	if status != config.SuccessStatus {
		return handleError(status), nil
	}

	return apiResponse(fmt.Sprintf("%s", post), status), nil
}

func createUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var params usecase.UserInteractorCreateParams
	requestBody := []byte(request.Body)
	err := json.Unmarshal(requestBody, &params)
	if err != nil {
		return handleError(400), nil
	}
	config.InitDbConf("")
	c := config.GetDbConf()
	sqlHandler, _ := infrastructure.NewSqlHandler(c)
	testController := controller.NewUserController(sqlHandler)

	res, status := testController.Create(params)
	if status != config.SuccessStatus {
		return handleError(status), nil
	}
	return apiResponse(fmt.Sprintf("%s", res), status), nil
}

func login(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var params usecase.UserInteractorCreateParams
	requestBody := []byte(request.Body)
	err := json.Unmarshal(requestBody, &params)
	if err != nil {
		return handleError(400), nil
	}
	config.InitDbConf("")
	c := config.GetDbConf()
	sqlHandler, _ := infrastructure.NewSqlHandler(c)
	testController := controller.NewUserController(sqlHandler)

	res, status := testController.Login(params)
	if status != config.SuccessStatus {
		return handleError(status), nil
	}
	return apiResponse(fmt.Sprintf("%s", res), status), nil
}

func logout(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	header := request.Headers["Authorization"]
	authenticationToken := strings.Split(header, " ")[1]
	config.InitDbConf("")
	c := config.GetDbConf()
	sqlHandler, _ := infrastructure.NewSqlHandler(c)
	testController := controller.NewUserController(sqlHandler)

	res, status := testController.Logout(authenticationToken)
	if status != config.SuccessStatus {
		return handleError(status), nil
	}
	return apiResponse(fmt.Sprintf("%s", res), status), nil
}

func health(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return apiResponse("success", config.SuccessStatus), nil
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
