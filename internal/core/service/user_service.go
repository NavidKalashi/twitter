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
	AccessTokenRepo  ports.AccessToken
	emailService     EmailService
}

func NewUserService(UserRepo ports.User, OTPRepo ports.OTP, RefreshTokenRepo ports.RefreshToken, AccessTokenRepo ports.AccessToken) *UserService {
	return &UserService{
		UserRepo:         UserRepo,
		OTPRepo:          OTPRepo,
		RefreshTokenRepo: RefreshTokenRepo,
		AccessTokenRepo:  AccessTokenRepo,
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

func (us *UserService) Verify(tokenString string, userID string, code uint) (string, string, error) {
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

	// check otp code
	otp, err := us.OTPRepo.FindByUserID(userID)

	if err != nil {
		return "", "", fmt.Errorf("failed to find OTP: %w", err)
	}

	expOtp := otp.CreatedAt.Add(2 * time.Minute)
	currentTime := time.Now()

	if currentTime.Unix() > expOtp.Unix() {
		return "", "", fmt.Errorf("OTP has expired")
	}

	if otp.Code != code {
		return "", "", fmt.Errorf("invalid OTP code")
	} else {
		otp.Verified = true
		if err := us.OTPRepo.Verified(otp); err != nil {
			return "", "", fmt.Errorf("failed to change otp status")
		}
	}

	// refresh token and access token
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return "", "", errors.New("invalid user ID format")
	}
	refreshToken, accessToken, err := jwtPackage.GenerateAccessAndRefresh(userID)
	if err != nil {
		return "", "", errors.New("refresh token not created")
	}

	us.RefreshTokenRepo.Create(userUUID, refreshToken)
	us.AccessTokenRepo.Set(userID, accessToken)

	return refreshToken, accessToken, nil
}

func (us *UserService) NewAccessToken(refreshTokenString string, userID string) (string, error) {
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
		userID := claims["user_id"].(string)

		newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": userID,
			"exp": time.Now().Add(time.Minute * 15).Unix(),
		})

		newAccessTokenString, err := newAccessToken.SignedString(secretKey)
		if err != nil {
			return "", err
		}
		
		us.AccessTokenRepo.Set(userID, newAccessTokenString)

		return newAccessTokenString, nil
	} else {
		return "", errors.New("invalid access token")
	}
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
