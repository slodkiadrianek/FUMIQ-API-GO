package main

import (
	"FUMIQ_API/config"
	"github.com/gin-gonic/gin"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	type Application struct {
		Router     *gin.Engine
		Config     *config.Config
		Logger     *logger.Logger
		Service    *services.Service
		Controller *controllers.Controller
	}

}
