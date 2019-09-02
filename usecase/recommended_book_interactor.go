package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
)

type RecommendedBookInteractor struct {
	RecommendedBookRepository RecommendedBookRepository
}

func (interactor *RecommendedBookInteractor) All(limit int) (model.RecommendedBooks, error) {
	recommendedBooks, err := interactor.RecommendedBookRepository.All(limit)
	return recommendedBooks, err
}
