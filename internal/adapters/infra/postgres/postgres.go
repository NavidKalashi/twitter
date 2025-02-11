package postgres

import (
	"fmt"
	"log"
	
	"github.com/NavidKalashi/twitter/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
)

type DB struct {
	Conn *gorm.DB
}

func InitDB(cfg *config.Config) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }

	if err := sqlDB.Ping(); err != nil {
        return nil, err
    }

	log.Println("Database connection established")
	return &DB{Conn: db}, nil
}

func (db *DB) GetDB() *gorm.DB {
	return db.Conn
}
