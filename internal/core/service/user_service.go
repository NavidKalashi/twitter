package service

import (
	"github.com/NavidKalashi/twitter/internal/adapters/infra/repository"
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
)

type UserService struct {
	repo *repository.UserGormRepository
}

func NewUserService(repo *repository.UserGormRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.RepoCreateUser(user)
}