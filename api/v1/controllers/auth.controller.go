package controllers

import (
	"FUMIQ_API/schemas"
	"FUMIQ_API/services"
	"FUMIQ_API/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Logger      utils.Logger
	AuthService *services.AuthService
}

func NewAuthController(logger utils.Logger, authService *services.AuthService) *AuthController {
	return &AuthController{
		Logger:      logger,
		AuthService: authService,
	}
}

func (a *AuthController) Register(c *gin.Context) {
	userData, _ := c.Get("validatedData")
	user, ok := userData.(*schemas.RegisterUser)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Proper data does not exist",
			},
		})
		return
	}
	_, err := a.AuthService.RegisterUser(c, user)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func (a *AuthController) Login(c *gin.Context) {
	userData, _ := c.Get("validatedData")
	user, ok := userData.(*schemas.LoginUser)
	fmt.Println(ok)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Proper data does not exist",
			},
		})
		return
	}
	token, err := a.AuthService.LoginUser(c, user)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a *AuthController) Verify(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"success": true, "data": gin.H{"user": gin.H{
		"id":        c.GetString("userId"),
		"firstName": c.GetString("firstName"), "lastName": c.GetString("lastName"),
	}}})
}

func (a *AuthController) Logout(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{})
}
