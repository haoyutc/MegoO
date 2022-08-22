package utils

import (
	"github.com/megoo/logger/zlog"
	"go.uber.org/zap"

	"github.com/beinan/fastid"
	"time"
)

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
