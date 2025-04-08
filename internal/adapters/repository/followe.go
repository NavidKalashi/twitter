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

func (fr *FollowRepository) GetFollowers(username string) ([]models.Follow, error) {
	var follow []models.Follow
	if err := fr.db.Where("following_name = ?", username).Find(&follow).Error; err != nil {
		return nil, err
	}
	return follow, nil
}