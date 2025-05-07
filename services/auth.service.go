package services

import (
	"FUMIQ_API/middleware"
	"FUMIQ_API/models"
	"FUMIQ_API/repositories"
	"FUMIQ_API/schemas"
	"FUMIQ_API/utils"
	"context"
	"errors"
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
	AuthMiddleware *middleware.AuthMiddleware
}

func NewAuthService(dbClient *mongo.Database, logger *utils.Logger, authRepository *repositories.UserRepository,
	authMiddleware *middleware.AuthMiddleware) *AuthService {
	return &AuthService{
		DbClient:       dbClient,
		Logger:         logger,
		AuthRepository: authRepository,
		AuthMiddleware: authMiddleware,
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
		a.Logger.Error(err.Error())
		return models.User{}, err
	}
	return insertRes, nil
}

func (a AuthService) LoginUser(ctx context.Context, user *schemas.LoginUser) (string, error) {
	res := a.DbClient.Collection("Users").FindOne(ctx, bson.M{"email": user.Email})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		a.Logger.Error("User with this email not found", user.Email)
		err := models.NewError(400, "User", "User with this email not found")
		return "", err
	}
	var userFromDb models.User
	fmt.Println(res)
	err := res.Decode(&userFromDb)
	fmt.Println(err)
	if err != nil {
		a.Logger.Error(err.Error())
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(user.Password))
	if err != nil {
		a.Logger.Error("Password is incorrect")
		err := models.NewError(400, "Password", "Password is incorrect")
		return "", err
	}
	token, err := a.AuthMiddleware.Sign(userFromDb)
	return token, nil
}
