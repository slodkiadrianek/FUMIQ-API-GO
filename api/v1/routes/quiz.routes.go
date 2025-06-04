package routes

import (
	"FUMIQ_API/api/v1/controllers"
	"FUMIQ_API/middleware"
	"FUMIQ_API/schemas"

	// "FUMIQ_API/schemas"

	"github.com/gin-gonic/gin"
)

type QuizRoutes struct {
	QuizController *controllers.QuizController
	AuthMiddleware *middleware.AuthMiddleware
}

func NewQuizRoutes(quizController *controllers.QuizController, authMiddleware *middleware.AuthMiddleware) *QuizRoutes {
	return &QuizRoutes{
		quizController,
		authMiddleware,
	}
}

func (q *QuizRoutes) SetupQuizRoutes(router *gin.RouterGroup) {
	quizGroup := router.Group("/quizzes")
	{
		quizGroup.POST("/", q.AuthMiddleware.Verify, middleware.ValidateRequestData[*schemas.CreateQuiz]("body"), q.QuizController.NewQuiz)
		quizGroup.GET("/users/:userId", q.AuthMiddleware.Verify, q.QuizController.GetAllQuizzes)
		quizGroup.GET("/:quizId", q.AuthMiddleware.Verify, q.QuizController.GetQuiz)
		quizGroup.PUT("/:quizId", q.AuthMiddleware.Verify, q.QuizController.UpdateQuiz)
		quizGroup.DELETE("/:quizId")
	}
}
