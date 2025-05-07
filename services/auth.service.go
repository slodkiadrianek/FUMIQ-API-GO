package services

import (
	"FUMIQ_API/models"
	"FUMIQ_API/repositories"
	"FUMIQ_API/schemas"
	"FUMIQ_API/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Logger         *utils.Logger
	AuthRepository *repositories.UserRepository
	DbClient       *mongo.Database
}

func NewAuthService(dbClient *mongo.Database, logger *utils.Logger, authRepository *repositories.UserRepository) *AuthService {
	return &AuthService{
		DbClient:       dbClient,
		Logger:         logger,
		AuthRepository: authRepository,
	}
}

func (a AuthService) RegisterUser(ctx context.Context, user *schemas.RegisterUser) (models.User, error) {
	res, err := a.DbClient.Collection("Users").Find(ctx, bson.M{"email": user.Email})
	if err != nil {
		a.Logger.Error(err.Error())
		return models.User{}, err
	}
	defer res.Close(ctx)
	if res.Next(ctx) {
		a.Logger.Error("User already exists")
		err := models.NewError(400, "User", "User already exists")
		fmt.Println(err.Error())
		return models.User{}, err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		a.Logger.Error(err.Error())
		return models.User{}, err
	}
	user.Password = string(bytes)

	insertRes, err := a.AuthRepository.InsertUser(ctx, user)
	if err != nil {

	}
	return insertRes, nil
}
