package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id        uuid.UUID `gorm:"type:uuid;primary key auto_increment:false"`
	username  string    `gorm:"type:varchar(255);not null;unique"`
	name      string    `gorm:"type:varchar(255);not null"`
	email     string    `gorm:"type:varchar(255);not null;unique"`
	password  string    `gorm:"type:varchar(255);not null"`
	bio       string    `gorm:"type:varchar(255)"`
	birthday  time.Time `gorm:"type:date"`
	createdAt time.Time `gorm:"type:timestamp"`
	updatedAt time.Time `gorm:"type:timestamp"`
}