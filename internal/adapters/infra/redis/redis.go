package redis

import (
	"context"
	"fmt"

	"github.com/NavidKalashi/twitter/internal/config"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rd *redis.Client
}

var ctx = context.Background()

func InitRedis(cfg *config.Config) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println("Redis connected", pong)
	return &Redis{rd: client}, nil
}

func (r *Redis) GetRedis() *redis.Client {
	return r.rd
}
