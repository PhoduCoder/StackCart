package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Locally running docker redis container
	})

	if err := client.Set(ctx, "name", "Gaurav", 0).Err(); err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}

// Create a client
// Set a key-value pair in Redis
// Get the value of the key from Redis
