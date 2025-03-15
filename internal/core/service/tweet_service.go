package service

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
)

type TweetService struct {
	tweetRepo ports.Tweet
	mediaRepo ports.Media
}

func NewTweetService(tweetRepo ports.Tweet, mediaRepo ports.Media) *TweetService {
	return &TweetService{tweetRepo: tweetRepo, mediaRepo: mediaRepo}
}

func (ts *TweetService) Create(text, username string) error {
	return ts.tweetRepo.Create(text, username)
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
