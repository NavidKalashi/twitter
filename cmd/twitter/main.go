package main

import (
	"log"

	"github.com/NavidKalashi/twitter/internal/adapters/api"
	"github.com/NavidKalashi/twitter/internal/adapters/api/controller"
	"github.com/NavidKalashi/twitter/internal/adapters/infra/postgres"
	"github.com/NavidKalashi/twitter/internal/adapters/repository"
	"github.com/NavidKalashi/twitter/internal/config"
	"github.com/NavidKalashi/twitter/internal/core/service"
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

	// otp
	otpRepository := repository.NewOTPRepository(db.GetDB())

	// user
	userRepository := repository.NewUserRepository(db.GetDB())
	userService := service.NewUserService(userRepository, otpRepository)
	userController := controller.NewUserController(userService)

	server := api.NewServer(userController)
	server.Start()
}
