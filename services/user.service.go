package services

import (
	"FUMIQ_API/middleware"
	"FUMIQ_API/models"
	"FUMIQ_API/repositories"
	"FUMIQ_API/schemas"
	"FUMIQ_API/utils"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
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

func (u *UserService) GetUser(ctx context.Context, userId string) (models.User, error) {
	user, err := u.UserRepository.GetUser(ctx, userId)
	if err != nil {
		u.Logger.Error(err.Error())
		return models.User{}, err
	}
	return user, nil
}

func (u *UserService) ChangePassword(ctx context.Context, userId string, passwords schemas.ChangePassword) error {
	user, err := u.UserRepository.GetUser(ctx, userId)
	if err != nil {
		u.Logger.Error(err.Error())
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwords.OldPassword))
	if err != nil {
		u.Logger.Error("Current password is incorrect")
		return err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwords.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		u.Logger.Error(err.Error())
		return err
	}
	user.Password = string(bytes)
	_, err = u.UserRepository.DbClient.Collection("Users").UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": bson.M{"password": user.Password}})
	if err != nil {
		u.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (u *UserService) DeleteUser(ctx context.Context, userId string, password schemas.PasswordBody) error {
	user, err := u.UserRepository.GetUser(ctx, userId)
	if err != nil {
		u.Logger.Error(err.Error())
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password.Password))
	if err != nil {
		u.Logger.Error(err.Error())
		return err
	}
	err = u.UserRepository.DeleteUser(ctx, userId)
	if err != nil {
		u.Logger.Error(err.Error())
		return err
	}
	return nil

}
