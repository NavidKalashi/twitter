package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	jwtPackage "github.com/NavidKalashi/twitter/pkg/jwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepo     ports.UserRepository
	OTPRepo      ports.OTPRepository
	emailService EmailService
}

func NewUserService(UserRepo ports.UserRepository, OTPRepo ports.OTPRepository) *UserService {
	return &UserService{
		UserRepo: UserRepo,
		OTPRepo:  OTPRepo,
	}
}

func generateOTP() uint {
	rand.Seed(time.Now().UnixNano())
	return uint(rand.Intn(900000) + 100000)
}

func (us *UserService) Register(user *models.User) (string, error) {

	// check email and username not exist
	email, err := us.UserRepo.EmailExist(user.Email)
	if err != nil {
		return "", err
	}
	if email != nil {
		return "", errors.New("email already exist")
	}

	username, err := us.UserRepo.UsernameExist(user.Username)
	if err != nil {
		return "", err
	}
	if username != nil {
		return "", errors.New("username already exist")
	}

	us.UserRepo.Register(user)

	// create otp for new user
	code := generateOTP()
	us.OTPRepo.Create(user, code)

	// generate jwt token
	token, err := jwtPackage.GenerateJwt(user.Email)
	if err != nil {
		return "", errors.New("token not created")
	}

	// send email
	us.emailService.SendOTP(user.Email, code)

	return token, nil
}

func (us *UserService) Verify(tokenString string, userID string, code uint) error {
	var secretKey = []byte("your_secret_key")
	claims := &jwt.MapClaims{}

	// token verify
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if token.Valid {
		if exp, ok := (*claims)["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return fmt.Errorf("token has expired")
			}
		} else {
			return fmt.Errorf("expiration claim missing or invalid")
		}

		if email, ok := (*claims)["email"].(string); ok {
			if email == "" {
				return fmt.Errorf("email claim is empty")
			}
		} else {
			return fmt.Errorf("email claim missing or invalid")
		}
	}

	// check otp code
	otp, err := us.OTPRepo.FindByUserID(userID)

	if err != nil {
		return fmt.Errorf("failed to find OTP: %w", err)
	}

	expOtp := otp.CreatedAt.Add(45 * time.Second)
	currentTime := time.Now()

	if currentTime.Unix() > expOtp.Unix() {
		return fmt.Errorf("OTP has expired")
	}

	if otp.Code != code {
		return fmt.Errorf("invalid OTP code")
	} else {
		otp.Verified = true
		if err := us.OTPRepo.Verified(otp); err != nil {
			return fmt.Errorf("failed to change otp status")
		}
	}

	return nil
}

func (us *UserService) Resend(id uuid.UUID) error {
	userByID, err := us.UserRepo.Get(id)
	if err != nil {
		return errors.New("user not found")
	}
	code := generateOTP()
	us.OTPRepo.Create(userByID, code)
	us.emailService.SendOTP(userByID.Email, code)
	return nil
}

func (us *UserService) Get(id uuid.UUID) (*models.User, error) {
	return us.UserRepo.Get(id)
}

func (us *UserService) Update(user *models.User) error {
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
