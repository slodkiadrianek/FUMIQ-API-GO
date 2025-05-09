package services

import (
	"FUMIQ_API/middleware"
	"FUMIQ_API/repositories"
	"FUMIQ_API/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserService struct {
	Logger         *utils.Logger
	UserRepository *repositories.UserRepository
	DbClient       *mongo.Database
	AuthMiddleware *middleware.AuthMiddleware
}

func NewUserService(logger *utils.Logger, userRepository *repositories.UserRepository, dbClient *mongo.Database,
	authMiddleware *middleware.AuthMiddleware) *UserService {
	return &UserService{
		Logger:         logger,
		UserRepository: userRepository,
		DbClient:       dbClient,
		AuthMiddleware: authMiddleware,
	}
}
