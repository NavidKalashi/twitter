package service

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
)

type UserService struct {
	userRepository *ports.UserGormRepository
}

func NewUserService(userRepository *ports.UserGormRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepository.RepoCreateUser(user)
}