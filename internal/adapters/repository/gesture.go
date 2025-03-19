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

func (gr *GestureRepository) Count(tweetID string, gestureType string) (int, error) {
	var count int64
	err := gr.db.Create(&models.Gesture{}).Where("tweet_id = ? AND type = ?", tweetID, gestureType).Count(&count).Error
	return int(count), err
}

func (r *GestureRepository) GetByUsername(username string, gesType string) (*models.Gesture, error) {
	var gesture models.Gesture
	if err := r.db.Where("username = ? AND type = ?", username, gesType).First(&gesture).Error; err != nil {
		return nil, err
	}
	return &gesture, nil
}