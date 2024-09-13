package redis

import (
	"context"
	"blockscan-go/internal/core/domain/cache"
	"time"
)

type CacheRepository struct {
	client cache.RedisClient
}

func NewRedisCacheRepository(client cache.RedisClient) *CacheRepository {
	return &CacheRepository{client: client}
}

func (r *CacheRepository) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(context.Background(), key, value, expiration).Err()
}

func (r *CacheRepository) Get(key string) (interface{}, error) {
	return r.client.Get(context.Background(), key).Result()
}

func (r *CacheRepository) MGet(keys []string) ([]interface{}, error) {
	return r.client.MGet(context.Background(), keys...).Result()
}

func (r *CacheRepository) Delete(key string) error {
	return r.client.Del(context.Background(), key).Err()
}

func (r *CacheRepository) Exists(key string) (bool, error) {
	result, err := r.client.Exists(context.Background(), key).Result()
	return result > 0, err
}

func (r *CacheRepository) Flush() error {
	return r.client.FlushAll(context.Background()).Err()
}

func (r *CacheRepository) Keys(key string) ([]string, error) {
	result, err := r.client.Keys(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}
