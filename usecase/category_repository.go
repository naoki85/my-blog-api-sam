package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
)

type CategoryRepository interface {
	All() (model.Categories, error)
	FindById(string) (model.Category, error)
	Create(CategoryCreateParams) error
	Update(CategoryCreateParams) error
	Delete(string) error
}
