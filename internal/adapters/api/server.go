package api

import (
	"github.com/NavidKalashi/twitter/internal/adapters/api/controller"
	"github.com/NavidKalashi/twitter/internal/adapters/api/middleware"
	"github.com/gin-gonic/gin"
)

var engine = gin.Default()

type Server struct {
	tweetController *controller.TweetController
}

func NewServer(tweetController *controller.TweetController) *Server {
	server := &Server{tweetController: tweetController}
	server.AddRoutes(tweetController)
	return server
}

func (s *Server) AddRoutes(tweetController *controller.TweetController) {
	engine.Use(middleware.AuthMiddleware())
	engine.POST("/tweets", s.tweetController.CreateTweet)
	engine.GET("/tweets", s.tweetController.GetTweets)
	engine.GET("/tweets/:id", s.tweetController.GetTweet)
	engine.PUT("/tweets/:id", s.tweetController.UpdateTweet)
	engine.DELETE("/tweets/:id", s.tweetController.DeleteTweet)
}

func (s *Server) Start() {
	engine.Run(":8080")
}
