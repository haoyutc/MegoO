package main

import (
	"context"
	"fmt"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v9"
	"log"
	"time"
)

func main() {
	// connect to redis.
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379", Network: "tcp"})
	defer client.Close()

	// create a new lock client
	locker := redislock.New(client)
	ctx := context.Background()

	// try to obtain lock
	lock, err := locker.Obtain(ctx, "lock_key", 100*time.Millisecond, nil)
	if err == redislock.ErrNotObtained {
		fmt.Println("Could not obtain lock!")
	} else if err != nil {
		log.Fatal(err)
	}

	// Don't forget to defer release.
	defer lock.Release(ctx)
	fmt.Println("I have a lock")

	// Sleep and check the remaining TTL.
	time.Sleep(time.Millisecond * 50)
	if ttl, err := lock.TTL(ctx); err != nil {
		log.Fatal(err)
	} else if ttl > 0 {
		fmt.Println("Yay, I still have my lock!")
	}

	// Extend my lock
	if err := lock.Refresh(ctx, 100*time.Millisecond, nil); err != nil {
		log.Fatal(err)
	}

	// Sleep a little longer, then check.
	time.Sleep(time.Millisecond * 100)
	if ttl, err := lock.TTL(ctx); err != nil {
		log.Fatal(err)
	} else if ttl == 0 {
		fmt.Println("Now, my lock has expired!")
	}

}
