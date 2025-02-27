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

func (r *OTPRepository) FindByUserID(userID string) (*models.OTP, error) {
    var otp models.OTP
    if err := r.db.Where("user_id = ?", userID).Last(&otp).Error
	err != nil {
        return nil, err
    }
    return &otp, nil
}

func (repo *OTPRepository) Verified(otp *models.OTP) error {
    return repo.db.Save(&otp).Error
}