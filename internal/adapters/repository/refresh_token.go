package repository

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) ports.RefreshToken {
	return &RefreshTokenRepository{db: db}
}

func (rt *RefreshTokenRepository) Create(userID uuid.UUID, refreshToken string) error {
	token := models.RefreshToken{
		UserID: userID,
		Value:  refreshToken,
	}
	return rt.db.Create(&token).Error
}

func (rt *RefreshTokenRepository) Delete(userID uuid.UUID) error {
	var refreshToken models.RefreshToken

	result := rt.db.Where("user_id = ?", userID).Delete(&refreshToken)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
