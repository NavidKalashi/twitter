package controller

import (
	"github.com/NavidKalashi/twitter/internal/core/service"
	"github.com/gin-gonic/gin"
)

type TweetController struct {
	tweetService *service.TweetService
}

func NewTweetController(tweetService *service.TweetService) *TweetController {
	return &TweetController{tweetService: tweetService}
}

func (c *TweetController) CreateTweet(ctx *gin.Context) {

}

func (c *TweetController) GetTweets(ctx *gin.Context) {

}

func (c *TweetController) GetTweet(ctx *gin.Context) {

}

func (c *TweetController) UpdateTweet(ctx *gin.Context) {

}

func (c *TweetController) DeleteTweet(ctx *gin.Context) {

}