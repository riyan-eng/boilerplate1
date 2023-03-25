package config

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnRedis() {
	addr := fmt.Sprintf("%s:%s", "localhost", "6379")
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	if err := Redis.Ping(ctx).Err(); err != nil {
		fmt.Println("redis: can't ping to redis")
		os.Exit(1)
	}
	fmt.Println("redis: connection opened to redis")
}
