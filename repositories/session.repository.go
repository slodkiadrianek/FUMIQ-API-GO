package repositories

import (
	"FUMIQ_API/config"
	"FUMIQ_API/utils"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SessionRepository struct {
	DbClient *mongo.Database
	Logger   *utils.Logger
	Caching  *config.CacheService
}

func NewSessionRepository(dbClient *mongo.Database, logger *utils.Logger, caching *config.CacheService) *SessionRepository {
	return &SessionRepository{
		DbClient: dbClient,
		Logger:   logger,
		Caching:  caching,
	}
}
