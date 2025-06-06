package services

import (
	"context"
	"fmt"

	"FUMIQ_API/models"
	"FUMIQ_API/repositories"
	"FUMIQ_API/utils"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SessionService struct {
	Logger            *utils.Logger
	SessionRepository *repositories.SessionRepository
	QuizRepository    *repositories.QuizRepository
	DbClient          *mongo.Database
}

func NewSessionService(logger *utils.Logger, sessionRepository *repositories.SessionRepository, QuizRepository *repositories.QuizRepository, dbClient *mongo.Database) *SessionService {
	return &SessionService{
		Logger:            logger,
		SessionRepository: sessionRepository,
		DbClient:          dbClient,
	}
}

func (s *SessionService) StartNewSession(ctx context.Context, quizId string, userId string) (models.Session, error) {
	ress, err := s.QuizRepository.GetQuizByQuizIdAndUserId(ctx, quizId, userId)
	if err != nil {
		return models.Session{}, err
	}
	fmt.Println(ress)

	res, err := s.SessionRepository.FindSesionByQuizIdAndUserId(ctx, quizId, userId)
	if err != nil {
		fmt.Println("HUJUS")
		fmt.Println(err.Error())
	}
	fmt.Println("HUJ")
	return res, nil
}
