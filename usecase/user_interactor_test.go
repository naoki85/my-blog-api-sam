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

func TestShouldCheckAuthenticationToken(t *testing.T) {
	sqlHandler, tearDown := SetupTest()
	defer tearDown()
	interactor := UserInteractor{
		UserRepository: &repository.UserRepository{
			sqlHandler,
		},
	}

	t.Run("find logged in user", func(t *testing.T) {
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
		user2, err := interactor.CheckAuthenticationToken(user.AuthenticationToken)
		if err != nil {
			t.Fatalf("Could not find user: %s", err.Error())
		}
		if user2.Id == 0 {
			t.Fatalf("Could not find user: %d", user2.Id)
		}
	})

	t.Run("could not find logged in user", func(t *testing.T) {
		_, err := interactor.CheckAuthenticationToken("hogehoge")
		if err == nil {
			t.Fatal("Wrong result")
		}
	})
}

func TestShouldUserLogout(t *testing.T) {
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

	err = interactor.Logout(user.AuthenticationToken)
	if err != nil {
		t.Fatalf("Could not logout: %s", err.Error())
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
		_, _ = sqlHandler.Execute("DELETE FROM recommended_books")
	}
}
