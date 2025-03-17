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
	"golang.org/x/crypto/bcrypt"
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

var secretKey = []byte("your_secret_key")

func generateOTP() uint {
	rand.Seed(time.Now().UnixNano())
	return uint(rand.Intn(900000) + 100000)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (us *UserService) Register(username, name, email, password, bio string, birthday time.Time) (string, error) {

	// check email and username not exist
	existEmail, err := us.UserRepo.EmailExist(email)
	if err != nil {
		return "", err
	}
	if existEmail != nil {
		return "", errors.New("email already exist")
	}

	existUsername, err := us.UserRepo.UsernameExist(username)
	if err != nil {
		return "", err
	}
	if existUsername != nil {
		return "", errors.New("username already exist")
	}

	hashPass, err := HashPassword(password)
	if err != nil {
		return "", err
	}

	us.UserRepo.Register(username, name, email, hashPass, bio, birthday)

	// create otp for new user
	code := generateOTP()
	us.OTPRepo.Set(email, code)

	// generate jwt token
	token, err := jwtPackage.OTPToken(email)
	if err != nil {
		return "", errors.New("token not created")
	}

	// send email
	us.emailService.SendOTP(email, code)

	return token, nil
}

func (us *UserService) Verify(tokenString string, code uint) (string, string, error) {
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

			user, err := us.UserRepo.GetByEmail(email)
			if err != nil {
				return "", "", err
			}
			fmt.Println("email: ", user.Email, user.ID, user.Username)

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
		} else {
			return "", "", fmt.Errorf("email claim missing or invalid")
		}
	}

	return "", "", nil
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
			"email":    email,
			"sub":      user.ID,
			"exp":      time.Now().Add(time.Hour * 24 * 2).Unix(),
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

func (us *UserService) Login(email string, password string) (string, string, error) {
	user, err := us.UserRepo.GetByEmail(email)
	if err != nil {
		return "", "", fmt.Errorf("email not found")
	}

	if user.OTPVerified {
		err = us.RefreshTokenRepo.Get(user.ID)
		if err == nil {
			return "", "", fmt.Errorf("user is loged in")
		}
		
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			fmt.Println("Invalid credentials")
		}

		refreshToken, accessToken, err := jwtPackage.GenerateAccessAndRefresh(user)
		if err != nil {
			return "", "", errors.New("refresh token not created")
		}

		err = us.RefreshTokenRepo.Create(user.ID, refreshToken)
		if err != nil {
			return "", "", err
		}

		return refreshToken, accessToken, nil
	} else {
		return "", "", fmt.Errorf("user not verified")
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

	token, err := jwtPackage.OTPToken(user.Email)
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