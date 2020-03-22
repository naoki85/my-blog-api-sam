package controller

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"log"
)

type CategoryController struct {
	Interactor usecase.CategoryInteractor
}

func NewCategoryController(dynamoDbHandler *dynamodb.DynamoDB) *CategoryController {
	return &CategoryController{
		Interactor: usecase.CategoryInteractor{
			CategoryRepository: &repository.CategoryRepository{
				DynamoDBHandler: dynamoDbHandler,
			},
		},
	}
}

func (controller *CategoryController) Index() ([]byte, int) {
	categories, err := controller.Interactor.Index()
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.NotFoundStatus
	}

	data := struct {
		Categories model.Categories
	}{categories}
	resp, err := json.Marshal(data)
	if err != nil {
		log.Printf("%s", err.Error())
		return resp, config.InternalServerErrorStatus
	}

	return resp, config.SuccessStatus
}

func (controller *CategoryController) Show(identifier string) ([]byte, int) {
	category, err := controller.Interactor.FindByIdentifier(identifier)
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.NotFoundStatus
	}

	resp, err2 := json.Marshal(category)

	if err2 != nil {
		log.Printf("%s", err2.Error())
		return resp, config.InternalServerErrorStatus
	}

	return resp, config.SuccessStatus
}

func (controller *CategoryController) Create(params usecase.CategoryCreateParams) ([]byte, int) {
	err := controller.Interactor.Create(params)
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.InvalidParameterStatus
	}

	data := struct {
		Message string
	}{"success"}
	resp, err := json.Marshal(data)
	if err != nil {
		log.Printf("%s", err.Error())
		return resp, config.InternalServerErrorStatus
	}
	return resp, config.SuccessStatus
}

func (controller *CategoryController) Update(params usecase.CategoryCreateParams) ([]byte, int) {
	err := controller.Interactor.Update(params)
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.InvalidParameterStatus
	}

	data := struct {
		Message string
	}{"success"}
	resp, err := json.Marshal(data)
	if err != nil {
		log.Printf("%s", err.Error())
		return resp, config.InternalServerErrorStatus
	}
	return resp, config.SuccessStatus
}

func (controller *CategoryController) Delete(identifier string) ([]byte, int) {
	err := controller.Interactor.Delete(identifier)
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.InvalidParameterStatus
	}

	data := struct {
		Message string
	}{"success"}
	resp, err := json.Marshal(data)
	if err != nil {
		log.Printf("%s", err.Error())
		return resp, config.InternalServerErrorStatus
	}
	return resp, config.SuccessStatus
}
