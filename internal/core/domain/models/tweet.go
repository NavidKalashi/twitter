package models

import (
	"time"
)

type Tweet struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Text      string    `gorm:"type:varchar(255)" json:"text"`
	CreatedBy string    `gorm:"type:varchar(255);not null" json:"created_by"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	DeletedAt time.Time `gorm:"type:timestamp" json:"deleted_at,omitempty"`
	Media     []Media   `gorm:"foreignKey:TweetID;references:ID"`
	Gesture   []Gesture `gorm:"foreignKey:TweetID;references:ID"`
}
