package usecase

import (
	"fmt"
	"github.com/naoki85/my-blog-api-sam/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
	UserRepository UserRepository
}

type UserInteractorCreateParams struct {
	Email    string
	Password string
}

func (interactor *UserInteractor) Create(params UserInteractorCreateParams) (bool, error) {
	var encryptedPassword []byte
	var err error
	encryptedPassword, err = bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	var userCreateParams = repository.UserCreateParams{
		Email:    params.Email,
		Password: fmt.Sprintf("%s", encryptedPassword),
	}
	return interactor.UserRepository.Create(userCreateParams)
}
