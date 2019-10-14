package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
)

type PostRepository interface {
	All() (model.Posts, int, error)
	FindById(int) (model.Post, error)
	Create(repository.PostCreateParams) error
}
