package usecase

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
	UserRepository UserRepository
}

type UserInteractorCreateParams struct {
	email    string
	password string
}

func (interactor *UserInteractor) Create(params UserInteractorCreateParams) (bool, error) {
	var encryptedPassword []byte
	var err error
	encryptedPassword, err = bcrypt.GenerateFromPassword([]byte(params.password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	var userCreateParams = UserCreateParams{
		email:    params.email,
		password: fmt.Sprintf("%s", encryptedPassword),
	}
	return interactor.UserRepository.Create(userCreateParams)
}
