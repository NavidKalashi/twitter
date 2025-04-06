package repository

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	"gorm.io/gorm"
)

type GestureRepository struct {
	db *gorm.DB
}

func NewGestureRepository(db *gorm.DB) ports.Gesture {
	return &GestureRepository{db: db}
}

func (gr *GestureRepository) Save(gesture *models.Gesture) error {
	return gr.db.Create(gesture).Error
}

func (gr *GestureRepository) Count(tweetID, username string) (int, error) {
	var count int64
	err := gr.db.Create(&models.Gesture{}).Where("tweet_id = ? AND username = ?", tweetID, username).Count(&count).Error
	return int(count), err
}

func (r *GestureRepository) GetByUsername(tweetID, username string) (*models.Gesture, error) {
	var gesture models.Gesture
	if err := r.db.Where("tweet_id = ? AND username = ?", tweetID, username).First(&gesture).Error; err != nil {
		return nil, err
	}
	return &gesture, nil
}