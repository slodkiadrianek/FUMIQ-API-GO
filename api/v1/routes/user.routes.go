package routes

import (
	"FUMIQ_API/api/v1/controllers"
	"FUMIQ_API/middleware"
	"FUMIQ_API/schemas"

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
		userGroup.GET("/:userId", u.AuthMiddleware.Verify, middleware.ValidateRequestData[*schemas.UserId]("params"), u.UserController.GetUser)
		userGroup.PATCH("/:userId", middleware.ValidateRequestData[*schemas.UserId]("params"), middleware.ValidateRequestData[*schemas.ChangePassword]("body"), u.AuthMiddleware.Verify, u.UserController.ChangePassword)
		userGroup.PUT("/:userId", middleware.ValidateRequestData[*schemas.UserId]("params"), middleware.ValidateRequestData[*schemas.UpdateUser]("body"), u.UserController.UpdateUser)
		userGroup.DELETE("/:userId", middleware.ValidateRequestData[*schemas.UserId]("params"), middleware.ValidateRequestData[*schemas.DeleteUser]("body"), u.AuthMiddleware.Verify, u.UserController.DeleteUser)
		userGroup.GET("/:userId/quizzes", middleware.ValidateRequestData[*schemas.UserId]("params"))
		userGroup.POST("/:userId/sessions", middleware.ValidateRequestData[*schemas.UserId]("params"), middleware.ValidateRequestData[*schemas.SessionId]("params"))
		userGroup.GET("/:userId/sessions/:sessionId", middleware.ValidateRequestData[*schemas.UserId]("params"), middleware.ValidateRequestData[*schemas.SessionId]("params"))
		userGroup.PATCH("/:userId/sessions/:sessionId", middleware.ValidateRequestData[*schemas.UserId]("params"), middleware.ValidateRequestData[*schemas.SessionId]("params"))
		userGroup.GET("/:userId/sessions/:sessionId/results", middleware.ValidateRequestData[*schemas.UserId]("params"), middleware.ValidateRequestData[*schemas.SessionId]("params"))

	}
}
