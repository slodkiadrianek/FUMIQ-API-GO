package routes

import "github.com/gin-gonic/gin"

func SetupQuizRoutes(router *gin.RouterGroup) {
	quizGroup := router.Group("/quizzes")
	{
		quizGroup.POST("/")
		quizGroup.GET("/:quizId")
		quizGroup.PATCH("/:quizId")
		quizGroup.PUT("/:quizId")
		quizGroup.DELETE("/:quizId")
	}
}
