package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/resend/resend-go/v2"
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
	OTP       []OTP
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	otp := OTP{
		UserID: u.ID,
		Code:   generateOTP(),
	}

	err = tx.Create(&otp).Error

	apiKey := "re_MzT3hsTe_F4JTUCmausSEMkipw2tC7QwT"
    client := resend.NewClient(apiKey)

    params := &resend.SendEmailRequest{
        From:    "Acme <onboarding@resend.dev>",
        To:       []string{"kalashinavid@gmail.com"},
        Html:    "<strong>your code is: </strong>" + fmt.Sprintf("%d", otp.Code),
        Subject: "Verfiy code",
    }

    client.Emails.Send(params)
	
	return err
}