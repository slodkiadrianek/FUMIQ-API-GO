package config

import (
	"FUMIQ_API/utils"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type CacheService struct {
	Client *redis.Client
	Logger utils.Logger
}

func ConnectToCache(cacheLink string) *CacheService {
	opt, err := redis.ParseURL(cacheLink)
	if err != nil {
		panic(err)
	}
	return &CacheService{
		Client: redis.NewClient(opt),
	}
}

func (c *CacheService) SetData(ctx context.Context, key string, data string, time time.Duration) error {
	err := c.Client.Set(ctx, key, data, time).Err()
	if err != nil {
		return err
	}
	return nil
}
func (c *CacheService) ExistData(ctx context.Context, key string) (int64, error) {
	result, err := c.Client.Exists(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (c *CacheService) GetData(ctx context.Context, key string) (string, error) {
	result, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (c *CacheService) DeleteData(ctx context.Context, key string) error {
	err := c.Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
