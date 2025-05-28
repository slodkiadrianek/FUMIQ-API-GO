package main

import (
	"FUMIQ_API/api/v1/controllers"
	"FUMIQ_API/api/v1/routes"
	"FUMIQ_API/config"
	"FUMIQ_API/middleware"
	"FUMIQ_API/repositories"
	"FUMIQ_API/services"
	"FUMIQ_API/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	configStruct := config.Config{}
	envVariables, err := configStruct.LoadConfig()
	if err != nil {
		panic("Error loading .env file")
	}
	cacheService := config.ConnectToCache(envVariables.CacheLink)
	logger := utils.NewLogger()
	loggerService := logger.CreateLogger()
	DbClient, err := config.Connect(envVariables.DatabaseLink)
	if err != nil {
		panic("Error connecting to database")
	}
	AuthMiddleware := middleware.NewAuthMiddleware(envVariables.JWTSecret, loggerService, cacheService)
	UserRepository := repositories.NewUserRepository(DbClient, &loggerService, cacheService)
	QuizRepository := repositories.NewQuizRepository(DbClient, &loggerService, cacheService)
	AuthService := services.NewAuthService(DbClient, &loggerService, UserRepository, AuthMiddleware)
	UserService := services.NewUserService(&loggerService, UserRepository, DbClient, AuthMiddleware)
	QuizService := services.NewQuizService(&loggerService, DbClient, QuizRepository)
	authController := controllers.NewAuthController(loggerService, AuthService)
	userController := controllers.NewUserController(loggerService, UserService)
	quizController := controllers.NewQuizController(loggerService, QuizService)

	AuthRoutes := routes.NewAuthRoutes(authController, AuthMiddleware)
	UserRoutes := routes.NewUserRoutes(userController, AuthMiddleware)
	quizRoutes := routes.NewQuizRoutes(&quizController, AuthMiddleware)
	routesConfig := routes.SetupRoutes{AuthRoutes: AuthRoutes, UserRoutes: UserRoutes, QuizRoutes: quizRoutes}
	router.Use(middleware.ErrorMiddleware())
	routesConfig.SetupRoutes(router)
	// gin.SetMode(gin.ReleaseMode)
	err = router.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}
	err = router.Run(":3008")
	if err != nil {
		panic(err)
	}
}
