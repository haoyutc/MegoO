package redismock

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/juju/errors"
	"testing"
	"time"
)

var ctx = context.TODO()

func NewsInfoForCache(rdb *redis.Client, newsId int) (info string, err error) {
	cacheKey := fmt.Sprintf("news_redis_cache_%d", newsId)
	info, err = rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		info = "test"
		err = rdb.Set(ctx, cacheKey, info, 30*time.Minute).Err()
	}
	return
}
func TestNewsInfoForCache(t *testing.T) {
	db, mock := redismock.NewClientMock()
	newsId := 123456789
	key := fmt.Sprintf("news_redis_cache_%d", newsId)

	mock.ExpectGet(key).RedisNil()
	mock.Regexp().ExpectSet(key, `[a-z]+`, 30*time.Minute).SetErr(errors.New("FAIL"))

	if _, err := NewsInfoForCache(db, newsId); err == nil || err.Error() != "FAIL" {
		t.Error("wrong error")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
