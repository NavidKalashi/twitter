package models

import (
	"time"

	"github.com/google/uuid"
)

type OTP struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	Code      uint      `gorm:"not null" json:"code"`
	User      User      `gorm:"foreignKey:UserID;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID    uuid.UUID `json:"user_id"`
	Verified  bool      `gorm:"default:false" json:"verified"`
}