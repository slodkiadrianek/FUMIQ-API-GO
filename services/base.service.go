package services

import (
	"FUMIQ_API/models"
	"FUMIQ_API/utils"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"FUMIQ_API/config"
)

type BaseService struct {
	Logger   *utils.Logger
	Caching  *config.CacheService
	DbClient *mongo.Database
}

func (b *BaseService) InsertToDatabaseAndCache(ctx context.Context, cache string, data interface{},
	collection string) error {
	stringData, _ := json.Marshal(data)
	err := b.Caching.SetData(ctx, cache, string(stringData), 1000)
	if err != nil {
		b.Logger.Error("Something went wrong during adding to data to cache")
		err := models.NewError(400, "Cache", "Something went wrong during adding to data to cache")
		return err
	}
	_, err = b.DbClient.Collection(collection).InsertOne(ctx, data)
	if err != nil {
		b.Logger.Error("Something went wrong during inserting to database")
		err := models.NewError(400, "Database", "Something went wrong during inserting to database")
		return err
	}
	return nil
}

func (b *BaseService) GetAllUserItems(ctx context.Context, cache string, userId string, collection string,
	dataType *interface{}) (interface{}, error) {

	cacheKey := cache + ":" + userId

	exists, err := b.Caching.ExistData(ctx, cacheKey)
	if err != nil {
		b.Logger.Error("Something went wrong during checking cache existence", userId)
		return nil, models.NewError(400, "Cache", "Something went wrong during checking cache existence")
	}

	if exists > 0 {
		data, err := b.Caching.GetData(ctx, cacheKey)
		if err != nil {
			b.Logger.Error("Something went wrong during getting user items from cache", userId)
			return nil, models.NewError(400, "Cache", "Something went wrong during getting user items from cache")
		}

		if err := json.Unmarshal([]byte(data), *dataType); err != nil {
			b.Logger.Error("Failed to unmarshal user items", err)
			return nil, models.NewError(400, "Cache", "Failed to unmarshal user items")
		}
		return &dataType, nil
	}
	res, err := b.DbClient.Collection(collection).Find(ctx, bson.D{{"userId", userId}})
	if err != nil {
		b.Logger.Error("Something went wrong during inserting to database")
		err := models.NewError(400, "Database", "Something went wrong during inserting to database")
		return nil, err
	}
	return res.Decode(&dataType), nil
}
