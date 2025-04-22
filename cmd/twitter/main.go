package main

import (
	"log"

	"github.com/NavidKalashi/twitter/internal/adapters/api"
	"github.com/NavidKalashi/twitter/internal/adapters/api/controller"
	"github.com/NavidKalashi/twitter/internal/adapters/infra/minio"
	"github.com/NavidKalashi/twitter/internal/adapters/infra/postgres"
	"github.com/NavidKalashi/twitter/internal/adapters/infra/rabbitmq"
	"github.com/NavidKalashi/twitter/internal/adapters/infra/redis"
	"github.com/NavidKalashi/twitter/internal/adapters/messaging"
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

	rabbit, err := rabbitmq.InitRabbitMQ(cfg)
	if err != nil {
		log.Fatalf("failed to initialize rabbitmq: %v", err)
	}
	defer func() {
		if err := rabbit.Close(); err != nil {
			log.Printf("failed to close rabbitmq: %v:", err)
		}
	}()

	// otp
	otpRepository := repository.NewOTPRepository(redis.GetRedis())

	// feed
	feedRepository := repository.NewFeedRepository(redis.GetRedis())

	// media
	mediaStorage := storage.NewMinioStorage(minio.GetMinio())
	mediaRepository := repository.NewMediaRepository(db.GetDB())

	// refresh token
	refreshTokenRepository := repository.NewRefreshTokenRepository(db.GetDB())

	// producer rabbitmq
	producerMessaging := messaging.NewRabbitMQProducer(rabbit.GetRabbit())
	produceService := service.NewProduceService(producerMessaging)

	// user
	userRepository := repository.NewUserRepository(db.GetDB())
	userService := service.NewUserService(userRepository, otpRepository, refreshTokenRepository)
	userController := controller.NewUserController(userService)

	// tweet
	tweetRepository := repository.NewTweetRepository(db.GetDB())
	tweetService := service.NewTweetService(tweetRepository, mediaRepository, mediaStorage)
	tweetController := controller.NewTweetController(tweetService, produceService)

	// gesture
	gestureRepository := repository.NewGestureRepository(db.GetDB(), redis.GetRedis())
	gestureService := service.NewGestureService(gestureRepository, tweetRepository)
	gestureController := controller.NewGestureService(gestureService, produceService)

	// follow
	followRepository := repository.NewFollowRepository(db.GetDB())
	followService := service.NewFollowService(followRepository, userRepository, tweetRepository, feedRepository)
	followController := controller.NewFollowController(followService)

	// consumer rabbitmq
	rabbitConsumer := messaging.NewRabbitMQConsumer(rabbit.GetRabbit())
	consumerService := service.NewConsumeService(rabbitConsumer, feedRepository, followRepository, gestureRepository, *tweetService, *gestureService)

	// feed 
	go func() {
		if err := consumerService.ConsumeFeed(); err != nil {
			log.Printf("error consuming feed events: %v", err)
		}
	}()

	// gesture
	go func() {
		if err := consumerService.ConsumeGesture(); err != nil {
			log.Printf("error consuming gesture events: %v", err)
		}
	}()

	r := gin.Default()

	server := api.NewServer(r, userController, tweetController, gestureController, followController)
	server.Start()
}
