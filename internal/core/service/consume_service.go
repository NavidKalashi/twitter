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
}

func NewConsumeService(consume ports.Consume, feedRepo ports.Feed, followRepo ports.Follow) *ConsumeService {
	return &ConsumeService{consume: consume,feedRepo: feedRepo, followRepo: followRepo}
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
			err = cs.feedRepo.Set(follower.FollowerName, string(tweetJSON))
			if err != nil {
				log.Printf("Failed to write feed to Redis: %v", err)
			}
		}
	})
	return nil
}
