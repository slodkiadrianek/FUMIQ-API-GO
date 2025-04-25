package routes

import "github.com/gin-gonic/gin"

func SetupUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/:userId")
		userGroup.PATCH("/:userId")
		userGroup.PUT("/:userId")
		userGroup.DELETE("/:userId")
		userGroup.GET("/:userId/quizzes")
		userGroup.POST("/:userId/sessions")
		userGroup.GET("/:userId/sessions/:sessionId")
		userGroup.PATCH("/:userId/sessions/:sessionId")
		userGroup.GET("/:userId/sessions/:sessionId/results")

	}
}
