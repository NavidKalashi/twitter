package controller

import (
	"net/http"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) CreateUserController(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user)
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.userService.CreateUser(&user)
	err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uc *UserController) GetUserController(c *gin.Context) {
	userIDStr := c.Param("id")
    userID, err := uuid.Parse(userIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
        return
    }

    user, err := uc.userService.GetUser(userID)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}

func (uc *UserController) UpdateUserController(c *gin.Context) {
    var user models.User

    if err := c.ShouldBindJSON(&user) 
    err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request users"})
    }

    err := uc.userService.UpdateUser(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    }

    c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

func (uc *UserController) DeleteUserController(c *gin.Context) {
	userIDStr := c.Param("id")
    userID, err := uuid.Parse(userIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
        return
    }

    err = uc.userService.DeleteUser(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}