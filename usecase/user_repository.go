package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
)

type UserRepository interface {
	FindBy(string, string) (model.User, error)
	UpdateAttribute(int, string, string) (bool, error)
	Create(repository.UserCreateParams) (bool, error)
}
