package repository

import (
    "fmt"
    "log"

    "github.com/NavidKalashi/twitter/internal/config"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
        config.Cfg.DB.Host, config.Cfg.DB.User, config.Cfg.DB.Password, config.Cfg.DB.Name, config.Cfg.DB.Port)
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Debug statement to confirm database connection
    fmt.Println("Database connection established")
}