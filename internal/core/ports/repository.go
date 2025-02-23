package ports

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/google/uuid"
)

type UserRepository interface {
	Register(user *models.User) error
	Get(id uuid.UUID) (*models.User, error)
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

type OTPRepository interface {
	Create(user *models.User, code uint) error
}