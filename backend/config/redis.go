package config

import (
	"exchangeapp/global"
	"log"

	"github.com/go-redis/redis"
)

func InitRedis() {
	// Initialize redis
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// Ping redis
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis, got error: %v", err)
	}
	// Set redis client
	global.RedisDB = RedisClient
}
