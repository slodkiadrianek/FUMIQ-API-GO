package services

import (
	"FUMIQ_API/models"
	"FUMIQ_API/repositories"
	"FUMIQ_API/schemas"
	"FUMIQ_API/utils"
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type QuizService struct {
	Logger         *utils.Logger
	DbClient       *mongo.Database
	QuizRepository *repositories.QuizRepository
}

func NewQuizService(logger *utils.Logger, dbClient *mongo.Database, quizRepository *repositories.QuizRepository) *QuizService {
	return &QuizService{
		Logger:         logger,
		DbClient:       dbClient,
		QuizRepository: quizRepository,
	}
}

func (q *QuizService) NewQuiz(ctx context.Context, quiz *schemas.CreateQuiz) (models.Quiz, error) {
	inserRes, err := q.QuizRepository.InsertQuiz(ctx, quiz)
	if err != nil {
		q.Logger.Error(err.Error())
		return models.Quiz{}, err
	}
	return inserRes, nil
}

func (q *QuizService) GetAllQuizzes(ctx context.Context, userId string) ([]models.Quiz, error) {
	res, err := q.QuizRepository.GetAllQuizzes(ctx, userId)
	if err != nil {
		q.Logger.Error(err.Error())
		return []models.Quiz{}, err
	}
	return res, nil
}
