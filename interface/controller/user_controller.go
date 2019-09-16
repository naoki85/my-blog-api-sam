package controller

import (
	"encoding/json"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/interface/database"
	"github.com/naoki85/my-blog-api-sam/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
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
