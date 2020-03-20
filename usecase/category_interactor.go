package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
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

func (interactor *CategoryInteractor) FindById(identifier string) (category model.Category, err error) {
	category, err = interactor.CategoryRepository.FindById(identifier)
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	return
}

func (interactor *CategoryInteractor) Create(params CategoryCreateParams) (err error) {
	err = interactor.CategoryRepository.Create(params)
	return
}

func (interactor *CategoryInteractor) Update(params CategoryCreateParams) (err error) {
	err = interactor.CategoryRepository.Update(params)
	return
}

func (interactor *CategoryInteractor) Delete(identifier string) (err error) {
	err = interactor.CategoryRepository.Delete(identifier)
	return
}
