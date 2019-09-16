package controller

import (
	"encoding/json"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
	"github.com/naoki85/my-blog-api-sam/usecase"
)

type RecommendedBookController struct {
	Interactor usecase.RecommendedBookInteractor
}

func NewRecommendedBookController(sqlHandler repository.SqlHandler) *RecommendedBookController {
	return &RecommendedBookController{
		Interactor: usecase.RecommendedBookInteractor{
			RecommendedBookRepository: &repository.RecommendedBookRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *RecommendedBookController) Index() ([]byte, int) {
	limit := 4
	recommendedBooks, err := controller.Interactor.RecommendedBookRepository.All(limit)
	if err != nil {
		return []byte{}, config.NotFoundStatus
	}

	data := struct {
		RecommendedBooks model.RecommendedBooks
	}{recommendedBooks}
	resp, err := json.Marshal(data)
	if err != nil {
		return resp, config.InternalServerErrorStatus
	}
	return resp, config.SuccessStatus
}
