package usecase

type UserCreateParams struct {
	email    string
	password string
}

type UserRepository interface {
	Create(UserCreateParams) (bool, error)
}
