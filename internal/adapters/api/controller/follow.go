package controller

import (
	"net/http"

	"github.com/NavidKalashi/twitter/internal/core/service"
	"github.com/gin-gonic/gin"
)

type FollowController struct {
	followService *service.FollowService
}

func NewFollowController(followService *service.FollowService) *FollowController {
	return &FollowController{followService: followService}
}

func (fc *FollowController) FollowingController(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found in context"})
		return
	}

	followingName := c.Param("following_name")

	usernameStr, ok := username.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username format"})
		return
	}

	err := fc.followService.FollowUser(usernameStr, followingName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "followed"})
}