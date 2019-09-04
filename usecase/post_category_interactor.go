package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
)

type PostCategoryInteractor struct {
	PostCategoryRepository PostCategoryRepository
}

func (interactor *PostCategoryInteractor) FindById(id int) (model.PostCategory, error) {
	postCategory, err := interactor.PostCategoryRepository.FindById(id)
	if err != nil {
		return postCategory, err
	}
	return postCategory, err
}
