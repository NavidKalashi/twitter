package jwt

import (
	"time"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your_secret_key")

func OTPToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(20 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateAccessAndRefresh(user *models.User) (string, string, error) {
	// access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"sub":      user.ID,
		"exp":      time.Now().Add(time.Hour * 24 * 2).Unix(),
	})
	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	// refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	return refreshTokenString, accessTokenString, nil
}
