package usecase

import (
	"fmt"
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
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
		log.Printf("%s", err.Error())
		return false, err
	}
	var userCreateParams = repository.UserCreateParams{
		Email:    params.Email,
		Password: fmt.Sprintf("%s", encryptedPassword),
	}
	return interactor.UserRepository.Create(userCreateParams)
}

func (interactor *UserInteractor) Login(params UserInteractorCreateParams) (model.User, error) {
	user, err := interactor.UserRepository.FindByEmail(params.Email)
	if err != nil {
		log.Printf("%s", err.Error())
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(params.Password))
	if err != nil {
		log.Printf("%s", err.Error())
		return user, err
	}
	err = interactor.updateToken(&user)
	if err != nil {
		log.Printf("%s", err.Error())
		return user, err
	}

	return user, err
}

func (interactor *UserInteractor) updateToken(user *model.User) error {
	authenticationToken := interactor.generateToken()
	_, err := interactor.UserRepository.UpdateAttribute(user.Id, "authentication_token", authenticationToken)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}
	user.AuthenticationToken = authenticationToken
	return nil
}

func (interactor *UserInteractor) generateToken() string {
	token := "hoge"
	return token
}
