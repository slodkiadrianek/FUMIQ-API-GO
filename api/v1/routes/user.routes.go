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
		userGroup.GET("/:userId", u.AuthMiddleware.Verify, u.UserController.GetUser)
		userGroup.PATCH("/:userId", middleware.ValidateRequestData[*schemas.UserId]("params"), middleware.ValidateRequestData[*schemas.ChangePassword]("body"), u.AuthMiddleware.Verify, u.UserController.ChangePassword)
		userGroup.PUT("/:userId", middleware.ValidateRequestData[*schemas.UserId]("params"), middleware.ValidateRequestData[*schemas.UpdateUser]("body"), u.UserController.UpdateUser)
		userGroup.DELETE("/:userId", middleware.ValidateRequestData[*schemas.UserId]("params"), middleware.ValidateRequestData[*schemas.DeleteUser]("body"), u.AuthMiddleware.Verify, u.UserController.DeleteUser)
		userGroup.POST("/:userId/sessions", u.AuthMiddleware.Verify, middleware.ValidateRequestData[*schemas.JoinQuiz]("body"), u.UserController.JoinSession)
		userGroup.GET("/:userId/sessions/:sessionId", u.AuthMiddleware.Verify, u.UserController.GetQuestions)
		userGroup.PATCH("/:userId/sessions/:sessionId", u.AuthMiddleware.Verify, u.UserController.SubmitAnswers)
		userGroup.GET("/:userId/sessions/:sessionId/results")

	}
}
