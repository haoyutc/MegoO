package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_ = rdb.FlushDB(ctx).Err()
	//if err != nil {
	//	panic(err)
	//}
	limiter := redis_rate.NewLimiter(rdb)
	res, err := limiter.Allow(ctx, "demo:1234", redis_rate.PerSecond(10))
	if err != nil {
		panic(err)
	}
	fmt.Println("allowed", res.Allowed, "remaining", res.Remaining)
	// output: allowed 1 remaining 9
}
