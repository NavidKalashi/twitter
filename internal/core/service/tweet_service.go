package service

import (
	"fmt"
	"time"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
)

type TweetService struct {
	tweetRepo ports.Tweet
	mediaRepo ports.Media
	storage   ports.Storage
}

func NewTweetService(tweetRepo ports.Tweet, mediaRepo ports.Media, storage ports.Storage) *TweetService {
	return &TweetService{tweetRepo: tweetRepo, mediaRepo: mediaRepo, storage: storage}
}

func (ts *TweetService) Create(text, username, fileType string, mediaFiles []string) (*models.Tweet, error) {
	tweet := &models.Tweet{
		Text:      text,
		CreatedBy: username,
		CreatedAt: time.Now(),
	}

	if err := ts.tweetRepo.Create(tweet); err != nil {
		return nil, err
	}

	for _, filePath := range mediaFiles {
		fileURL, err := ts.storage.UploadMedia(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to upload media: %v", err)
		}

		media := &models.Media{
			TweetID:   tweet.ID,
			Type:      fileType,
			FileName:  filePath,
			FileURL:   fileURL,
			CreatedAt: time.Now(),
		}

		if err := ts.mediaRepo.SaveMedia(media); err != nil {
			return nil, fmt.Errorf("failed to save media: %v", err)
		}
	}

	return tweet, nil
}

func (ts *TweetService) DeleteAll(username string) error {
	return ts.tweetRepo.DeleteAll(username)
}

func (ts *TweetService) Delete(username, tweetID string) error {
	return ts.tweetRepo.Delete(username, tweetID)
}

func (ts *TweetService) GetTweets() ([]models.Tweet, error) {
	return ts.tweetRepo.GetTweets()
}
