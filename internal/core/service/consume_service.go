package service

import (
	"encoding/json"
	"log"

	"github.com/NavidKalashi/twitter/internal/core/domain/events"
	"github.com/NavidKalashi/twitter/internal/core/ports"
)

type ConsumeService struct {
	consume        ports.Consume
	feedRepo       ports.Feed
	followRepo     ports.Follow
	gestureRepo    ports.Gesture
	tweetService   TweetService
	gestureService GestureService
}

func NewConsumeService(consume ports.Consume, feedRepo ports.Feed, followRepo ports.Follow, gestureRepo ports.Gesture, tweetService TweetService, gestureService GestureService) *ConsumeService {
	return &ConsumeService{
		consume:        consume,
		feedRepo:       feedRepo,
		followRepo:     followRepo,
		gestureRepo:    gestureRepo,
		tweetService:   tweetService,
		gestureService: gestureService,
	}
}

func (cs *ConsumeService) ConsumeFeed() error {
	cs.consume.ConsumeFeedEvents(func(e events.Feed) {
		log.Printf("feed event received: %+v", e)

		followers, err := cs.followRepo.GetFollowers(e.Username)
		if err != nil {
			log.Println(err)
		}

		tweetJSON, err := json.Marshal(e)
		if err != nil {
			log.Printf("Failed to marshal tweet: %v", err)
			return
		}

		if err := cs.tweetService.Create(e.Text, e.Username, e.MediaType, e.FileNames); err != nil {
			log.Printf("Failed to write tweet to postgres: %v", err)
			return
		}

		for _, follower := range followers {
			f := follower.FollowerName
			go func(f string) {
				if err := cs.feedRepo.Set(f, string(tweetJSON)); err != nil {
					log.Printf("Failed to write feed to Redis: %v", err)
				}
			}(f)
		}

	})
	return nil
}

func (cs *ConsumeService) ConsumeGesture() error {
	cs.consume.ConsumeGestureEvents(func(e events.Gesture) {
		log.Printf("gesture event received: %+v", e)

		if err:= cs.gestureService.Add(e.TweetID, e.Username, e.GestureType); err != nil {
			log.Printf("Failed to write gesture to postgres: %v", err)
		}

		if err := cs.gestureRepo.Set(e.TweetID, e.GestureType, e.Username); err != nil {
			log.Printf("Failed to write gesture to redis: %v", err)
		}
	})
	return nil
}
