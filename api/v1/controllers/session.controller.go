package controllers

import (
	// "net/http"

	"FUMIQ_API/services"
	"FUMIQ_API/utils"

	"github.com/gin-gonic/gin"
)

type SessionController struct {
	Logger         *utils.Logger
	SessionService *services.SessionService
}

func NewSessionController(logger *utils.Logger, sessionService *services.SessionService) *SessionController {
	return &SessionController{
		Logger:         logger,
		SessionService: sessionService,
	}
}

func (s *SessionController) StartNewSession(c *gin.Context) {
	// quizIdData, _ := c.Get("validatedParams")
	// quizId, ok := quizIdData.(string)
	// if !ok {
	// 	s.Logger.Error("Proper data does not exist")
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": gin.H{
	// 			"category":    "Validation",
	// 			"description": "Something went wrong",
	// 		},
	// 	})
	// 	return
	// }
	quizId := c.Param("quizId")
	userId := c.GetString("userId")
	res, err := s.SessionService.StartNewSession(c, quizId, userId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, gin.H{"success": true, "data": gin.H{
		"quiz": res,
	}})
}
