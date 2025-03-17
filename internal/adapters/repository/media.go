package repository

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	"gorm.io/gorm"
)

type MediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) ports.Media {
	return &MediaRepository{db: db}
}

func (mr *MediaRepository) SaveMedia(media *models.Media) error {
	return mr.db.Create(media).Error
}