package controllers

import (
	"FUMIQ_API/schemas"
	"FUMIQ_API/services"
	"FUMIQ_API/utils"

	"github.com/gin-gonic/gin"
	"net/http"
)

type QuizController struct {
	Logger      utils.Logger
	QuizService *services.QuizService
}

func NewQuizController(logger utils.Logger, quizService *services.QuizService) QuizController {
	return QuizController{
		Logger:      logger,
		QuizService: quizService,
	}
}

func (q *QuizController) NewQuiz(c *gin.Context) {
	var data *schemas.CreateQuiz
	err := c.ShouldBind(data)
	if err != nil {
		q.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Something went wrong",
			},
		})
		return
	}
	res, err := q.QuizService.NewQuiz(c, data)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"quiz": res}})
}

func (q *QuizController) GetAllQuizzes(c *gin.Context) {
	userId := c.Param("userId")
	res, err := q.QuizService
}
