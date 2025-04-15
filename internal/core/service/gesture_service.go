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

func (gs *GestureService) AddView(tweetID, username, typeStr string) (*models.Gesture, error) {
	gesture := &models.Gesture{
		TweetID:  tweetID,
		Username: username,
		Type:     typeStr,
	}

	return gesture, gs.gestureRepo.Save(gesture)
}
