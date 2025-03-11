package service

import (
	// "context"
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
	UserRepo         ports.User
	OTPRepo          ports.OTP
	RefreshTokenRepo ports.RefreshToken
	emailService     EmailService
}

func NewUserService(UserRepo ports.User, OTPRepo ports.OTP, RefreshTokenRepo ports.RefreshToken) *UserService {
	return &UserService{
		UserRepo:         UserRepo,
		OTPRepo:          OTPRepo,
		RefreshTokenRepo: RefreshTokenRepo,
	}
}

func generateOTP() uint {
	rand.Seed(time.Now().UnixNano())
	return uint(rand.Intn(900000) + 100000)
}

var secretKey = []byte("your_secret_key")

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
	us.OTPRepo.Set(user.Email, code)

	// generate jwt token
	token, err := jwtPackage.GenerateJwt(user.Email)
	if err != nil {
		return "", errors.New("token not created")
	}

	// send email
	us.emailService.SendOTP(user.Email, code)

	return token, nil
}

func (us *UserService) Verify(tokenString string, email string, code uint) (string, string, error) {
	claims := &jwt.MapClaims{}

	// token verify
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return "", "", err
	}

	if token.Valid {
		if exp, ok := (*claims)["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return "", "", fmt.Errorf("token has expired")
			}
		} else {
			return "", "", fmt.Errorf("expiration claim missing or invalid")
		}

		if email, ok := (*claims)["email"].(string); ok {
			if email == "" {
				return "", "", fmt.Errorf("email claim is empty")
			}
		} else {
			return "", "", fmt.Errorf("email claim missing or invalid")
		}
	}

	user, err := us.UserRepo.GetByEmail(email)
	if err != nil {
		return "", "", err
	}

	// check otp code
	otp, err := us.OTPRepo.Get(user.Email)
	if err != nil {
		return "", "", fmt.Errorf("failed to find OTP: %w", err)
	}

	if otp != code {
		return "", "", fmt.Errorf("invalid OTP code")
	} else {
		if err := us.UserRepo.Verified(user, true); err != nil {
			return "", "", fmt.Errorf("failed to change otp status")
		}
	}

	// refresh token and access token
	refreshToken, accessToken, err := jwtPackage.GenerateAccessAndRefresh(user)
	if err != nil {
		return "", "", errors.New("refresh token not created")
	}

	err = us.RefreshTokenRepo.Create(user.ID, refreshToken)
	if err != nil {
		return "", "", err
	}

	return refreshToken, accessToken, nil
}

func (us *UserService) NewAccessToken(refreshTokenString string) (string, error) {
	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil || !refreshToken.Valid {
		return "", fmt.Errorf("invalid refres token")
	}

	// generate new access token
	if claims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
		email := claims["email"].(string)

		user, err := us.UserRepo.GetByEmail(email)
		if err != nil {
			return "", fmt.Errorf("fail to find user")
		}

		newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"email": email,
			"sub":   user.ID,
			"exp":   time.Now().Add(time.Minute * 15).Unix(),
		})

		newAccessTokenString, err := newAccessToken.SignedString(secretKey)
		if err != nil {
			return "", err
		}

		return newAccessTokenString, nil
	} else {
		return "", errors.New("invalid access token")
	}
}

func (us *UserService) Logout(userID string) error {
	user, err := us.GetByID(userID)
    if err != nil {
        return fmt.Errorf("failed to get user by id: %w", err)
    }

    err = us.UserRepo.Verified(user, false)
    if err != nil {
        return fmt.Errorf("failed to unverify user: %w", err)
    }

    err = us.RefreshTokenRepo.Delete(user.ID)
    if err != nil {
        return fmt.Errorf("failed to delete refresh token: %w", err)
    }

    return nil
}

func (us *UserService) Resend(email string) (string, error) {
	code := generateOTP()
	user, err := us.UserRepo.GetByEmail(email)
	if err != nil {
		return "", err
	}
	us.OTPRepo.Set(user.Email, code)
	us.emailService.SendOTP(user.Email, code)

	token, err := jwtPackage.GenerateJwt(user.Email)
	if err != nil {
		return "", errors.New("token not created")
	}
	return token, nil
}

func (us *UserService) GetByEmail(email string) (*models.User, error) {
	return us.UserRepo.GetByEmail(email)
}

func (us *UserService) GetByID(userID string) (*models.User, error) {
	return us.UserRepo.GetByID(userID)
}

func (us *UserService) Edit(user *models.User) error {
	existingUser, err := us.UserRepo.GetByEmail(user.Email)
	if err != nil {
		return errors.New("user not found")
	}

	if user.Name == "" {
		user.Name = existingUser.Name
	}
	if user.Email == "" {
		user.Email = existingUser.Email
	}
	user.UpdatedAt = time.Now()

	return us.UserRepo.Edit(user)
}

func (us *UserService) Delete(id uuid.UUID) error {
	return us.UserRepo.Delete(id)
}