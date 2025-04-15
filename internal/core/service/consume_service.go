package service

import (
	"encoding/json"
	"log"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
)

type ConsumeService struct {
	consume    ports.Consume
	feedRepo   ports.Feed
	followRepo ports.Follow
	gestureRepo ports.Gesture
}

func NewConsumeService(consume ports.Consume, feedRepo ports.Feed, followRepo ports.Follow, gestureRepo ports.Gesture) *ConsumeService {
	return &ConsumeService{consume: consume,feedRepo: feedRepo, followRepo: followRepo, gestureRepo: gestureRepo}
}

func (cs *ConsumeService) ConsumeFeed() error {
	cs.consume.ConsumeFeedEvents(func(e models.Tweet) {
		log.Printf("feed event received: %+v", e)

		followers, err := cs.followRepo.GetFollowers(e.CreatedBy)
		if err != nil {
			log.Println(err)
		}

		tweetJSON, err := json.Marshal(e)
		if err != nil {
			log.Printf("Failed to marshal tweet: %v", err)
			return
		}

		for _, follower := range followers {
			go func(f string) {
				err := cs.feedRepo.Set(f, string(tweetJSON))
				if err != nil {
					log.Printf("Failed to write feed to Redis: %v", err)
				}
			}(follower.FollowerName)
		}
	})
	return nil
}

func (cs *ConsumeService) ConsumeGesture() error {
	cs.consume.ConsumeGestureEvents(func(e models.Gesture) {
		log.Printf("gesture event received: %+v", e)

		err := cs.gestureRepo.Set(e.TweetID, e.Type, e.Username)
		if err != nil {
			log.Printf("Failed to write gesture to redis: %v", err)
		}
	})
	return nil
}