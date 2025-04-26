package services

import (
	"FUMIQ_API/middleware"
	"FUMIQ_API/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	BaseService
	Auth *middleware.AuthMiddleware
}

func (a AuthService) RegisterUser(ctx context.Context, user *models.User) error {
	res, err := a.DbClient.Collection("Users").Find(ctx, bson.M{"email": user.Email})
	if err != nil {
		a.Logger.Error(err.Error())
		return err
	}
	if res != nil {
		a.Logger.Error("User already exists")
		return models.NewError(400, "User", "User already exists")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		a.Logger.Error(err.Error())
		return err
	}
	user.Password = string(bytes)
	err = a.InsertToDatabaseAndCache(ctx, "User", user, "Users")
	if err != nil {
		a.Logger.Error(err.Error())
		return err
	}
	token, err := a.Auth.Sign(*user)
	if err != nil {
		a.Logger.Error(err.Error())
		return err
	}
	fmt.Println(token)
	return nil
}
