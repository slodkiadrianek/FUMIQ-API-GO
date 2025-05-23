package controllers

import (
	"FUMIQ_API/schemas"
	"FUMIQ_API/services"
	"FUMIQ_API/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	Logger      utils.Logger
	UserService *services.UserService
}

func NewUserController(logger utils.Logger, authService *services.UserService) *UserController {
	return &UserController{
		Logger:      logger,
		UserService: authService,
	}
}

func (u *UserController) GetUser(c *gin.Context) {
	userId := c.Param("userId")
	user, err := u.UserService.GetUser(c, userId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"user": user}})
}

func (u *UserController) ChangePassword(c *gin.Context) {
	userId := c.Param("userId")
	var passwords schemas.ChangePassword
	err := c.BindJSON(passwords)
	if err != nil {
		u.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Something went wrong",
			},
		})
		return
	}
	err = u.UserService.ChangePassword(c, userId, passwords)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

func (u *UserController) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	var password schemas.DeleteUser
	err := c.BindJSON(password)
	if err != nil {
		u.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Something went wrong",
			},
		})
	}
	err = u.UserService.DeleteUser(c, userId, password)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
