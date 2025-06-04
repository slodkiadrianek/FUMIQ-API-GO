package services

import (
	"FUMIQ_API/repositories"
	"FUMIQ_API/utils"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SessionService struct {
	Logger            *utils.Logger
	SessionRepository *repositories.SessionRepository
	DbClient          *mongo.Database
}

func NewSessionService(logger *utils.Logger, sessionRepository *repositories.SessionRepository, dbClient *mongo.Database) *SessionService {
	return &SessionService{
		Logger:            logger,
		SessionRepository: sessionRepository,
		DbClient:          dbClient,
	}
}
