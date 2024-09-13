package redis

import (
	"github.com/redis/go-redis/v9"
	"blockscan-go/internal/config"
)

func ConnectCache(config *config.Config) (cache *redis.Client) {
	client := redis.NewClient(&redis.Options{
		Addr: config.RedisHost + ":" + config.RedisPort,
	})
	cache = client
	return cache
}
