package main

import (
	"log"
    "fmt"

	"github.com/NavidKalashi/twitter/internal/adapters/repository"
	"github.com/NavidKalashi/twitter/internal/config"
    "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func main() {
	cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    db, err := InitDB(cfg)
    if err != nil {
        log.Fatalf("failed to initialize database: %v", err)
    }
    
    repository.NewRepository(db)

    fmt.Println("App Name:", cfg.Twitter)
    fmt.Println("Database Host:", cfg.DB.Host)
    fmt.Println("Port:", cfg.Port)
}