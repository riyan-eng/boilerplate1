package config

import (
	"context"
	"fmt"
	"log"

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
		log.Fatal("can't ping to redis")
	}
	fmt.Println("connection opened to redis")
}
