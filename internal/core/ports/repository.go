package ports

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/google/uuid"
)

type User interface {
	Register(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(userID string) (*models.User, error)
	Edit(user *models.User) error
	Delete(userID uuid.UUID) error
	EmailExist(email string) (*models.User, error)
	UsernameExist(username string) (*models.User, error)
	Verified(user *models.User, sit bool) error
	// GetUsers() ([]models.User, error)
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
	Delete(userID uuid.UUID) error
}