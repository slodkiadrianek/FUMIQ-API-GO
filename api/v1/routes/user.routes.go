package routes

import (
	"FUMIQ_API/api/v1/controllers"
	"FUMIQ_API/middleware"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	UserController *controllers.UserController
	AuthMiddleware *middleware.AuthMiddleware
}

func NewUserRoutes(userController *controllers.UserController, authMiddleware *middleware.AuthMiddleware) *UserRoutes {
	return &UserRoutes{
		UserController: userController,
		AuthMiddleware: authMiddleware,
	}
}

func (u *UserRoutes) SetupUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/:userId", u.AuthMiddleware.Verify, u.UserController.GetUser)
		userGroup.PATCH("/:userId", u.AuthMiddleware.Verify, u.UserController.ChangePassword)
		userGroup.PUT("/:userId")
		userGroup.DELETE("/:userId")
		userGroup.GET("/:userId/quizzes")
		userGroup.POST("/:userId/sessions")
		userGroup.GET("/:userId/sessions/:sessionId")
		userGroup.PATCH("/:userId/sessions/:sessionId")
		userGroup.GET("/:userId/sessions/:sessionId/results")

	}
}
