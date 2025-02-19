package models

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OTP struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	Code      uint      `gorm:"not null" json:"code"`
	User      User      `gorm:"foreignKey:UserID;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID    uuid.UUID `json:"user_id"`
	Verified  bool      `gorm:"default:false" json:"verified"`
}

func generateOTP() uint {
	rand.Seed(time.Now().UnixNano())
	return uint(rand.Intn(900000) + 100000)
}

func (u *OTP) BeforeCreate(tx *gorm.DB) (err error) {
	u.Code = generateOTP()
	return nil
}
