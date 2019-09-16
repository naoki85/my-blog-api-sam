package database

import "time"

type UserRepository struct {
	SqlHandler
}

type UserCreateParams struct {
	Email    string
	Password string
}

func (repo *UserRepository) Create(params UserCreateParams) (bool, error) {
	query := "INSERT INTO users (email, encrypted_password, created_at, updated_at) VALUES (?, ?, ?, ?)"
	now := time.Now().Format("2006-01-02 03-04-05")
	_, err := repo.SqlHandler.Execute(query, params.Email, params.Password, now, now)
	if err != nil {
		return false, err
	}

	return true, nil
}
