package usecase

import "github.com/naoki85/my-blog-api-sam/repository"

type UserRepository interface {
	Create(params repository.UserCreateParams) (bool, error)
}
