package controllers

import (
	"FUMIQ_API/schemas"
	"FUMIQ_API/services"
	"FUMIQ_API/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	Logger      utils.Logger
	AuthService *services.AuthService
}

func (a *AuthController) Register(c *gin.Context) {
	var user schemas.RegisterUser
	err := c.BindJSON(&user)
	if err != nil {
		a.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := a.AuthService.RegisterUser(c, &user)
	if err != nil {
		a.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "data": gin.H{
		"user": data,
	}})
}
