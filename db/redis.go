package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	client := ExampleNewClient()
	result := client.HGetAll(ctx, "zhangsan")
	userId, err := result.Result()
	if err == nil {
		fmt.Print(userId["1"])
	}
	//fmt.Print(result)
}

func ExampleNewClient() redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.11.120:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return *rdb
}
