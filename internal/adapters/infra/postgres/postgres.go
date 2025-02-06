package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/NavidKalashi/twitter/internal/config"
	_ "github.com/lib/pq"
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

func (db *DB) GetDB() *sql.DB {
	return db.Conn
}
