package routes

import (
	"FUMIQ_API/api/v1/controllers"
	"FUMIQ_API/middleware"
	"FUMIQ_API/schemas"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	AuthController *controllers.AuthController
	AuthMiddleware *middleware.AuthMiddleware
}

func NewAuthRoutes(authController *controllers.AuthController, authMiddleware *middleware.AuthMiddleware) *AuthRoutes {
	return &AuthRoutes{
		AuthController: authController,
		AuthMiddleware: authMiddleware,
	}
}

func (a *AuthRoutes) SetupAuthRoutes(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	{
<<<<<<< HEAD
		authGroup.POST("/register", middleware.ValidateRequestData[*schemas.RegisterUser]("body"), a.AuthController.Register)
		authGroup.POST("/login", middleware.ValidateRequestData[*schemas.LoginUser]("body"), a.AuthController.Login)
=======
		authGroup.POST("/register", middleware.ValidateRequestData[schemas.RegisterUser](schemas.RegisterSchema), a.AuthController.Register)
		authGroup.POST("/login", middleware.ValidateRequestData[schemas.LoginUser](schemas.LoginSchema), a.AuthController.Login)
>>>>>>> e50232b (VALIDATION)
		authGroup.GET("/check", a.AuthMiddleware.Verify, a.AuthController.Verify)
		authGroup.POST("/logout", a.AuthMiddleware.BlackList, a.AuthController.Logout)
		authGroup.POST("/reset-password")
		authGroup.POST("/reset-password/:token")
		authGroup.GET("/activate/:token")

	}
}
