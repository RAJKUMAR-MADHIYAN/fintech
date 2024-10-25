package services

import (
	"blaster_balu/models"
	"blaster_balu/repository"
	"errors"
)

type UserService interface {
	Signup(user models.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Signup(user models.User) error {
	if user.UserName == "" || user.Email == "" || user.Password == "" {
		return errors.New("missing required fields")
	}
	return s.repo.CreateUser(user)
}
