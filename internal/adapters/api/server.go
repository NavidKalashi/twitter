package api

import (
	"github.com/NavidKalashi/twitter/internal/adapters/api/controller"
	"github.com/gin-gonic/gin"
)

type Server struct {
	userController  *controller.UserController
	engine          *gin.Engine
}

func NewServer(userController *controller.UserController) *Server {
	server := &Server{
		userController:  userController,
	}
	server.engine = gin.Default()
	server.AddRoutes(userController)
	return server
}

func (s *Server) AddRoutes(userController *controller.UserController) {
	s.engine.POST("/user", userController.CreateUserController)
	s.engine.GET("/user/:id", userController.GetUserController)
	s.engine.DELETE("/user/:id", userController.DeleteUserController)
	s.engine.PUT("/user/update", userController.UpdateUserController)

}

func (s *Server) Start() {
	s.engine.SetTrustedProxies([]string{"127.0.0.1"})
	s.engine.Run(":8080")
}
