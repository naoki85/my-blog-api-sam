package controller

import (
	"encoding/json"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/repository"
	"github.com/naoki85/my-blog-api-sam/usecase"
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
		return []byte{}, config.NotFoundStatus
	}

	data := struct {
		Message string
	}{"success"}
	resp, err := json.Marshal(data)
	if err != nil {
		return resp, config.InternalServerErrorStatus
	}
	return resp, config.SuccessStatus
}
