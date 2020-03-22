package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
)

type CategoryRepository interface {
	All() (model.Categories, error)
	FindByIdentifier(string) (model.Category, error)
	Create(repository.CategoryCreateParams) error
	Update(repository.CategoryCreateParams) error
	Delete(string) error
}
