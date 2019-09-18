package repository

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"log"
	"time"
)

type UserRepository struct {
	SqlHandler
}

type UserCreateParams struct {
	Email    string
	Password string
}

func (repo *UserRepository) FindByEmail(email string) (user model.User, err error) {
	query := "SELECT id, email, encrypted_password FROM users WHERE email = ? LIMIT 1"
	rows, err := repo.SqlHandler.Query(query, email)
	if err != nil {
		log.Printf("%s", err.Error())
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Email, &user.EncryptedPassword)
		if err != nil {
			log.Printf("%s", err.Error())
			return user, err
		}
		break
	}
	return user, err
}

func (repo *UserRepository) UpdateAttribute(id int, field string, param string) (bool, error) {
	query := "UPDATE users SET " + field + " = ?, updated_at = ? WHERE id = ? LIMIT 1"
	now := time.Now().Format("2006-01-02 03-04-05")
	_, err := repo.SqlHandler.Execute(query, param, now, id)
	if err != nil {
		log.Printf("%s", err.Error())
		return false, err
	}

	return true, nil
}

func (repo *UserRepository) Create(params UserCreateParams) (bool, error) {
	query := "INSERT INTO users (email, encrypted_password, created_at, updated_at) VALUES (?, ?, ?, ?)"
	now := time.Now().Format("2006-01-02 03-04-05")
	_, err := repo.SqlHandler.Execute(query, params.Email, params.Password, now, now)
	if err != nil {
		log.Printf("%s", err.Error())
		return false, err
	}

	return true, nil
}