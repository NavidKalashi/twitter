package repository

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserGormRepository {
	return &UserGormRepository{db: db}
}

func (r *UserGormRepository) RepoCreateUser(user *models.User) error {
    return r.db.Create(user).Error
}