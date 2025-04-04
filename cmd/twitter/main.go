package main

import (
	"log"

	"github.com/NavidKalashi/twitter/internal/adapters/api"
	"github.com/NavidKalashi/twitter/internal/adapters/api/controller"
	"github.com/NavidKalashi/twitter/internal/adapters/infra/minio"
	"github.com/NavidKalashi/twitter/internal/adapters/infra/postgres"
	"github.com/NavidKalashi/twitter/internal/adapters/infra/redis"
	"github.com/NavidKalashi/twitter/internal/adapters/repository"
	"github.com/NavidKalashi/twitter/internal/adapters/storage"
	"github.com/NavidKalashi/twitter/internal/config"
	"github.com/NavidKalashi/twitter/internal/core/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := postgres.InitDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	redis, err := redis.InitRedis(cfg)
	if err != nil {
		log.Fatalf("failed to initialize redis: %v", err)
	}

	minio, err := minio.InitMinio(cfg)
	if err != nil {
		log.Fatalf("failed to initialize minio: %v", err)
	}

	// otp
	otpRepository := repository.NewOTPRepository(redis.GetRedis())

	// media
	mediaStorage := storage.NewMinioStorage(minio.GetMinio())
	mediaRepository := repository.NewMediaRepository(db.GetDB())

	// refresh token
	refreshTokenRepository := repository.NewRefreshTokenRepository(db.GetDB())

	// user
	userRepository := repository.NewUserRepository(db.GetDB())
	userService := service.NewUserService(userRepository, otpRepository, refreshTokenRepository)
	userController := controller.NewUserController(userService)

	// tweet
	tweetRepository := repository.NewTweetRepository(db.GetDB())
	tweetService := service.NewTweetService(tweetRepository, mediaRepository, mediaStorage)
	tweetController := controller.NewTweetController(tweetService)

	// gesture
	gestureRepository := repository.NewGestureRepository(db.GetDB())
	gestureService := service.NewGestureService(gestureRepository, tweetRepository)
	gestureController := controller.NewGestureService(gestureService)

	r := gin.Default()

	server := api.NewServer(r, userController, tweetController, gestureController)
	server.Start()
}
