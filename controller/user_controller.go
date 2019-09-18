package controller

import (
	"encoding/json"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/repository"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"log"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler repository.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &repository.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(params usecase.UserInteractorCreateParams) ([]byte, int) {
	res, err := controller.Interactor.Create(params)
	if err != nil || res == false {
		log.Printf("%s", err.Error())
		return []byte{}, config.NotFoundStatus
	}

	data := struct {
		Message string
	}{"success"}
	resp, err := json.Marshal(data)
	if err != nil {
		log.Printf("%s", err.Error())
		return resp, config.InternalServerErrorStatus
	}
	return resp, config.SuccessStatus
}

func (controller *UserController) Login(params usecase.UserInteractorCreateParams) ([]byte, int) {
	user, err := controller.Interactor.Login(params)
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.InvalidParameterStatus
	}

	data := struct {
		AuthenticationToken string
	}{user.AuthenticationToken}
	resp, err := json.Marshal(data)
	if err != nil {
		log.Printf("%s", err.Error())
		return resp, config.InternalServerErrorStatus
	}
	return resp, config.SuccessStatus
}
