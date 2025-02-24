package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(email, username string) (string, error) {
	var secretKey = []byte("your_secret_key")

	claims := jwt.MapClaims{
		"email": email,
		"exp":  time.Now().Add(5 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}