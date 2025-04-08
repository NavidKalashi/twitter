package repository

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	"gorm.io/gorm"
)

type FollowRepository struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) ports.Follow {
	return &FollowRepository{db: db}
}

func (fr *FollowRepository) Save(follow *models.Follow) error {
	return fr.db.Create(follow).Error
}
