package utils

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var Cache *redis.ClusterClient

func InitCache() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: Config.Redis.Address,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("failed to connect cluster client: %v", err)
	}

	log.Println("success connect redis")
	Cache = client
}
