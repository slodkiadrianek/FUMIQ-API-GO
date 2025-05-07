package routes

import (
	"FUMIQ_API/api/v1/controllers"
	"FUMIQ_API/middleware"
	"FUMIQ_API/schemas"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	AuthController *controllers.AuthController
}

func (a *AuthRoutes) SetupAuthRoutes(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", middleware.ValidateRequestData[schemas.RegisterUser], a.AuthController.Register)
		authGroup.POST("/login", middleware.ValidateRequestData[schemas.LoginUser])
		authGroup.GET("/check")
		authGroup.POST("/logout")
		authGroup.POST("/reset-password")
		authGroup.POST("/reset-password/:token")
		authGroup.GET("/activate/:token")

	}
}
