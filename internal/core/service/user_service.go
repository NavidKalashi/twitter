package service

import (
	"errors"
	"math/rand"
	"time"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	"github.com/NavidKalashi/twitter/pkg/jwt"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepo ports.UserRepository
	OTPRepo	ports.OTPRepository
	EmailService
}

func NewUserService(UserRepo ports.UserRepository, OTPRepo	ports.OTPRepository) *UserService {
	return &UserService{
		UserRepo: UserRepo,
		OTPRepo: OTPRepo,
	}
}

func generateOTP() uint {
	rand.Seed(time.Now().UnixNano())
	return uint(rand.Intn(900000) + 100000)
}

func (us *UserService) Register(user *models.User, otp *models.OTP) error {

	// check email and username not exist
	email, err := us.UserRepo.EmailExist(user.Email)
	if err != nil {
		return err
	}
	if email != nil {
		return errors.New("email already exist")
	}

	username, err := us.UserRepo.UsernameExist(user.Username)
	if err != nil {
		return err
	}
	if username != nil {
		return errors.New("username already exist")
	}
	
	
	us.UserRepo.Register(user)
	
	// create otp for new user
	code := generateOTP()
	us.OTPRepo.Create(user, code)
	
	// generate jwt token
	tokenString := jwt.GenerateToken(otp, code)

	// send email
	us.EmailService.SendOTP(user.Email, code)
	
	// return token
	return tokenString
}

func (us *UserService) Get(id uuid.UUID) (*models.User, error) {
	return us.UserRepo.Get(id)
}

func (us *UserService) Update(user *models.User) error{
	existingUser, err := us.UserRepo.Get(user.ID)
	if err != nil {
		return errors.New("user not found")
	}

	if user.Name == "" {
		user.Name = existingUser.Name
	}
	if user.Email == "" {
		user.Email = existingUser.Email
	}

	return us.UserRepo.Update(user)
}

func (us *UserService) Delete(id uuid.UUID) error {
	return us.UserRepo.Delete(id)
}

// verify email service method

// step 1: verify jwt signiture
// step 2: get user last otp code
// step 3: compare otp codes
// step 4: change user email status to verified