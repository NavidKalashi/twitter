package repository

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	"gorm.io/gorm"
)

type OTPRepository struct {
	db *gorm.DB
}

func NewOTPRepository(db *gorm.DB) ports.OTPRepository{
	return &OTPRepository{db: db}
}

func (or *OTPRepository) Create(user *models.User, code uint) error {
	otp := models.OTP{
		UserID: user.ID,
		Code:   code,
	}
	return or.db.Create(&otp).Error
}