package api

import (
	"github.com/NavidKalashi/twitter/internal/adapters/api/controller"
	"github.com/NavidKalashi/twitter/internal/adapters/api/middleware"
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
	s.engine.POST("/register", userController.RegisterController)
	s.engine.POST("/verify-email/:id", userController.VerifyController)
	s.engine.POST("/refresh", userController.RefreshController)
	s.engine.POST("/send-code-again/:id", userController.ResendController)
	s.engine.GET("/users", middleware.AuthMiddleware())
	s.engine.GET("/user/:id", userController.GetController)
	s.engine.DELETE("/user/:id", userController.UpdateController)
	s.engine.PUT("/user/update", userController.DeleteController)
}

func (s *Server) Start() {
	s.engine.SetTrustedProxies([]string{"127.0.0.1"})
	s.engine.Run(":8080")
}
