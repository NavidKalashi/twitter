package main

import (
    "fmt"
    "github.com/NavidKalashi/twitter/internal/config"
    "github.com/NavidKalashi/twitter/internal/adapters/repository"
)

func main() {
    config.LoadConfig()
    repository.InitDB()
    fmt.Println("Twitter:", config.Cfg.Twitter)
	fmt.Println("Database Host:", config.Cfg.DB.Host)
}