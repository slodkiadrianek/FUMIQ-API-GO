package routes

import "github.com/gin-gonic/gin"

type SetupRoutes struct {
	AuthRoutes *AuthRoutes
}

func (s *SetupRoutes) SetupRoutes(router *gin.Engine) {
	routesGroup := router.Group("/api/v1")

	s.AuthRoutes.SetupAuthRoutes(routesGroup)
	SetupUserRoutes(routesGroup)
	SetupQuizRoutes(routesGroup)

}
