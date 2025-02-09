package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(ctx *gin.Context) {

}

func (s *UserService) GetUser(ctx *gin.Context) {
	
}

func (s *UserService) UpdateUser(ctx *gin.Context) {

}

func (s *UserService) DeleteUser(ctx *gin.Context) {

}

func (s *UserService) GetUsers(ctx *gin.Context) {

}