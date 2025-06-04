package controllers

import (
	"FUMIQ_API/services"
	"FUMIQ_API/utils"
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
