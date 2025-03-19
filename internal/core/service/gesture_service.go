package service

import (
	"fmt"

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

func (gs *GestureService) AddView(tweetID string, username string, gesType string) error {
	gesture := &models.Gesture{
		TweetID:  tweetID,
		Username: username,
		Type:     "view",
	}

	_, err := gs.gestureRepo.GetByUsername(username, gesType)
	if err != nil {
		tweet, err := gs.tweetRepo.GetByID(tweetID)
		if err != nil {
			return fmt.Errorf("id is not valid: %v", err)
		}
		tweet.Views++
		gs.tweetRepo.Update(tweet)
	} else {
		return fmt.Errorf("you view this tweet before")
	}

	return gs.gestureRepo.Save(gesture)
}

func (gs *GestureService) AddLike(tweetID string, username string, gesType string) error {
	gesture := &models.Gesture{
		TweetID:  tweetID,
		Username: username,
		Type:     "like",
	}

	_, err := gs.gestureRepo.GetByUsername(username, gesType)
	if err != nil {
		tweet, err := gs.tweetRepo.GetByID(tweetID)
		if err != nil {
			return fmt.Errorf("id is not valid: %v", err)
		}
		tweet.Likes++
		gs.tweetRepo.Update(tweet)
	} else {
		return fmt.Errorf("you like this tweet before")
	}

	return gs.gestureRepo.Save(gesture)
}

func (gs *GestureService) AddRetweet(tweetID string, username string, gesType string) error {
	gesture := &models.Gesture{
		TweetID:  tweetID,
		Username: username,
		Type:     "retweet",
	}

	_, err := gs.gestureRepo.GetByUsername(username, gesType)
	if err != nil {
		tweet, err := gs.tweetRepo.GetByID(tweetID)
		if err != nil {
			return fmt.Errorf("id is not valid: %v", err)
		}
		tweet.Retweet++
		gs.tweetRepo.Update(tweet)
	} else {
		return fmt.Errorf("you retweet this tweet before")
	}

	return gs.gestureRepo.Save(gesture)
}
