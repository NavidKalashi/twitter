package models

import (
	"time"

	"gorm.io/gorm"
)

type Gesture struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	TweetID   string         `gorm:"not null" json:"tweet_id"`
	Username  string         `gorm:"type:varchar(255);not null" json:"username"`
	Type      string         `gorm:"type:varchar(50);not null" json:"type"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
