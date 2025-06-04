package services

import (
	"context"

	"FUMIQ_API/models"
	"FUMIQ_API/repositories"
	"FUMIQ_API/schemas"
	"FUMIQ_API/utils"

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
		q.Logger.Error(err.Error(), userId)
		return []models.Quiz{}, err
	}
	return res, nil
}

func (q *QuizService) GetQuiz(ctx context.Context, quizId string) (models.Quiz, error) {
	res, err := q.QuizRepository.GetQuiz(ctx, quizId)
	if err != nil {
		q.Logger.Error(err.Error(), quizId)
		return models.Quiz{}, err
	}
	return res, nil
}

func (q *QuizService) UpdateQuiz(ctx context.Context, quizId string, updateData schemas.CreateQuiz) error {
	err := q.QuizRepository.UpdateQuiz(ctx, quizId, updateData)
	if err != nil {
		q.Logger.Error(err.Error(), quizId)
		return err
	}
	return nil
}

func (q *QuizService) DeleteQuiz(ctx context.Context, quizId string) error {
	err := q.QuizRepository.DeleteQuiz(ctx, quizId)
	if err != nil {
		q.Logger.Error(err.Error(), quizId)
		return err
	}
	return nil
}
