package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(otp *models.OTP, code uint) error {
	secretKey := []byte("YourSecretKey")

	claims := jwt.MapClaims{
		"code": code,
		"exp": time.Now().Add(5 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatalf("error in creating token: %V", err)
	}
	fmt.Println("jwt token:", tokenString)
	
	return nil
}