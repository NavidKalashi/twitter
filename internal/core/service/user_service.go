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

func (us *UserService) Register(user *models.User) error {

	// check email and username not exist
	email, err := us.repo.EmailExist(user.Email)
	if err != nil {
		return err
	}
	if email != nil {
		return errors.New("email already exist")
	}

	username, err := us.repo.UsernameExist(user.Username)
	if err != nil {
		return err
	}
	if username != nil {
		return errors.New("username already exist")
	}
	
	return us.repo.Register(user)
	
	// create otp for new user
	
	// generate jwt token 
	
	// send email
	
	// return token
}

func (us *UserService) Get(id uuid.UUID) (*models.User, error) {
	return us.repo.Get(id)
}

func (us *UserService) Update(user *models.User) error{
	existingUser, err := us.repo.Get(user.ID)
	if err != nil {
		return errors.New("user not found")
	}

	if user.Name == "" {
		user.Name = existingUser.Name
	}
	if user.Email == "" {
		user.Email = existingUser.Email
	}

	return us.repo.Update(user)
}

func (us *UserService) Delete(id uuid.UUID) error {
	return us.repo.Delete(id)
}

// verify email service method

// step 1: verify jwt signiture
// step 2: get user last otp code
// step 3: compare otp codes
// step 4: change user email status to verified