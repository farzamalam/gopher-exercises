package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()

	fmt.Println(pong, err)
}
