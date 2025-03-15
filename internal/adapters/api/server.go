package api

import (
	"github.com/NavidKalashi/twitter/internal/adapters/api/controller"
	"github.com/NavidKalashi/twitter/internal/adapters/api/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	userController *controller.UserController
	tweetController *controller.TweetController
	engine         *gin.Engine
}

func NewServer(engine *gin.Engine, userController *controller.UserController, tweetController *controller.TweetController) *Server {
	server := &Server{
		userController: userController,
		tweetController: tweetController,
		engine: engine,
	}
	server.AddRoutes(userController, tweetController)
	return server
}

func (s *Server) AddRoutes(userController *controller.UserController, tweetController *controller.TweetController) {
	authRoutes := s.engine.Group("/protected")
	s.engine.POST("/register", userController.RegisterController)
	s.engine.POST("/verify-email", userController.VerifyController)
	s.engine.POST("/refresh", userController.RefreshController)
	s.engine.POST("/send-code-again", userController.ResendController)
	s.engine.POST("/login", userController.LoginController)
	s.engine.GET("/tweets", tweetController.GetController)
	authRoutes.Use(middleware.AuthMiddleware())
	{
		authRoutes.GET("/profile", userController.GetController)
		authRoutes.DELETE("/logout", userController.LogoutController)
		authRoutes.PUT("/edit", userController.EditController)
		authRoutes.POST("/create-tweet", tweetController.CreateController)
		authRoutes.DELETE("/delete-all-tweet", tweetController.DeleteAllController)
		authRoutes.DELETE("/delete-tweet/:id", tweetController.DeleteController)
	}
}

func (s *Server) Start() {
	s.engine.SetTrustedProxies([]string{"127.0.0.1"})
	s.engine.Run(":8080")
}
