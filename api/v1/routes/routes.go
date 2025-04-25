package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	routesGroup := router.Group("/api/v1")

	SetupAuthRoutes(routesGroup)
	SetupUserRoutes(routesGroup)
	SetupQuizRoutes(routesGroup)

}
