package main

import (
	"FUMIQ_API/api/v1/controllers"
	"FUMIQ_API/api/v1/routes"
	"FUMIQ_API/config"
	"FUMIQ_API/middleware"
	"FUMIQ_API/repositories"
	"FUMIQ_API/services"
	"FUMIQ_API/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
	AuthMiddleware := middleware.AuthMiddleware{
		Secret:  envVariables.JWTSecret,
		Logger:  loggerService,
		Caching: cacheService,
	}
	if err != nil {
		panic("Error connecting to database")
	}
	UserRepository := repositories.NewUserRepository(DbClient, &loggerService, cacheService)
	BaseService := services.BaseService{
		DbClient: DbClient,
		Logger:   &loggerService,
		Caching:  cacheService,
	}
	AuthService := services.NewAuthService(DbClient, &loggerService, UserRepository, &AuthMiddleware)
	authController := controllers.AuthController{Logger: loggerService, AuthService: AuthService}
	AuthRoutes := routes.AuthRoutes{AuthController: &authController}
	routesConfig := routes.SetupRoutes{AuthRoutes: &AuthRoutes}
	fmt.Println(BaseService)
	router.Use(middleware.ErrorMiddleware())
	routesConfig.SetupRoutes(router)
	//gin.SetMode(gin.ReleaseMode)
	err = router.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}
	err = router.Run(":3008")
	if err != nil {
		panic(err)
	}
}
