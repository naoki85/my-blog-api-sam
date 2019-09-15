package database

type UserRepository struct {
	SqlHandler
}

type UserCreateParams struct {
	email    string
	password string
}

func (repo *UserRepository) Create(params UserCreateParams) (bool, error) {
	query := "INSERT INTO users (email, encrypted_password) VALUES (?, ?)"
	_, err := repo.SqlHandler.Execute(query, params.email, params.password)
	if err != nil {
		return false, err
	}

	return true, nil
}
