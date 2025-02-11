package ports

import "github.com/NavidKalashi/twitter/internal/core/domain/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
	GetUsers() ([]models.User, error)
}