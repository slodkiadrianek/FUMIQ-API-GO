package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"FUMIQ_API/config"
	"FUMIQ_API/models"
	"FUMIQ_API/schemas"
	"FUMIQ_API/utils"

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
	quiz, err := models.NewQuiz(data.UserId, data.Title, data.Description, data.TimeLimit, data.Questions)
	if err != nil {
		q.Logger.Error("Something went wrong during quiz creation")
		return models.Quiz{}, models.NewError(400, "Quiz Creation", "Something went wrong during quiz creation")
	}
	_, err = q.DbClient.Collection("Quizzes").InsertOne(ctx, quiz)
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
		q.Logger.Error("Failed to convert quiz id to object id", err)
		return models.Quiz{}, models.NewError(400, "Database", "Failed to convert quiz id to object id")
	}
	res := q.DbClient.Collection("Quizzes").FindOne(ctx, bson.M{"_id": objectID})
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

func (q *QuizRepository) GetAllQuizzes(ctx context.Context, userId string) ([]models.Quiz, error) {
	cacheKey := "Quizzes-" + userId
	exist, err := q.Caching.ExistData(ctx, cacheKey)
	if err != nil {
		q.Logger.Error("Something went wrong during checking cache existence", userId)
		return []models.Quiz{}, models.NewError(400, "Cache", "Something went wrong during checking cache existence")
	}
	if exist > 0 {
		data, err := q.Caching.GetData(ctx, cacheKey)
		if err != nil {
			q.Logger.Error("Something went wrong during getting user from cache", userId)
			return []models.Quiz{}, models.NewError(400, "Cache", "Something went wrong during getting user from cache")
		}
		var quizzes []models.Quiz
		err = json.Unmarshal([]byte(data), &quizzes)
		if err != nil {
			q.Logger.Error("Failed to unmarshal user", err)
			return []models.Quiz{}, models.NewError(400, "Cache", "Failed to unmarshal user")
		}
		return quizzes, nil
	}
	ObjectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		q.Logger.Error("Failed to convert user id to object id", err)
		return []models.Quiz{}, models.NewError(400, "Database", "Failed to convert user id to object id")
	}
	res, err := q.DbClient.Collection("Quizzes").Find(ctx, bson.M{"userId": ObjectID})
	if err != nil {
		q.Logger.Error("Something went wrong during taking data from database")
		return []models.Quiz{}, models.NewError(400, "Database", "Something went wrong during taking data from database")
	}
	var quizzes []models.Quiz
	err = res.All(ctx, &quizzes)
	if err != nil {
		q.Logger.Error("Failed to decode quizzes", err)
		return []models.Quiz{}, models.NewError(400, "Database", "Failed to decode quizzes")
	}
	dataBytes, err := json.Marshal(quizzes)
	if err != nil {
		q.Logger.Error("Failed to marshal data for caching")
		return []models.Quiz{}, models.NewError(500, "Cache", "Failed to marshal data for caching")
	}
	err = q.Caching.SetData(ctx, cacheKey, string(dataBytes), 1000)
	if err != nil {
		q.Logger.Error("Something went wrong during adding data to cache")
		q.Logger.Error("Cache operation failed but database insert was successful")
	}
	return quizzes, nil
}

func (q *QuizRepository) UpdateQuiz(ctx context.Context, quizId string, updateData schemas.CreateQuiz) error {
	quiz, err := models.NewQuiz(updateData.UserId, updateData.Title, updateData.Description, updateData.TimeLimit, updateData.Questions)
	if err != nil {
		q.Logger.Error("Something went wrong during quiz creation")
		return models.NewError(400, "Quiz Creation", "Something went wrong during quiz creation")
	}
	ObjectID, err := primitive.ObjectIDFromHex(quizId)
	if err != nil {
		q.Logger.Error("Failed to convert user id to object id", err)
		return models.NewError(400, "Database", "Failed to convert user id to object id")
	}
	_, err = q.DbClient.Collection("Quizzes").UpdateByID(ctx, ObjectID, quiz)
	if err != nil {
		q.Logger.Error("Something went wrong during inserting  to database an updated data")
		return models.NewError(400, "Database", "Something went wrong during inserting to database an updated data")
	}
	cacheKey := "Quiz-" + quizId
	dataBytes, err := json.Marshal(quiz)
	if err != nil {
		q.Logger.Error("Failed to marshal data for caching")
		return models.NewError(500, "Cache", "Failed to marshal data for caching")
	}
	err = q.Caching.SetData(ctx, cacheKey, string(dataBytes), 1000)
	if err != nil {
		q.Logger.Error("Something went wrong during adding data to cache")
		q.Logger.Error("Cache operation failed but database insert was successful")
	}
	return nil
}

func (q *QuizRepository) DeleteQuiz(ctx context.Context, quizId string) error {
	ObjectID, err := primitive.ObjectIDFromHex(quizId)
	if err != nil {
		q.Logger.Error("Failed to convert user id to object id", err)
		return models.NewError(400, "Database", "Failed to convert user id to object id")
	}
	_, err = q.DbClient.Collection("Quizzes").DeleteOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		q.Logger.Error("Something went wrong during inserting  to database an updated data")
		return models.NewError(400, "Database", "Something went wrong during inserting to database an updated data")
	}
	cacheKey := "Quiz-" + quizId
	err = q.Caching.DeleteData(ctx, cacheKey)
	if err != nil {
		q.Logger.Error("Something went wrong during deleting data from cache")
		q.Logger.Error("Cache operation failed but database delete was successful")
	}
	return nil
}

func (q *QuizRepository) GetQuizByQuizIdAndUserId(ctx context.Context, quizId string, userId string) error {
	userObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		q.Logger.Error("Failed to convert user id to object id", err)
		return models.NewError(400, "Database", "Failed to convert user id to object id")
	}
	quizObjectId, err := primitive.ObjectIDFromHex(quizId)
	if err != nil {
		q.Logger.Error("Failed to convert quiz id to object id", err)
		return models.NewError(400, "Database", "Failed to convert quiz id to object id")
	}
	var data *models.Quiz
	res := q.DbClient.Collection("Quizzes").FindOne(ctx, bson.M{"_id": quizObjectId, "userId": userObjectId})
	err = res.Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			q.Logger.Error("Quiz not found", userId)
			return models.NewError(400, "Quiz", "Quiz  not found for "+userId)
		} else {
			q.Logger.Error("Something went wrong during finding a quiz", quizId)
			return models.NewError(400, "Quiz", "Something went wrong during finding quiz")
		}
	}
	return nil
}
