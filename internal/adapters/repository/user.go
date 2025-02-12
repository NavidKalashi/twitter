package repository

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) ports.UserRepository {
	return &UserGormRepository{db: db}
}

func (r *UserGormRepository) CreateUser(user *models.User) error {
    return r.db.Create(user).Error
}
