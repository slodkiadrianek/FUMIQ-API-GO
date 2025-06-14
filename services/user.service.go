package services

import (
	"context"

	"FUMIQ_API/middleware"
	"FUMIQ_API/models"
	"FUMIQ_API/repositories"
	"FUMIQ_API/schemas"
	"FUMIQ_API/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Logger            *utils.Logger
	UserRepository    *repositories.UserRepository
	SessionRepository *repositories.SessionRepository
	DbClient          *mongo.Database
	AuthMiddleware    *middleware.AuthMiddleware
}

func NewUserService(logger *utils.Logger, userRepository *repositories.UserRepository, dbClient *mongo.Database, sessionRepository *repositories.SessionRepository,
	authMiddleware *middleware.AuthMiddleware,
) *UserService {
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

func (u *UserService) DeleteUser(ctx context.Context, userId string, password schemas.DeleteUser) error {
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

func (u *UserService) UpdateUser(ctx context.Context, userId string, updateUser schemas.UpdateUser) error {
	err := u.UserRepository.UpdateUser(ctx, userId, updateUser)
	if err != nil {
		u.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (u *UserService) JoinSession(ctx context.Context, userId, code string) (string, error) {
	res, err := u.SessionRepository.JoinSession(ctx, userId, code)
	if err != nil {
		return "", nil
	}
	return res, nil
}

func (u *UserService) SubmitAnswers(ctx context.Context, userId, sessionId string) error {
	err := u.SessionRepository.SubmitAnswers(ctx, userId, sessionId)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) GetQuestions(ctx context.Context, userId, sessionId string) (models.SessionQuestions, error) {
	queryRes, err := u.SessionRepository.GetQuestions(ctx, userId, sessionId)
	if err != nil {
		return models.SessionQuestions{}, err
	}
	var user models.Competitor
	doesUserExist := 0
	for _, v := range queryRes.Competitors {
		if v.UserID.Hex() == userId {
			user = v
			doesUserExist = 1
		}
	}
	if doesUserExist != 0 {

		if user.Finished {
			return models.SessionQuestions{}, models.NewError(400, "Session", "You have already finished this session")
		}
		serviceRes := models.SessionQuestions{
			ID:         queryRes.ID,
			Quiz:       queryRes.QuizID,
			Competitor: user,
		}
		return serviceRes, nil
	}
	return models.SessionQuestions{}, models.NewError(400, "Session", "You have to join session")
}
