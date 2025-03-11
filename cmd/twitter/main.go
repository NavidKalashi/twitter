package main

import (
	"log"

	"github.com/NavidKalashi/twitter/internal/adapters/api"
	"github.com/NavidKalashi/twitter/internal/adapters/api/controller"
	"github.com/NavidKalashi/twitter/internal/adapters/infra/postgres"
	"github.com/NavidKalashi/twitter/internal/adapters/infra/redis"
	"github.com/NavidKalashi/twitter/internal/adapters/repository"
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

	// otp
	otpRepository := repository.NewOTPRepository(redis.GetRedis())
	
	// refresh token
	refreshTokenRepository := repository.NewRefreshTokenRepository(db.GetDB())

	// user
	userRepository := repository.NewUserRepository(db.GetDB())
	userService := service.NewUserService(userRepository, otpRepository, refreshTokenRepository)
	userController := controller.NewUserController(userService)

	r := gin.Default()

	server := api.NewServer(r, userController)
	server.Start()
}
