package controllers

import (
	"FUMIQ_API/services"
	"FUMIQ_API/utils"
)

type QuizController struct {
	Logger      utils.Logger
	QuizService *services.QuizService
}
