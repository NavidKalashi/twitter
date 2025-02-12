package ports

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"gorm.io/gorm"
)

type UserGormRepository struct{
	db *gorm.DB
}

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
	GetUsers() ([]models.User, error)
}

func NewUserGormRepository(db *gorm.DB) *UserGormRepository {
	return &UserGormRepository{db: db}
}

func (r *UserGormRepository) RepoCreateUser(user *models.User) error {
	return r.db.Create(user).Error
}