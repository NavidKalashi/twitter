package api

import (
	"github.com/NavidKalashi/twitter/internal/adapters/api/controller"
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
	server.AddRoutes(userController)
	return server
}

func (s *Server) AddRoutes(userController *controller.UserController) {
	s.engine.POST("/users", userController.ControllerCreateUser)
}

func (s *Server) Start() {
	s.engine.SetTrustedProxies([]string{"127.0.0.1"})
	s.engine.Run(":8080")
}
