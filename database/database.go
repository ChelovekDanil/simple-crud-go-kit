package database

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

const (
	defaultRedisURI = "localhost:6379"
	RedisURI        = "REDIS_URI"
)

// Connect return connection in database
func Connect(ctx context.Context) (*redis.Client, error) {
	var redisUri string
	if redisUri = os.Getenv(RedisURI); redisUri == "" {
		redisUri = defaultRedisURI
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: redisUri,
	})
	return rdb, nil
}
