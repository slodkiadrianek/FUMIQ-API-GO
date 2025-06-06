package repositories

import (
	"context"

	"FUMIQ_API/config"
	"FUMIQ_API/models"
	"FUMIQ_API/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func (s *SessionRepository) FindSesionByQuizIdAndUserId(ctx context.Context, quizId string, userId string) (models.Session, error) {
	userObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		s.Logger.Error("Failed to convert  id to object id", err)
		return models.Session{}, models.NewError(400, "Database", "Failed to convert id to object id")
	}
	quizObjectId, err := primitive.ObjectIDFromHex(quizId)
	if err != nil {
		s.Logger.Error("Failed to convert  id to object id", err)
		return models.Session{}, models.NewError(400, "Database", "Failed to convert id to object id")
	}
	var data models.Session
	res := s.DbClient.Collection("Sessions").FindOne(ctx, bson.M{"quizId": quizObjectId, "userId": userObjectId, "isActive": true})
	err = res.Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			s.Logger.Error("Quiz not found", userId)
			return models.Session{}, models.NewError(400, "Quiz", "Quiz  not found for "+userId)
		} else {
			s.Logger.Error("Something went wrong during finding a quiz", quizId)
			return models.Session{}, models.NewError(400, "Quiz", "Something went wrong during finding quiz")
		}
	}
	return data, nil
}
