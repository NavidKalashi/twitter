package service

import (
	"github.com/NavidKalashi/twitter/internal/adapters/infra/repository"
)

type UserService struct {
	repo *repository.UserGormRepository
}

func NewUserService(repo *repository.UserGormRepository) *UserService {
	return &UserService{repo: repo}
}