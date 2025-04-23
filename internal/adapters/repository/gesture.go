package repository

import (
	"context"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type GestureRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewGestureRepository(db *gorm.DB, redis *redis.Client) ports.Gesture {
	return &GestureRepository{db: db, redis: redis}
}

func (gr *GestureRepository) Save(gesture *models.Gesture) error {
	return gr.db.Create(gesture).Error
}

func (gr *GestureRepository) Set(tweetID, gestureType, username string) error {
	ctx := context.Background()
	key := "gesture_" + tweetID
	value := username + "_" + gestureType

	items, err := gr.redis.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return err
	}
	for _, item := range items {
		if item == value {
			return nil
		}
	}

	return gr.redis.LPush(ctx, key, value).Err()
}

func (r *GestureRepository) Exists(tweetID, username, typeStr string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Gesture{}).
		Where("tweet_id = ? AND username = ? AND type = ?", tweetID, username, typeStr).
		Count(&count).Error
	return count > 0, err
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
