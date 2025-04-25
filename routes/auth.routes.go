package routes

import "github.com/gin-gonic/gin"

func SetupAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register")
		authGroup.POST("/login")
		authGroup.GET("/check")
		authGroup.POST("/logout")
		authGroup.POST("/reset-password")
		authGroup.POST("/reset-password/:token")
		authGroup.GET("/activate/:token")

	}
}
