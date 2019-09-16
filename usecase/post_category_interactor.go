package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"log"
)

type PostCategoryInteractor struct {
	PostCategoryRepository PostCategoryRepository
}

func (interactor *PostCategoryInteractor) FindById(id int) (model.PostCategory, error) {
	postCategory, err := interactor.PostCategoryRepository.FindById(id)
	if err != nil {
		log.Printf("%s", err.Error())
		return postCategory, err
	}
	return postCategory, err
}
