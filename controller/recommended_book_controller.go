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

type RecommendedBookController struct {
	Interactor usecase.RecommendedBookInteractor
}

func NewRecommendedBookController(sqlHandler repository.SqlHandler,
	dynamoDbHandler *dynamodb.DynamoDB) *RecommendedBookController {
	return &RecommendedBookController{
		Interactor: usecase.RecommendedBookInteractor{
			RecommendedBookRepository: &repository.RecommendedBookRepository{
				SqlHandler:      sqlHandler,
				DynamoDBHandler: dynamoDbHandler,
			},
			IdCounterRepository: &repository.IdCounterRepository{
				DynamoDBHandler: dynamoDbHandler,
			},
		},
	}
}

func (controller *RecommendedBookController) Index() ([]byte, int) {
	limit := 4
	recommendedBooks, err := controller.Interactor.All(limit)
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.NotFoundStatus
	}

	data := struct {
		RecommendedBooks model.RecommendedBooks
	}{recommendedBooks}
	resp, err := json.Marshal(data)
	if err != nil {
		log.Printf("%s", err.Error())
		return resp, config.InternalServerErrorStatus
	}
	return resp, config.SuccessStatus
}

func (controller *RecommendedBookController) Create(params usecase.RecommendedBookInteractorCreateParams) ([]byte, int) {
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
