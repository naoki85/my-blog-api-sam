package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
)

type UserRepository interface {
	FindByEmail(string) (model.User, error)
	FindByAuthenticationToken(string) (model.User, error)
	UpdateAttribute(string, string, string) error
	Create(repository.UserCreateParams) error
}
