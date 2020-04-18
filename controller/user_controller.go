package controller

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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

type SesHandler interface {
	SendMail(string, string, string) error
}

func NewUserController(dynamoDbHandler *dynamodb.DynamoDB, sesHandler SesHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &repository.UserRepository{
				DynamoDBHandler: dynamoDbHandler,
			},
			IdCounterRepository: &repository.IdCounterRepository{
				DynamoDBHandler: dynamoDbHandler,
			},
			SesHandler: sesHandler,
		},
	}
}

func (controller *UserController) Create(params usecase.UserInteractorCreateParams) ([]byte, int) {
	err := controller.Interactor.Create(params)
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

func (controller *UserController) Login(params usecase.UserInteractorLoginParams) ([]byte, int) {
	_, err := controller.Interactor.Login(params)
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.InvalidParameterStatus
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

func (controller *UserController) OnetimeToken(params usecase.UserInteractorOnetimeTokenParams) ([]byte, int) {
	user, err := controller.Interactor.OnetimeToken(params)
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

func (controller *UserController) Logout(params string) ([]byte, int) {
	token, err := jwt.Parse(params, func(token *jwt.Token) (interface{}, error) {
		return signingKey(), nil
	})
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.InvalidParameterStatus
	}

	err = controller.Interactor.Logout(fmt.Sprintf("%s", token.Claims.(jwt.MapClaims)["accessToken"]))
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.InternalServerErrorStatus
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

func (controller *UserController) LoginStatus(params string) ([]byte, int) {
	token, err := jwt.Parse(params, func(token *jwt.Token) (interface{}, error) {
		return signingKey(), nil
	})
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.InvalidParameterStatus
	}

	_, err = controller.Interactor.CheckAuthenticationToken(fmt.Sprintf("%s", token.Claims.(jwt.MapClaims)["accessToken"]))
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
