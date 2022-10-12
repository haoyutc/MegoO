package cache_test

import (
	"context"
	"fmt"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func Test_basicUsage(t *testing.T) {
	ring := redis.NewRing(&redis.RingOptions{Addrs: map[string]string{
		"server1": ":6379",
		"server2": ":6380",
	}})

	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	ctx := context.TODO()
	key := "person:lisi"
	lisi := &Person{
		Name: "lisi",
		Age:  24,
	}
	if err := mycache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: lisi,
		TTL:   time.Hour,
	}); err != nil {
		panic(err)
	}

	var wanted Person
	if err := mycache.Get(ctx, key, &wanted); err == nil {
		fmt.Println(wanted)
	}
}

func Test_AdvancedUsage(t *testing.T) {
	ring := redis.NewRing(&redis.RingOptions{Addrs: map[string]string{
		"server1": ":6379",
		"server2": ":6380",
	}})
	mycache := cache.New(&cache.Options{Redis: ring, LocalCache: cache.NewTinyLFU(1000, time.Minute)})

	lisi := new(Person)
	err := mycache.Once(&cache.Item{
		Key:   "lisi",
		Value: lisi,
		Do: func(item *cache.Item) (interface{}, error) {
			return &Person{
				Name: "Lisi",
				Age:  24,
			}, nil
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(lisi)
	// Output: &{Lisi 24}
}
