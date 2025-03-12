package api

import (
	"github.com/NavidKalashi/twitter/internal/adapters/api/controller"
	"github.com/NavidKalashi/twitter/internal/adapters/api/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	userController *controller.UserController
	engine         *gin.Engine
}

func NewServer(engine *gin.Engine, userController *controller.UserController) *Server {
	server := &Server{
		userController: userController,
		engine: engine,
	}
	server.AddRoutes(userController)
	return server
}

func (s *Server) AddRoutes(userController *controller.UserController) {
	authRoutes := s.engine.Group("/protected")
	s.engine.POST("/register", userController.RegisterController)
	s.engine.POST("/verify-email", userController.VerifyController)
	s.engine.POST("/refresh", userController.RefreshController)
	s.engine.POST("/send-code-again", userController.ResendController)
	s.engine.POST("/login", userController.LoginController)
	authRoutes.Use(middleware.AuthMiddleware())
	{
		authRoutes.GET("/profile", userController.GetController)
		authRoutes.DELETE("/logout", userController.LogoutController)
		authRoutes.PUT("/edit", userController.EditController)
	}
}

func (s *Server) Start() {
	s.engine.SetTrustedProxies([]string{"127.0.0.1"})
	s.engine.Run(":8080")
}
