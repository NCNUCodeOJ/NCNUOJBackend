package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis 服務的位址
		Password: "",               //
		DB:       0,                // 對應 reids 0-15 的 db，測試連接
	})

	rdb.HSet(ctx, "user", "key1", "value1", "key2", "value2")
	rdb.HSet(ctx, "user", []string{"key3", "value3", "key4", "value4"})
	rdb.HSet(ctx, "user", map[string]interface{}{"key5": "value5", "key6": "value6"})
	user, err := rdb.HGetAll(ctx, "user").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
