package api

import (
	"github.com/NavidKalashi/twitter/internal/adapters/api/controller"
	"github.com/NavidKalashi/twitter/internal/adapters/api/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	userController    *controller.UserController
	tweetController   *controller.TweetController
	gestureController *controller.GestureControlelr
	followController  *controller.FollowController
	engine            *gin.Engine
}

func NewServer(engine *gin.Engine, userController *controller.UserController, tweetController *controller.TweetController, gestureController *controller.GestureControlelr, followController *controller.FollowController) *Server {
	server := &Server{
		userController:    userController,
		tweetController:   tweetController,
		gestureController: gestureController,
		followController:  followController,
		engine:            engine,
	}
	server.AddRoutes(userController, tweetController, gestureController, followController)
	return server
}

func (s *Server) AddRoutes(userController *controller.UserController, tweetController *controller.TweetController, gestureController *controller.GestureControlelr, followController *controller.FollowController) {
	authRoutes := s.engine.Group("/protected")
	s.engine.POST("/auth/register", userController.RegisterController)
	s.engine.POST("/auth/verify-email", userController.VerifyController)
	s.engine.POST("/auth/refresh", userController.RefreshController)
	s.engine.POST("/auth/resend-code", userController.ResendController)
	s.engine.POST("/auth/login", userController.LoginController)
	s.engine.POST("/users/search", userController.SearchController)
	s.engine.GET("/tweets", tweetController.GetController)
	
	authRoutes.Use(middleware.AuthMiddleware())
	{
		authRoutes.GET("/users/me", userController.GetController)
		authRoutes.PUT("/users/me", userController.EditController)
		authRoutes.DELETE("/auth/logout", userController.LogoutController)
	
		authRoutes.POST("/tweets", tweetController.CreateController)
		authRoutes.DELETE("/tweets", tweetController.DeleteAllController)
		authRoutes.DELETE("/tweets/:id", tweetController.DeleteController)
	
		authRoutes.POST("/tweets/:tweet_id/gestures", gestureController.AddViewController)
	
		authRoutes.POST("/users/:username/follow", followController.FollowingController)
		authRoutes.DELETE("/users/:username/follow", followController.UnfollowController)
		authRoutes.GET("/users/me/followers", followController.GetFollowersController)
	}	
}

func (s *Server) Start() {
	s.engine.SetTrustedProxies([]string{"127.0.0.1"})
	s.engine.Run(":8080")
}
