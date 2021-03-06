package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
)

type RecommendedBookRepository interface {
	All() (model.RecommendedBooks, error)
	Create(params repository.RecommendedBookCreateParams) error
}
