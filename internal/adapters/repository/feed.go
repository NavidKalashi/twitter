package repository

import (
	"context"

	"github.com/NavidKalashi/twitter/internal/core/ports"
	"github.com/redis/go-redis/v9"
)

type FeedRepository struct {
	redis *redis.Client
}

func NewFeedRepository(redis *redis.Client) ports.Feed {
	return &FeedRepository{redis: redis}
}

func (fr *FeedRepository) Set(username, tweet string) error {
	ctx := context.Background()
	key := "feed_" + username
	return fr.redis.LPush(ctx, key, tweet).Err()
}

func (fr *FeedRepository) Get(username string) ([]string, error) {
	ctx := context.Background()
	key := "feed_" + username

	feedList, err := fr.redis.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return feedList, nil
}