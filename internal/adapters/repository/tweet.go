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

func (tr *TweetRepository) Create(text, username string) error {
	tweet := &models.Tweet{
		Text: text,
		CreatedBy: username,
	}
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
    result := tr.db.Find(&tweets)
    if result.Error != nil {
        return nil, result.Error
    }
    return tweets, nil
}