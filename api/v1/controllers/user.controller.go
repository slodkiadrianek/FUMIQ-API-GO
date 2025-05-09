package controllers

import (
	"FUMIQ_API/services"
	"FUMIQ_API/utils"
	"github.com/gin-gonic/gin"
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

func (u *UserController) GetUser(c *gin.Context) {}
