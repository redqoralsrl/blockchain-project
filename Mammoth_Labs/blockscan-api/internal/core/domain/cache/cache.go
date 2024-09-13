package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	MGet(ctx context.Context, keys ...string) *redis.SliceCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Exists(ctx context.Context, keys ...string) *redis.IntCmd
	FlushAll(ctx context.Context) *redis.StatusCmd
	Keys(ctx context.Context, pattern string) *redis.StringSliceCmd
	Close() error
}

type Repository interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (interface{}, error)
	MGet(keys []string) ([]interface{}, error)
	Delete(key string) error
	Exists(key string) (bool, error)
	Flush() error
	Keys(key string) ([]string, error)
}

type UseCase interface {
	// 캐시에 데이터를 추가하거나 업데이트
	AddOrUpdateItem(key string, value interface{}, expiration time.Duration) error

	//캐시에서 데이터를 검색
	RetrieveItem(key string, dest interface{}) error

	// 캐시에서 여러 데이터를 검색
	RetrieveMultiItem(keys []string, dest interface{}) error

	// 캐시에서 특정 데이터를 무효화
	InvalidateItem(key string) error

	//주어진 키가 캐시에 존재하는지 여부를 확인
	ExistsItem(key string) (bool, error)

	// 캐시의 모든 데이터를 무효화
	InvalidateAll() error

	// 키값에 특정 값을 포함하고 있는지 확인
	FindKeys(key string) ([]string, error)
}
