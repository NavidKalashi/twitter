package repository

import (
	"fmt"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	"gorm.io/gorm"
)

type TweetRepository struct {
	db *gorm.DB
}

func NewTweetRepository(db *gorm.DB) ports.Tweet {
	return &TweetRepository{db: db}
}

func (tr *TweetRepository) Create(tweet *models.Tweet) error {
	return tr.db.Create(tweet).Error
}

func (tr *TweetRepository) DeleteAll(username string) error {
	result := tr.db.Where("created_by = ?", username).Delete(&models.Tweet{})
	return result.Error
}

func (tr *TweetRepository) Delete(username, tweetID string) error {
	result := tr.db.Where("id = ? AND created_by = ?", tweetID, username).Delete(&models.Tweet{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("tweet not found or unauthorized")
	}
	return nil
}

func (tr *TweetRepository) GetTweets() ([]models.Tweet, error) {
	var tweets []models.Tweet
	err := tr.db.Preload("Media").Find(&tweets).Error
	return tweets, err
}

func (tr *TweetRepository) GetByID(tweetID string) (*models.Tweet, error) {
	var tweet models.Tweet
	err := tr.db.Where("id = ?", tweetID).First(&tweet).Error
	return &tweet, err
}

func (tr *TweetRepository) Update(tweet *models.Tweet) error {
	return tr.db.Save(tweet).Error
}

func (tr *TweetRepository) GetByUsername(username string) ([]models.Tweet, error) {
	var tweets []models.Tweet
	if err := tr.db.Preload("Media").Preload("Gesture").Where("created_by = ?", username).Find(&tweets).Error; err != nil {
		return nil, err
	}
	return tweets, nil
}