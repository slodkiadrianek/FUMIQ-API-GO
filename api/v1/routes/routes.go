package routes

import "github.com/gin-gonic/gin"

type SetupRoutes struct {
	AuthRoutes    *AuthRoutes
	UserRoutes    *UserRoutes
	QuizRoutes    *QuizRoutes
	SessionRoutes *SessionRoutes
}

func (s *SetupRoutes) SetupRoutes(router *gin.Engine) {
	routesGroup := router.Group("/api/v1")

	s.AuthRoutes.SetupAuthRoutes(routesGroup)
	s.UserRoutes.SetupUserRoutes(routesGroup)
	s.QuizRoutes.SetupQuizRoutes(routesGroup)
	s.SessionRoutes.SetupSessionRoutes(routesGroup)
}
