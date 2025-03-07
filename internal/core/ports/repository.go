package ports

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/google/uuid"
)

type User interface {
	Register(user *models.User) error
	Get(id uuid.UUID) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uuid.UUID) error
	EmailExist(email string) (*models.User, error)
	UsernameExist(username string) (*models.User, error)
	// GetUsers() ([]models.User, error)
}

type EmailService interface {
	Send(to string, message string) error
	SendOTP(to string, code uint) error
}

type OTP interface {
	Create(user *models.User, code uint) error
	FindByUserID(userID string) (*models.OTP, error)
	Verified(otp *models.OTP) error
}

type RefreshToken interface {
	Create(userID uuid.UUID, refreshToken string) error
}

type AccessToken interface {
	Set(userID string, accessToken string) error
	Delete(userID string) error
}