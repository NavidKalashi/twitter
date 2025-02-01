package main

import (
	"log"
    "fmt"

	"github.com/NavidKalashi/twitter/internal/config"
    "github.com/NavidKalashi/twitter/internal/infra/database"
)

func main() {
	cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    db, err := database.InitDB(cfg)
    if err != nil {
        log.Fatalf("failed to initialize database: %v", err)
    }
    
    fmt.Println(db)
    fmt.Println("App Name:", cfg.Twitter)
    fmt.Println("Database Host:", cfg.DB.Host)
    fmt.Println("Port:", cfg.Port)
}