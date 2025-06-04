package controllers

import (
	"net/http"

	"FUMIQ_API/schemas"
	"FUMIQ_API/services"
	"FUMIQ_API/utils"

	"github.com/gin-gonic/gin"
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
	quizData, _ := c.Get("validatedData")
	quiz, ok := quizData.(*schemas.CreateQuiz)
	if !ok {
		q.Logger.Error("Proper data does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Something went wrong",
			},
		})
		return
	}
	res, err := q.QuizService.NewQuiz(c, quiz)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "data": gin.H{"quiz": res}})
}

func (q *QuizController) GetAllQuizzes(c *gin.Context) {
	userIdData, _ := c.Get("validatedParams")
	userId, ok := userIdData.(string)
	if !ok {
		q.Logger.Error("Proper data does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Something went wrong",
			},
		})
		return
	}
	res, err := q.QuizService.GetAllQuizzes(c, userId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, gin.H{"success": true, "data": gin.H{"quizzes": res}})
}

func (q *QuizController) GetQuiz(c *gin.Context) {
	quizIdData, _ := c.Get("validatedParams")
	quizId, ok := quizIdData.(string)

	if !ok {
		q.Logger.Error("Proper data does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Something went wrong",
			},
		})
		return
	}
	res, err := q.QuizService.GetQuiz(c, quizId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, gin.H{"success": true, "data": gin.H{"quiz": res}})
}

func (q *QuizController) UpdateQuiz(c *gin.Context) {
	quizIdData, _ := c.Get("validatedParams")
	quizId, ok := quizIdData.(string)
	if !ok {
		q.Logger.Error("Proper data does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation Params",
				"description": "Something went wrong",
			},
		})
		return
	}
	updateQuizData, _ := c.Get("validatedData")
	updateQuiz, ok := updateQuizData.(schemas.CreateQuiz)
	if !ok {
		q.Logger.Error("Proper data does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation Body",
				"description": "Something went wrong",
			},
		})
		return
	}
	err := q.QuizService.UpdateQuiz(c, quizId, updateQuiz)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

func (q *QuizController) DeleteQuiz(c *gin.Context) {
	quizIdData, _ := c.Get("validatedParams")
	quizId, ok := quizIdData.(string)
	if !ok {
		q.Logger.Error("Proper data does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation Params",
				"description": "Something went wrong",
			},
		})
		return
	}
	err := q.QuizService.DeleteQuiz(c, quizId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
