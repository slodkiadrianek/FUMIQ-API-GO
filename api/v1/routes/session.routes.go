package routes

import (
	"FUMIQ_API/api/v1/controllers"
	"FUMIQ_API/middleware"
	"github.com/gin-gonic/gin"
)

type SessionRoutes struct {
	SessionController *controllers.SessionController
	AuthMiddleware    *middleware.AuthMiddleware
}

func NewSessionRoutes(sessionController *controllers.SessionController, authMiddleware *middleware.AuthMiddleware) *SessionRoutes {
	return &SessionRoutes{
		SessionController: sessionController,
		AuthMiddleware:    authMiddleware,
	}
}

func (q *QuizRoutes) SetupSessionRoutes(router *gin.RouterGroup) {
	sessionGroup := router.Group("/quizzes")
	{
		sessionGroup.POST("/:quizId/sessions", q.AuthMiddleware.Verify)
		sessionGroup.GET("/:quizId/sessions", q.AuthMiddleware.Verify, q.QuizController.GetAllQuizzes)
		sessionGroup.PATCH("/:quizId/sessions/:sessionId", q.AuthMiddleware.Verify, q.QuizController.GetQuiz)
		sessionGroup.GET("/:quizId/sessions/:sessionId", q.AuthMiddleware.Verify, q.QuizController.GetQuiz)
		sessionGroup.GET("/:quizId/sessions/:sessionId/results", q.AuthMiddleware.Verify, q.QuizController.GetQuiz)
	}
}
