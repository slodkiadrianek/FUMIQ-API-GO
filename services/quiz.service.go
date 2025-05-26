package services

import (
	"FUMIQ_API/repositories"
	"FUMIQ_API/utils"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type QuizService struct {
	Logger         *utils.Logger
	DbClient       *mongo.Database
	QuizRepository *repositories.QuizRepository
}
