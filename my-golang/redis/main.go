package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// Replace options as needed (Addr, Password, DB)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password by default
		DB:       0,  // use default DB
	})

	// Optional: ping to verify connection
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(fmt.Errorf("redis ping failed: %w", err))
	}

	// Example: set and get a key
	if err := rdb.Set(ctx, "example:key", "hello redis", 10*time.Minute).Err(); err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "example:key").Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("value:", val)
	}

	// Close when done
	if err := rdb.Close(); err != nil {
		fmt.Println("close error:", err)
	}
}
