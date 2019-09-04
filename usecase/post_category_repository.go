package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
)

type PostCategoryRepository interface {
	FindById(int) (model.PostCategory, error)
}
