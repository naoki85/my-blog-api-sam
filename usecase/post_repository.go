package usecase

import "github.com/naoki85/my-blog-api-sam/model"

type PostRepository interface {
	Index(int) (model.Posts, error)
	FindById(int) (model.Post, error)
	GetPostsCount() (int, error)
}
