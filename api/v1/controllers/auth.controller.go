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
	_, err = a.AuthService.RegisterUser(c, &user)
	if err != nil {
		a.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (a *AuthController) Login(c *gin.Context) {
	var user schemas.LoginUser
	err := c.BindJSON(&user)
	if err != nil {
		a.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := a.AuthService.LoginUser(c, &user)
	if err != nil {
		a.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
