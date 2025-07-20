package utils

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis() error {
	redisURL := os.Getenv("UPSTASH_REDIS_URL")
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return err
	}
	RedisClient = redis.NewClient(opt)

	// 測試連線
	if err := RedisClient.Ping(ctx).Err(); err != nil {
		return err
	}
	log.Println("✅ Redis connected successfully")
	return nil
}
