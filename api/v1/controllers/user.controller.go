package controllers

import (
	"net/http"

	"FUMIQ_API/schemas"
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

func (u *UserController) GetUser(c *gin.Context) {
	userIdData, _ := c.Get("validatedParams")
	userId, ok := userIdData.(string)
	if !ok {
		u.Logger.Error("Proper data does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Something went wrong",
			},
		})
		return
	}
	user, err := u.UserService.GetUser(c, userId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"user": user}})
}

func (u *UserController) ChangePassword(c *gin.Context) {
	userIdData, _ := c.Get("validatedParams")
	userId, ok := userIdData.(string)
	if !ok {
		u.Logger.Error("Proper data does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Something went wrong",
			},
		})
		return
	}
	passwordsData, _ := c.Get("validatedData")
	passwords, ok := passwordsData.(schemas.ChangePassword)
	if !ok {
		u.Logger.Error("Proper data does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Something went wrong",
			},
		})
		return
	}
	err := u.UserService.ChangePassword(c, userId, passwords)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

func (u *UserController) DeleteUser(c *gin.Context) {
	userIdData, _ := c.Get("validatedParams")
	userId, ok := userIdData.(string)
	if !ok {
		u.Logger.Error("Proper data does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Something went wrong",
			},
		})
		return
	}
	passwordData, _ := c.Get("validatedData")
	password, ok := passwordData.(schemas.DeleteUser)
	if !ok {
		u.Logger.Error("Proper data does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Something went wrong",
			},
		})
		return
	}
	err := u.UserService.DeleteUser(c, userId, password)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

func (u *UserController) UpdateUser(c *gin.Context) {
	userIdData, _ := c.Get("validatedParams")
	userId, ok := userIdData.(string)
	if !ok {
		u.Logger.Error("Proper data does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Something went wrong",
			},
		})
		return
	}
	userData, _ := c.Get("validatedData")
	user, ok := userData.(schemas.UpdateUser)
	if !ok {
		u.Logger.Error("Proper data does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Something went wrong",
			},
		})
		return
	}
	err := u.UserService.UpdateUser(c, userId, user)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

func (u *UserController) JoinSession(c *gin.Context) {
	userId := c.Param("userId")
	// codeData, _ := c.Get("validatedData")
	// code, ok := codeData.(*schemas.JoinQuiz)
	// if !ok {
	// 	u.Logger.Error("Proper data does not exist")
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": gin.H{
	// 			"category":    "Validation",
	// 			"description": "Something went wrong",
	// 		},
	// 	})
	// 	return
	// }
	res, err := u.UserService.JoinSession(c, userId, "318618")
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, gin.H{"success": true, "data": gin.H{"quiz": gin.H{"_id": res}}})
}

func (u *UserController) GetQuestions(c *gin.Context) {
	userId := c.Param("userId")
	sessionId := c.Param("sessionId")
	res, err := u.UserService.GetQuestions(c, userId, sessionId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, gin.H{"success": true, "data": gin.H{
		"quiz": res,
	}})
}

func (u *UserController) SubmitAnswers(c *gin.Context) {
	userId := c.Param("userId")
	sessionId := c.Param("sessionId")
	err := u.UserService.SubmitAnswers(c, userId, sessionId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(204, gin.H{})
}
