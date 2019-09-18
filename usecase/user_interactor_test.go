package usecase

import (
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/infrastructure"
	"github.com/naoki85/my-blog-api-sam/repository"
	"testing"
)

func TestShouldCreateUserAndLogin(t *testing.T) {
	sqlHandler, tearDown := SetupTest()
	defer tearDown()
	interactor := UserInteractor{
		UserRepository: &repository.UserRepository{
			sqlHandler,
		},
	}

	params := UserInteractorCreateParams{
		Email:    "test@example.com",
		Password: "hogehoge",
	}

	_, err := interactor.Create(params)
	if err != nil {
		t.Fatalf("Could not create user: %s", err.Error())
	}
	user, err := interactor.Login(params)
	if err != nil {
		t.Fatalf("Could not login: %s", err.Error())
	}
	if user.AuthenticationToken == "" {
		t.Fatal("Should set authorization_token to user")
	}
}

func SetupTest() (repository.SqlHandler, func()) {
	config.InitDbConf("")
	c := config.GetDbConf()
	sqlHandler, err := infrastructure.NewSqlHandler(c)
	if err != nil {
		panic(err.Error())
	}

	return sqlHandler, func() {
		_, _ = sqlHandler.Execute("DELETE FROM users")
	}
}
