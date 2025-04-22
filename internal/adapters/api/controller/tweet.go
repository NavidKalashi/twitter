package controller

import (
	"net/http"

	"github.com/NavidKalashi/twitter/internal/core/domain/events"
	"github.com/NavidKalashi/twitter/internal/core/service"
	"github.com/gin-gonic/gin"
)

type TweetController struct {
	tweetService   *service.TweetService
	produceService *service.ProduceService
}

func NewTweetController(tweetService *service.TweetService, produceService *service.ProduceService) *TweetController {
	return &TweetController{tweetService: tweetService, produceService: produceService}
}

func (tc *TweetController) CreateController(c *gin.Context) {
	var feed events.Feed

	if err := c.BindJSON(&feed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
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

	feed.Username = usernameStr

	err := tc.produceService.ProduceFeed(feed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tweet created"})
}

func (tc *TweetController) GetController(c *gin.Context) {
	tweets, err := tc.tweetService.GetTweets()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tweets not found"})
	}

	c.JSON(http.StatusOK, gin.H{"tweets": tweets})
}

func (tc *TweetController) DeleteAllController(c *gin.Context) {
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

	err := tc.tweetService.DeleteAll(usernameStr)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "tweet not deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tweets deleted successfully"})
}

func (tc *TweetController) DeleteController(c *gin.Context) {
	tweetID := c.Param("id")

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

	err := tc.tweetService.Delete(usernameStr, tweetID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "tweet not deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tweet deleted"})
}
