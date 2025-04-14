package controller

import (
	"encoding/json"
	"net/http"

	"github.com/NavidKalashi/twitter/internal/core/service"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type GestureControlelr struct {
	gestureService *service.GestureService
	channel        *amqp.Channel
}

func NewGestureService(gestureService *service.GestureService, channel *amqp.Channel) *GestureControlelr {
	return &GestureControlelr{gestureService: gestureService, channel: channel}
}

func (gc *GestureControlelr) AddViewController(c *gin.Context) {
	var gesture struct {
		TweetID     string   `json:"tweetID"`
		TypeStr     string   `json:"typeStr"`
	}

	if err := c.BindJSON(&gesture); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	
	body, err := json.Marshal(gesture)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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

	err = gc.channel.Publish(
		"feed_exchange",
		"", // routing key for fanout
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = gc.gestureService.AddView(gesture.TweetID, usernameStr, gesture.TypeStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "view added"})
}
