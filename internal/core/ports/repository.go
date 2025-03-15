package ports

import (
	"time"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/google/uuid"
)

type User interface {
	Register(username, name, email, hashPass, bio string, birthday time.Time) error
	GetByEmail(email string) (*models.User, error)
	GetByID(userID string) (*models.User, error)
	Edit(user *models.User) error
	EmailExist(email string) (*models.User, error)
	UsernameExist(username string) (*models.User, error)
	Verified(user *models.User, sit bool) error
}

type EmailService interface {
	Send(to string, message string) error
	SendOTP(to string, code uint) error
}

type OTP interface {
	Set(email string, code uint) error
	Get(email string) (uint, error)
}

type RefreshToken interface {
	Create(userID uuid.UUID, refreshToken string) error
	Get(userID uuid.UUID) error
	Delete(userID uuid.UUID) error
}

type Tweet interface {
	Create(text, username string) error
	DeleteAll(username string) error
	Delete(username, tweetID string) error
	GetTweets() ([]models.Tweet, error)
}

type Media interface {
	
}