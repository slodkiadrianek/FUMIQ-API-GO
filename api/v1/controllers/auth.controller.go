package controllers

import (
	"FUMIQ_API/schemas"
	"FUMIQ_API/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	Logger utils.Logger
	//AuthService *services.AuthService
	//ctx         *gin.Context
}

func (a *AuthController) Register(c *gin.Context) {
	var user schemas.RegisterUser
	if err := c.ShouldBindJSON(&user); err != nil {
		a.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(user)
}
