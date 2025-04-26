package main

import (
	"FUMIQ_API/api/v1/controllers"
	"FUMIQ_API/api/v1/routes"
	"FUMIQ_API/config"
	"FUMIQ_API/middleware"
	"FUMIQ_API/services"
	"FUMIQ_API/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.<

func main() {
	router := gin.Default()
	configStruct := config.Config{}
	envVariables, err := configStruct.LoadConfig()
	if err != nil {
		panic("Error loading .env file")
	}
	cacheService := config.ConnectToCache(envVariables.CacheLink)
	logger := utils.Logger{}
	loggerService := logger.CreateLogger()
	DbClient, err := config.Connect(envVariables.DatabaseLink)
	if err != nil {
		panic("Error connecting to database")
	}
	BaseService := services.BaseService{
		DbClient: DbClient,
		Logger:   &loggerService,
		Caching:  cacheService,
	}
	authController := controllers.AuthController{Logger: loggerService}
	AuthRoutes := routes.AuthRoutes{AuthController: &authController}

	routesConfig := routes.SetupRoutes{AuthRoutes: &AuthRoutes}
	//authMiddleware := middleware.AuthMiddleware{
	//	Secret:  envVariables.JWTSecret,
	//	Caching: cacheService,
	//	Logger:  loggerService,
	//	Ctx:     &gin.Context{},
	//}

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
