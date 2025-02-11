package controller

import (
	"github.com/NavidKalashi/twitter/internal/core/service"
)

type TweetController struct {
	tweetService *service.TweetService
}

func NewTweetController(tweetService *service.TweetService) *TweetController {
	return &TweetController{tweetService: tweetService}
}