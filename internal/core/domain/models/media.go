package models

import (
	"time"

	"gorm.io/gorm"
)

type Media struct {
	ID        string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TweetID   string         `gorm:"not null" json:"tweet_id"`
	Type      string         `gorm:"type:varchar(50);not null" json:"type"`
	FileName  string         `gorm:"type:varchar(255);not null" json:"file_name"`
	FileURL   string         `gorm:"type:varchar(255);not null" json:"file_url"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
