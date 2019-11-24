package internal

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func GetRedisClient() *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	if len(redisHost) == 0 {
		redisHost = "localhost:6379"
	}

	fmt.Printf("Using redis host %s\n", redisHost)

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}
