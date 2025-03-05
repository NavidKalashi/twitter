package models

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
	Value     string    `gorm:"not null" json:"code"`
	User      User      `gorm:"foreignKey:UserID;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID    uuid.UUID `json:"user_id"`
}
