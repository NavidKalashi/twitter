package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(email string) (string, error) {
	var secretKey = []byte("your_secret_key")
	claims := jwt.MapClaims{
		"email": email,
		"exp":  time.Now().Add(20 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(&secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateAccessAndRefresh(userID string) (string, string, error) {
	var secretKey = []byte("your_secret_key")
	// access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	   })
	accessTokenString, err := accessToken.SignedString(&secretKey)
	if err != nil {
		return "", "", err
	}

	// refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	   })
	refreshTokenString, err := refreshToken.SignedString(&secretKey)
	if err != nil {
		return "", "", err
	}

	return refreshTokenString, accessTokenString, nil
}