package repository

import (
	"blaster_balu/models"
	"database/sql"
	"errors"
)

type UserRepository interface {
	CreateUser(user models.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(user models.User) error {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)"

	_, err := r.db.Exec(query, user.UserName, user.Email, user.Password)
	if err != nil {
		return errors.New("could not create user")
	}
	return nil
}
