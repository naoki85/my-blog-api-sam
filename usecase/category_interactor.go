package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
	"log"
)

type CategoryInteractor struct {
	CategoryRepository CategoryRepository
}

type CategoryCreateParams struct {
	Identifier string `json:"identifier"`
	JpName     string `json:"jpname"`
	Color      string `json:"color"`
}

func (interactor *CategoryInteractor) Index() (categories model.Categories, err error) {
	categories, err = interactor.CategoryRepository.All()
	if err != nil {
		log.Printf("%s", err.Error())
	}

	return
}

func (interactor *CategoryInteractor) FindByIdentifier(identifier string) (category model.Category, err error) {
	category, err = interactor.CategoryRepository.FindByIdentifier(identifier)
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	return
}

func (interactor *CategoryInteractor) Create(params CategoryCreateParams) (err error) {
	inputParams := repository.CategoryCreateParams{
		Identifier: params.Identifier,
		JpName:     params.JpName,
		Color:      params.Color,
	}
	err = interactor.CategoryRepository.Create(inputParams)
	return
}

func (interactor *CategoryInteractor) Update(params CategoryCreateParams) (err error) {
	inputParams := repository.CategoryCreateParams{
		Identifier: params.Identifier,
		JpName:     params.JpName,
		Color:      params.Color,
	}
	err = interactor.CategoryRepository.Update(inputParams)
	return
}

func (interactor *CategoryInteractor) Delete(identifier string) (err error) {
	err = interactor.CategoryRepository.Delete(identifier)
	return
}
