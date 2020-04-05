package usecase

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"time"
)

type UserInteractor struct {
	UserRepository      UserRepository
	IdCounterRepository IdCounterRepository
	SesHandler          SesHandler
}

type UserRepository interface {
	FindByEmail(string) (model.User, error)
	FindByAuthenticationToken(string) (model.User, error)
	UpdateAttribute(string, string, string) error
	Create(repository.UserCreateParams) error
}

type SesHandler interface {
	SendMail(string, string, string) error
}

type UserInteractorCreateParams struct {
	Id       int
	Email    string
	Password string
}

var randSrc = rand.NewSource(time.Now().UnixNano())

const (
	rsLetters       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rsLetterIdxBits = 6
	rsLetterIdxMask = 1<<rsLetterIdxBits - 1
	rsLetterIdxMax  = 63 / rsLetterIdxBits
)

func (interactor *UserInteractor) Create(params UserInteractorCreateParams) (err error) {
	var encryptedPassword []byte
	encryptedPassword, err = bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	maxId, err := interactor.IdCounterRepository.FindMaxIdByIdentifier("Users")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	newId := maxId + 1
	_, err = interactor.IdCounterRepository.UpdateMaxIdByIdentifier("Users", newId)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	var userCreateParams = repository.UserCreateParams{
		Id:       newId,
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

func (interactor *UserInteractor) Logout(authenticationToken string) (err error) {
	user, err := interactor.UserRepository.FindByAuthenticationToken(authenticationToken)
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}
	err = interactor.UserRepository.UpdateAttribute(user.Email, "AuthenticationToken", "-")
	if err != nil {
		log.Printf("%s", err.Error())
	}

	return
}

func (interactor *UserInteractor) CheckAuthenticationToken(authenticationToken string) (model.User, error) {
	user, err := interactor.UserRepository.FindByAuthenticationToken(authenticationToken)
	if err != nil {
		log.Printf("%s", err.Error())
		return user, err
	}
	if user.Id == 0 {
		err = errors.New("not found")
	}
	layout := "2006-01-02 15-04-05"
	loc, _ := time.LoadLocation("Asia/Tokyo")
	t, err := time.ParseInLocation(layout, user.AuthenticationTokenExpiredAt, loc)
	if err != nil {
		log.Println(err.Error())
		return user, err
	}
	now := time.Now()
	if now.After(t) {
		err = errors.New("not found")
	}
	return user, err
}

func (interactor *UserInteractor) updateToken(user *model.User) error {
	authenticationToken := interactor.generateToken()
	interactor.SesHandler.SendMail(user.Email, "Test", authenticationToken)
	err := interactor.UserRepository.UpdateAttribute(user.Email, "AuthenticationToken", authenticationToken)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}
	expiredAt := time.Now().Add(6 * time.Hour).Format("2006-01-02 15-04-05")
	err = interactor.UserRepository.UpdateAttribute(user.Email, "AuthenticationTokenExpiredAt",
		expiredAt)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}
	user.AuthenticationToken = authenticationToken
	user.AuthenticationTokenExpiredAt = expiredAt
	return nil
}

func (interactor *UserInteractor) generateToken() string {
	token := rand2String(16)
	token = time.Now().Format("20060102150405") + token
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	return encoded
}

func rand2String(n int) string {
	b := make([]byte, n)
	cache, remain := randSrc.Int63(), rsLetterIdxMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), rsLetterIdxMax
		}
		idx := int(cache & rsLetterIdxMask)
		if idx < len(rsLetters) {
			b[i] = rsLetters[idx]
			i--
		}
		cache >>= rsLetterIdxBits
		remain--
	}
	return string(b)
}
