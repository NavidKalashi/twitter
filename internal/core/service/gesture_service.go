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

func (gs *GestureService) Add(tweetID, username, typeStr string) error {
	exists, err := gs.gestureRepo.Exists(tweetID, username, typeStr)
	if err != nil {
		return err
	}
	if exists {
		return nil 
	}

	gesture := &models.Gesture{
		TweetID:  tweetID,
		Username: username,
		Type:     typeStr,
	}

	return gs.gestureRepo.Save(gesture)
}
