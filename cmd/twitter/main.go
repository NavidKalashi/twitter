package main

import (
    "log"

	"github.com/NavidKalashi/twitter/internal/adapters/repository"
	"github.com/NavidKalashi/twitter/internal/config"
    "github.com/NavidKalashi/twitter/internal/infra/database"
    "github.com/NavidKalashi/twitter/internal/infra/server"
)

func main() {
    cfg, err := config.LoadConfig()
    if (err != nil) {
        log.Fatalf("Failed to load config: %v", err)
    }

    db, err := database.InitDB(cfg)
    if (err != nil) {
        log.Fatalf("failed to initialize database: %v", err)
    }
    
    repository.NewRepository(db)

    fmt.Println("App Name:", cfg.Twitter)
    fmt.Println("Database Host:", cfg.DB.Host)
    fmt.Println("Port:", cfg.Port)
}