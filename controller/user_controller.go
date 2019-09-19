package controller

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/repository"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"log"
	"os"
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
	_, err := controller.Interactor.Create(params)
	if err != nil {
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

	resp, err := generateJwtToken(user.AuthenticationToken)
	if err != nil {
		log.Printf("%s", err.Error())
		return resp, config.InternalServerErrorStatus
	}
	return resp, config.SuccessStatus
}

func generateJwtToken(base string) ([]byte, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["accessToken"] = base
	tokenString, err := token.SignedString(signingKey())
	return []byte(tokenString), err
}

func signingKey() []byte {
	if len(os.Getenv("SIGNINGKEY")) == 0 {
		return []byte("hogehoge")
	} else {
		return []byte(os.Getenv("SIGNINGKEY"))
	}
}
