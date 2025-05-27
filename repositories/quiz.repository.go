package repositories

import (
	"FUMIQ_API/config"
	"FUMIQ_API/models"
	"FUMIQ_API/schemas"
	"FUMIQ_API/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type QuizRepository struct {
	DbClient *mongo.Database
	Logger   *utils.Logger
	Caching  *config.CacheService
}

func NewQuizRepository(dbClient *mongo.Database, logger *utils.Logger, caching *config.CacheService) *QuizRepository {
	return &QuizRepository{
		DbClient: dbClient,
		Logger:   logger,
		Caching:  caching,
	}
}

func (q *QuizRepository) InsertQuiz(ctx context.Context, data *schemas.CreateQuiz) (models.Quiz, error) {
	quiz := models.NewQuiz(data.UserId, data.Title, data.Description, data.TimeLimit, data.Questions)
	_, err := q.DbClient.Collection("Quizzes").InsertOne(ctx, quiz)

	if err != nil {
		q.Logger.Error("Something went wrong during inserting to database")
		return models.Quiz{}, models.NewError(400, "Database", "Something went wrong during inserting to database")
	}
	dataBytes, err := json.Marshal(quiz)
	if err != nil {
		q.Logger.Error("Failed to marshal data for caching")
		return models.Quiz{}, models.NewError(500, "Cache", "Failed to marshal data for caching")
	}
	err = q.Caching.SetData(ctx, fmt.Sprintf("Quiz-%s", quiz.ID), string(dataBytes), 1000)
	if err != nil {
		q.Logger.Error("Something went wrong during adding data to cache")
		q.Logger.Error("Cache operation failed but database insert was successful")
	}
	return quiz, nil
}

func (q *QuizRepository) GetQuiz(ctx context.Context, quizId string) (models.Quiz, error) {
	cacheKey := "Quiz-" + quizId
	exist, err := q.Caching.ExistData(ctx, cacheKey)
	if err != nil {
		q.Logger.Error("Something went wrong during checking cache existence", quizId)
		return models.Quiz{}, models.NewError(400, "Cache", "Something went wrong during checking cache existence")
	}
	if exist > 0 {
		data, err := q.Caching.GetData(ctx, cacheKey)
		if err != nil {
			q.Logger.Error("Something went wrong during getting user from cache", quizId)
			return models.Quiz{}, models.NewError(400, "Cache", "Something went wrong during getting user from cache")
		}
		var quiz models.Quiz
		err = json.Unmarshal([]byte(data), &quiz)
		if err != nil {
			q.Logger.Error("Failed to unmarshal user", err)
			return models.Quiz{}, models.NewError(400, "Cache", "Failed to unmarshal user")
		}
		return quiz, nil
	}
	objectID, err := primitive.ObjectIDFromHex(quizId)

	if err != nil {
		q.Logger.Error("Failed to convert user id to object id", err)
		return models.Quiz{}, models.NewError(400, "Database", "Failed to convert user id to object id")
	}
	res := q.DbClient.Collection("Users").FindOne(ctx, bson.D{{"_id", objectID}})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		q.Logger.Error("Something went wrong during taking data from database")
		return models.Quiz{}, models.NewError(400, "Database", "Something went wrong during taking data from database")
	}
	var quiz models.Quiz
	err = res.Decode(&quiz)

	if err != nil {
		q.Logger.Error("Failed to decode user", err)
		return models.Quiz{}, models.NewError(400, "Database", "Failed to decode user")
	}
	dataBytes, err := json.Marshal(quiz)
	if err != nil {
		q.Logger.Error("Failed to marshal data for caching")
		return models.Quiz{}, models.NewError(500, "Cache", "Failed to marshal data for caching")
	}
	err = q.Caching.SetData(ctx, cacheKey, string(dataBytes), 1000)
	if err != nil {
		q.Logger.Error("Something went wrong during adding data to cache")
		q.Logger.Error("Cache operation failed but database insert was successful")
	}
	return quiz, nil
}

func (q *QuizRepository) GetAllQuizzes(userId string) ([]models.Quiz, error) {
	ObjectID, err := primitive.ObjectIDFromHex(userId)
}
