package usecase

import "github.com/naoki85/my-blog-api-sam/interface/database"

type UserRepository interface {
	Create(params database.UserCreateParams) (bool, error)
}
