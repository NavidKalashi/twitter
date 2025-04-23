package controller

import (
	"net/http"

	"github.com/NavidKalashi/twitter/internal/core/domain/events"
	"github.com/NavidKalashi/twitter/internal/core/service"
	"github.com/gin-gonic/gin"
)

type GestureControlelr struct {
	gestureService *service.GestureService
	ProduceService *service.ProduceService
}

func NewGestureService(gestureService *service.GestureService, ProduceService *service.ProduceService) *GestureControlelr {
	return &GestureControlelr{gestureService: gestureService, ProduceService: ProduceService}
}

func (gc *GestureControlelr) AddViewController(c *gin.Context) {
	var gesture events.Gesture

	if err := c.BindJSON(&gesture); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

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

	gesture.Username = usernameStr
	gesture.TweetID = tweetID

	err := gc.ProduceService.ProducerGesture(gesture)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "gesture added"})
}
