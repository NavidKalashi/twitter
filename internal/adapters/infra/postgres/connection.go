package database

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // Import the PostgreSQL driver
    "github.com/NavidKalashi/twitter/internal/config"
)

type DB struct {
    Conn *sql.DB
}

func InitDB(cfg *config.Config) (*DB, error) {
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    log.Println("Database connection established")
    return &DB{Conn: db}, nil
}