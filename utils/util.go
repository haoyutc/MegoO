package utils

import (
	"github.com/megoo/logger/zlog"
	"go.uber.org/zap"
	"math/rand"
	"unsafe"

	"github.com/beinan/fastid"
	"time"
)

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

const (
	letterBytes   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStr(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// 优雅统计耗时

func TimeCost(methodName, jobId string) func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		zlog.Info("time cost", zap.String("method", methodName),
			zap.String("jobId", jobId),
			zap.Duration("usedTime", tc),
		)
	}
}

// GenerateUID 使用雪花算法生成全局id
func GenerateUID() int64 {
	return fastid.CommonConfig.GenInt64ID()
}
