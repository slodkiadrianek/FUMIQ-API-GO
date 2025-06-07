package repositories

import (
	"context"
	"encoding/json"

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

func (s *SessionRepository) CreateNewSession(ctx context.Context, quizId, userId string, code int) (models.Session, error) {
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
	session := models.NewSession(userObjectId, quizObjectId, code)
	_, err = s.DbClient.Collection("Sessions").InsertOne(ctx, session)
	if err != nil {
		s.Logger.Error("Something went wrong during inserting new session to database", session.ID)
		return models.Session{}, models.NewError(400, "Database", "Something went wrong during inserting new session to database")
	}
	return session, nil
}

func (s *SessionRepository) FindSessionByCode(ctx context.Context, code int) bool {
	var data models.Session
	res := s.DbClient.Collection("Sessions").FindOne(ctx, bson.M{"code": code, "isActive": true})
	err := res.Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			s.Logger.Info("Unique code has been found", code)
			return false
		} else {
			s.Logger.Error("This quiz code is in use", code)
			return true
		}
	}
	return false
}

func (s *SessionRepository) FindAllUserSessions(ctx context.Context, quizId string) ([]models.Session, error) {
	cacheKey := "Sessions-" + quizId
	exist, err := s.Caching.ExistData(ctx, cacheKey)
	if err != nil {
		s.Logger.Error("Something went wrong during checking cache existence", quizId)
		return []models.Session{}, models.NewError(400, "Cache", "Something went wrong during checking cache existence")
	}

	if exist > 0 {
		var data []models.Session
		res, err := s.Caching.GetData(ctx, cacheKey)
		if err != nil {
			s.Logger.Error("Something went wrong during getting sessions from cache", quizId)
			return []models.Session{}, models.NewError(400, "Cache", "Something went wrong during getting sessions from cache")
		}

		err = json.Unmarshal([]byte(res), &data)
		if err != nil {
			s.Logger.Error("Failed to unmarshal sessions", err)
			return []models.Session{}, models.NewError(400, "Cache", "Failed to unmarshal sessions")
		}

		return data, nil
	}
	quizObjectId, err := primitive.ObjectIDFromHex(quizId)
	if err != nil {
		s.Logger.Error("Failed to convert quiz id to object id", err)
		return []models.Session{}, models.NewError(400, "Database", "Failed to convert quiz id to object id")
	}
	res, err := s.DbClient.Collection("Sessions").Find(ctx, bson.M{"quizId": quizObjectId})
	if err != nil {
		s.Logger.Error("Something went wrong during taking data from database", quizId)
		return []models.Session{}, models.NewError(400, "Database", "Something went wrong during taking data from database")
	}
	var data []models.Session
	err = res.All(ctx, &data)
	if err != nil {
		s.Logger.Error("Something went wrong during decoding data", res)
		return []models.Session{}, models.NewError(400, "Database", "Something went wrong during decoding data")
	}
	bodyBytes, err := json.Marshal(data)
	if err != nil {
		s.Logger.Error("Failed to marshal data for caching")
		return []models.Session{}, models.NewError(500, "Cache", "Failed to marshal data for caching")
	}
	err = s.Caching.SetData(ctx, cacheKey, string(bodyBytes), 1000)
	if err != nil {
		s.Logger.Error("Something went wrong during adding data to cache")
		s.Logger.Error("Cache operation failed but database insert was successful")
	}
	return data, nil
}

func (s *SessionRepository) EndSession(ctx context.Context, quizId, sessionId string) error {
	quizObjectId, err := primitive.ObjectIDFromHex(quizId)
	if err != nil {
		s.Logger.Error("Failed to convert quiz id to object id", err)
		return []models.Session{}, models.NewError(400, "Database", "Failed to convert quiz id to object id")
	}
	sessionIdObjectId, err := primitive.ObjectIDFromHex(sessionId)
	if err != nil {
		s.Logger.Error("Failed to convert quiz id to object id", err)
		return []models.Session{}, models.NewError(400, "Database", "Failed to convert quiz id to object id")
	}
}
