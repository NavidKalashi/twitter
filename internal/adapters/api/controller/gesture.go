package controller

import (
	"net/http"

	"github.com/NavidKalashi/twitter/internal/core/service"
	"github.com/gin-gonic/gin"
)

type GestureControlelr struct {
	gestureService *service.GestureService
}

func NewGestureService(gestureService *service.GestureService) *GestureControlelr {
	return &GestureControlelr{gestureService: gestureService}
}

func (gc *GestureControlelr) AddViewController(c *gin.Context) {
	tweetID := c.Param("tweet_id")

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found in context"})
		return
	}

	usernameStr, ok := username.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username format"})
		return
	}

	err := gc.gestureService.AddView(tweetID, usernameStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "view added"})
}

func (gc *GestureControlelr) AddLikeController(c *gin.Context) {
	tweetID := c.Param("tweet_id")

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found in context"})
		return
	}

	usernameStr, ok := username.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username format"})
		return
	}

	err := gc.gestureService.AddLike(tweetID, usernameStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "like added"})
}

func (gc *GestureControlelr) AddRetweetController(c *gin.Context) {
	tweetID := c.Param("tweet_id")

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found in context"})
		return
	}

	usernameStr, ok := username.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username format"})
		return
	}

	err := gc.gestureService.AddRetweet(tweetID, usernameStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "retweet added"})
}