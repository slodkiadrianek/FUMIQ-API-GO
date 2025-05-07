package repositories

import (
	"FUMIQ_API/config"
	"FUMIQ_API/models"
	"FUMIQ_API/schemas"
	"FUMIQ_API/utils"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	DbClient *mongo.Database
	Logger   *utils.Logger
	Caching  *config.CacheService
}

func NewUserRepository(dbClient *mongo.Database, logger *utils.Logger, caching *config.CacheService) *UserRepository {
	return &UserRepository{
		DbClient: dbClient,
		Logger:   logger,
		Caching:  caching,
	}
}

func (u *UserRepository) InsertUser(ctx context.Context, data *schemas.RegisterUser) (models.User, error) {
	user := models.NewUser(data.FirstName, data.LastName, data.Email, data.Password)
	_, err := u.DbClient.Collection("Users").InsertOne(ctx, user)
	if err != nil {
		u.Logger.Error("Something went wrong during inserting to database")
		return models.User{}, models.NewError(400, "Database", "Something went wrong during inserting to database")
	}
	dataBytes, err := json.Marshal(user)
	if err != nil {
		u.Logger.Error("Failed to marshal data for caching")
		return models.User{}, models.NewError(500, "Cache", "Failed to marshal data for caching")
	}
	err = u.Caching.SetData(ctx, fmt.Sprintf("User-%s", user.ID), string(dataBytes), 1000)
	if err != nil {
		u.Logger.Error("Something went wrong during adding data to cache")
		u.Logger.Error("Cache operation failed but database insert was successful")
	}
	return user, nil
}

func (u *UserRepository) GetUser(ctx context.Context, userId string) (models.User, error) {
	cacheKey := "User-" + userId
	exists, err := u.Caching.ExistData(ctx, cacheKey)
	if err != nil {
		u.Logger.Error("Something went wrong during checking cache existence", userId)
		return models.User{}, models.NewError(400, "Cache", "Something went wrong during checking cache existence")
	}
	if exists > 0 {
		data, err := u.Caching.GetData(ctx, cacheKey)
		if err != nil {
			u.Logger.Error("Something went wrong during getting user from cache", userId)
			return models.User{}, models.NewError(400, "Cache", "Something went wrong during getting user from cache")
		}
		var user models.User
		err = json.Unmarshal([]byte(data), &user)
		if err != nil {
			u.Logger.Error("Failed to unmarshal user", err)
			return models.User{}, models.NewError(400, "Cache", "Failed to unmarshal user")
		}
		return user, nil
	}
	res, err := u.DbClient.Collection("Users").Find(ctx, bson.D{{"ID", userId}})
	if err != nil {
		u.Logger.Error("Something went wrong during inserting to database")
		return models.User{}, models.NewError(400, "Database", "Something went wrong during inserting to database")
	}
	var user models.User
	err = res.Decode(&user)
	if err != nil {
		u.Logger.Error("Failed to decode user", err)
		return models.User{}, models.NewError(400, "Database", "Failed to decode user")
	}
	dataBytes, err := json.Marshal(user)
	if err != nil {
		u.Logger.Error("Failed to marshal data for caching")
		return models.User{}, models.NewError(500, "Cache", "Failed to marshal data for caching")
	}
	err = u.Caching.SetData(ctx, cacheKey, string(dataBytes), 1000)
	if err != nil {
		u.Logger.Error("Something went wrong during adding data to cache")
		u.Logger.Error("Cache operation failed but database insert was successful")
	}
	return user, nil
}
