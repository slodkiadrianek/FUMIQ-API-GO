package main

import (
	"FUMIQ_API/api/v1/routes"
	"FUMIQ_API/config"
	"FUMIQ_API/middleware"
	"FUMIQ_API/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.<

func main() {
	configStruct := config.Config{}
	envVariables, err := configStruct.LoadConfig()
	if err != nil {
		panic("Error loading .env file")
	}
	cacheService := config.ConnectToCache(envVariables.CacheLink)
	logger := utils.Logger{}
	loggerService := logger.CreateLogger()

	authSevice := middleware.AuthMiddleware{
		Secret:  envVariables.JWTSecret,
		Caching: cacheService,
		Logger:  loggerService,
		Ctx:     &gin.Context{},
	}
	fmt.Println(authSevice)
	router := gin.Default()
	router.Use(middleware.ErrorMiddleware())
	routes.SetupRoutes(router)
	err = router.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}
	err = router.Run(":3007")
	if err != nil {
		panic(err)
	}
}
