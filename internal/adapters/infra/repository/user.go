package repository

import (
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserGormRepository {
	return &UserGormRepository{db: db}
}