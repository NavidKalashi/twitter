package repository

import (
	"errors"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) ports.UserRepository {
	return &UserGormRepository{db: db}
}

func (r *UserGormRepository) CreateUser(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *UserGormRepository) GetUser(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error
	err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserGormRepository) UpdateUser(user *models.User) error {
	result := r.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no user updated")
	}
	return nil
}

func (r *UserGormRepository) DeleteUser(id uuid.UUID) error {
	var user models.User
	var otp models.OTP
	r.db.Where("user_id = ?", id).Delete(&otp)
	result := r.db.Delete(&user, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}