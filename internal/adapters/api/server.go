package api

import (
	"github.com/NavidKalashi/twitter/internal/adapters/api/controller"
	"github.com/NavidKalashi/twitter/internal/adapters/api/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	tweetController *controller.TweetController
	userController  *controller.UserController
	engine          *gin.Engine
}

func NewServer(tweetController *controller.TweetController, userController *controller.UserController) *Server {
	server := &Server{
		tweetController: tweetController,
		userController:  userController,
	}
	server.engine = gin.Default()
	server.AddRoutes(tweetController, userController)
	return server
}

func (s *Server) AddRoutes(tweetController *controller.TweetController, userController *controller.UserController) {
	s.engine.Use(middleware.AuthMiddleware())
	s.engine.POST("/tweets", tweetController.CreateTweet)
	s.engine.GET("/tweets", tweetController.GetTweets)
	s.engine.GET("/tweets/:id", tweetController.GetTweet)
	s.engine.PUT("/tweets/:id", tweetController.UpdateTweet)
	s.engine.DELETE("/tweets/:id", tweetController.DeleteTweet)

	s.engine.POST("/user", userController.CreateUser)
	s.engine.GET("/user", userController.GetUser)
	s.engine.GET("/user/:id", userController.GetUser)
	s.engine.PUT("/user/:id", userController.UpdateUser)
	s.engine.DELETE("/user/:id", userController.DeleteUser)
}

func (s *Server) Start() {
	s.engine.SetTrustedProxies([]string{"127.0.0.1"})
	s.engine.Run(":8080")
}
