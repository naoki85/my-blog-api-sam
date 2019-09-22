package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
)

type RecommendedBookInteractor struct {
	RecommendedBookRepository RecommendedBookRepository
}

type RecommendedBookInteractorCreateParams struct {
	Link      string
	ImageUrl  string
	ButtonUrl string
}

func (interactor *RecommendedBookInteractor) All(limit int) (model.RecommendedBooks, error) {
	recommendedBooks, err := interactor.RecommendedBookRepository.All(limit)
	return recommendedBooks, err
}

func (interactor *RecommendedBookInteractor) Create(params RecommendedBookInteractorCreateParams) error {
	var inputParams = repository.RecommendedBookCreateParams{
		Link:      params.Link,
		ImageUrl:  params.ImageUrl,
		ButtonUrl: params.ButtonUrl,
	}
	err := interactor.RecommendedBookRepository.Create(inputParams)
	return err
}
