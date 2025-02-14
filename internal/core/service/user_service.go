package service

import (
	"errors"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	"github.com/google/uuid"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) CreateUser(user *models.User) error {
	return us.repo.CreateUser(user)
}

func (us *UserService) GetUser(id uuid.UUID) (*models.User, error) {
	return us.repo.GetUser(id)
}

func (us *UserService) UpdateUser(user *models.User) error{
	existingUser, err := us.repo.GetUser(user.ID)
	if err != nil {
		return errors.New("user not found")
	}

	if user.Name == "" {
		user.Name = existingUser.Name
	}
	if user.Email == "" {
		user.Email = existingUser.Email
	}

	return us.repo.UpdateUser(user)
}

func (us *UserService) DeleteUser(id uuid.UUID) error {
	return us.repo.DeleteUser(id)
}