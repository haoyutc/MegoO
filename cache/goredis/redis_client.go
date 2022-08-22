package goredis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

func RedisClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Set(ctx, "key", "value", 0)
	if err != nil {
		panic(err)
	}

	val, err2 := rdb.Get(ctx, "key").Result()
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("key: %s", val)

	val2, err3 := rdb.Get(ctx, "key2").Result()
	if err3 == redis.Nil {
		fmt.Println("key2 is not exist")
	} else if err3 != nil {
		panic(err3)
	} else {
		fmt.Printf("key2: %s", val2)
	}

	rdb.SetNX(ctx, "key3", "expire", 10*time.Second).Result()
	rdb.SetNX(ctx, "key", "value", redis.KeepTTL)

}
