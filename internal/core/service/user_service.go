package service

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
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