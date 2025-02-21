package service

import (
	"errors"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	"github.com/google/uuid"
)

type UserService struct {
	repo ports.UserRepository
	// otp repo
	// email repo
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// change to Regiter
func (us *UserService) CreateUser(user *models.User) error {
	// check email and username shouldn't esxist
	
	return us.repo.CreateUser(user)

	// create otp for new user

	// generate jwt token 

	// send email

	// return token
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

// verify email service method

// step 1: verify jwt signiture
// step 2: get user last otp code
// step 3: compare otp codes
// step 4: change user email status to verified