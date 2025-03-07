package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/NavidKalashi/twitter/internal/core/ports"
	"github.com/redis/go-redis/v9"
)

type AccessTokenRepository struct {
	redis *redis.Client
}

func NewAccessTokenRepository(redis *redis.Client) ports.AccessToken {
	return &AccessTokenRepository{redis: redis}
}

func (r *AccessTokenRepository) Set(userID string, token string) error {
	ctx := context.Background()
	key := fmt.Sprintf("access_token_%s", userID)
	return r.redis.Set(ctx, key, token, 15 * time.Minute).Err()
}

func (r *AccessTokenRepository) Delete(userID string) error {
	ctx := context.Background()
	key := fmt.Sprintf("access_token_%s", userID)
	return r.redis.Del(ctx, key).Err()
}