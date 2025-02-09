package controller

import (
	"github.com/NavidKalashi/twitter/internal/core/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(ctx *gin.Context) {

}

func (c *UserController) GetUsers(ctx *gin.Context) {
	c.userService.GetUsers(ctx)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	c.userService.GetUser(ctx)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	c.userService.UpdateUser(ctx)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	c.userService.DeleteUser(ctx)
}