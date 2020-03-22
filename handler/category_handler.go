package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/controller"
	"github.com/naoki85/my-blog-api-sam/infrastructure"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"log"
)

func categories() (events.APIGatewayProxyResponse, error) {
	c := initConf()
	dynamoDbHandler, _ := infrastructure.NewDynamoDbHandler(c)
	categoryController := controller.NewCategoryController(dynamoDbHandler)
	categories, status := categoryController.Index()

	if status != config.SuccessStatus {
		return handleError(status), nil
	}

	return apiResponse(fmt.Sprintf("%s", categories), status), nil
}

func category(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cIdentifier := request.PathParameters["identifier"]

	c := initConf()
	dynamoDbHandler, _ := infrastructure.NewDynamoDbHandler(c)
	categoryController := controller.NewCategoryController(dynamoDbHandler)

	category, status := categoryController.Show(cIdentifier)
	if status != config.SuccessStatus {
		return handleError(status), nil
	}

	return apiResponse(fmt.Sprintf("%s", category), status), nil
}

func createCategory(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var params usecase.CategoryCreateParams
	requestBody := []byte(request.Body)
	err := json.Unmarshal(requestBody, &params)
	if err != nil {
		return handleError(400), nil
	}

	c := initConf()
	dynamoDbHandler, _ := infrastructure.NewDynamoDbHandler(c)
	categoryController := controller.NewCategoryController(dynamoDbHandler)
	category, status := categoryController.Create(params)

	if status != config.SuccessStatus {
		return handleError(status), nil
	}

	return apiResponse(fmt.Sprintf("%s", category), status), nil
}

func updateCategory(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cIdentifier := request.PathParameters["identifier"]

	var params usecase.CategoryCreateParams
	requestBody := []byte(request.Body)
	err := json.Unmarshal(requestBody, &params)
	if err != nil {
		log.Println("fail to unmarshal")
		return handleError(400), nil
	}
	params.Identifier = cIdentifier

	c := initConf()
	dynamoDbHandler, _ := infrastructure.NewDynamoDbHandler(c)
	categoryController := controller.NewCategoryController(dynamoDbHandler)
	data, status := categoryController.Update(params)

	if status != config.SuccessStatus {
		return handleError(status), nil
	}

	return apiResponse(fmt.Sprintf("%s", data), status), nil
}

func deleteCategory(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cIdentifier := request.PathParameters["identifier"]

	c := initConf()
	dynamoDbHandler, _ := infrastructure.NewDynamoDbHandler(c)
	categoryController := controller.NewCategoryController(dynamoDbHandler)
	data, status := categoryController.Delete(cIdentifier)

	if status != config.SuccessStatus {
		return handleError(status), nil
	}

	return apiResponse(fmt.Sprintf("%s", data), status), nil
}
