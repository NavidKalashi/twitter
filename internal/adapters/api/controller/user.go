package controller

import (
	"github.com/NavidKalashi/twitter/internal/core/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}