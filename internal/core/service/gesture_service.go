package service

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
)

type GestureService struct {
	gestureRepo ports.Gesture
	tweetRepo   ports.Tweet
}

func NewGestureService(gestureRepo ports.Gesture, tweetRepo ports.Tweet) *GestureService {
	return &GestureService{gestureRepo: gestureRepo, tweetRepo: tweetRepo}
}

func (gs *GestureService) AddView(tweetID string, username string) error {
	gesture := &models.Gesture{
		TweetID:  tweetID,
		Username: username,
		Type:     "view",
	}

	return gs.gestureRepo.Save(gesture)
}

func (gs *GestureService) AddLike(tweetID string, username string) error {
	gesture := &models.Gesture{
		TweetID:  tweetID,
		Username: username,
		Type:     "like",
	}

	return gs.gestureRepo.Save(gesture)
}

func (gs *GestureService) AddRetweet(tweetID string, username string) error {
	gesture := &models.Gesture{
		TweetID:  tweetID,
		Username: username,
		Type:     "retweet",
	}

	return gs.gestureRepo.Save(gesture)
}
