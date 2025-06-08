package services

import (
	"context"
	"fmt"
	"math/rand"

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

func NewSessionService(logger *utils.Logger, sessionRepository *repositories.SessionRepository, quizRepository *repositories.QuizRepository, dbClient *mongo.Database) *SessionService {
	return &SessionService{
		Logger:            logger,
		SessionRepository: sessionRepository,
		QuizRepository:    quizRepository,
		DbClient:          dbClient,
	}
}

func (s *SessionService) StartNewSession(ctx context.Context, quizId string, userId string) (models.Session, error) {
	if s.QuizRepository == nil {
		fmt.Println("SIGMA")
	}
	ress, err := s.QuizRepository.GetQuizByQuizIdAndUserId(ctx, quizId, userId)
	if err != nil {
		return models.Session{}, err
	}
	fmt.Println(ress)

	res, err := s.SessionRepository.FindSesionByQuizIdAndUserId(ctx, quizId, userId)
	if err != nil {
		if err.Error() == "Quiz error : Quiz  not found for "+userId {
			error := true
			code := rand.Intn(900000) + 100000
			for error {
				code := rand.Intn(900000) + 100000
				res := s.SessionRepository.FindSessionByCode(ctx, code)
				error = res
			}
			res, err := s.SessionRepository.CreateNewSession(ctx, quizId, userId, code)
			if err != nil {
				return models.Session{}, err
			}
			return res, nil
		}
	}
	return res, nil
}

func (s *SessionService) GetInfoAboutSessions(ctx context.Context, quizId string) ([]models.SessionInfo, error) {
	res, err := s.SessionRepository.FindAllUserSessions(ctx, quizId)
	if err != nil {
		return []models.SessionInfo{}, err
	}
	responseData := []models.SessionInfo{}
	for _, v := range res {
		responseData = append(responseData, models.SessionInfo{
			StartedAt:            v.StartedAt,
			EndedAt:              v.EndedAt,
			AmountOfParticipants: len(v.Competitors),
		})
	}

	return responseData, nil
}

func (s *SessionService) EndQuizSession(ctx context.Context, quizId, sessionId string) error {
	err := s.SessionRepository.EndSession(ctx, quizId, sessionId)
	if err != nil {
		return err
	}
	return nil
}
