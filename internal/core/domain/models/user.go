package models

import (
    "time"

    "github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
    Username  string    `gorm:"type:varchar(255);not null;unique" json:"username"`
    Name      string    `gorm:"type:varchar(255);not null" json:"name"`
    Email     string    `gorm:"type:varchar(255);not null;unique" json:"email"`
    Password  string    `gorm:"type:varchar(255);not null" json:"password"`
    Bio       string    `gorm:"type:varchar(255)" json:"bio"`
    Birthday  time.Time `gorm:"type:date" json:"birthday"`
    CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
    UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}