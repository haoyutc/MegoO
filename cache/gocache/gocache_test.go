package gocache

import (
	"github.com/bytedance/sonic"
	"github.com/megoo/logger/zlog"
	"github.com/megoo/utils"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"testing"
	"time"
)

type Person struct {
	ID   int64
	Name string
	Age  int
}

// 字节JSON序列化库
func TestSonic(t *testing.T) {
	var data Person
	output, err := sonic.Marshal(&data)
	if err != nil {
		zlog.Error("Sonic marshal err", zap.Error(err))
	}
	zlog.Info("Marshal result", zap.Any("output", output))
	err = sonic.Unmarshal(output, &data)
	if err != nil {
		zlog.Error("Sonic unmarshal err", zap.Error(err))
	}
	zlog.Info("Unmarshal result", zap.Any("data", data))
}

type Client struct {
	uinCache      *cache.Cache
	notExistCache *cache.Cache
	goLimit       *utils.Limiter
}

func NewAccountClient() *Client {
	return &Client{
		// appid -> uin，永远会变
		uinCache: cache.New(cache.NoExpiration, 10*time.Minute),
		// 不存在缓存设置为5分钟
		notExistCache: cache.New(5*time.Minute, 1*time.Minute),
		// 单节点限制10qps访问
		goLimit: utils.NewLimiter(10),
	}
}

func TestGoCache(t *testing.T) {

	ac := NewAccountClient()

	if val, ok := ac.uinCache.Get("key"); ok {
		zlog.Info("Get data from cache", zap.String("uin", val.(string)))

	}

	if _, ok := ac.notExistCache.Get("key"); ok {
	}
	ac.goLimit.Add()
	defer ac.goLimit.Done()
	// TODO 业务

}
