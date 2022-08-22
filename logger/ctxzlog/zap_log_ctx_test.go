package ctxzlog

import (
	"github.com/gin-gonic/gin"
	"github.com/megoo/logger/zlog"
	"go.uber.org/zap"
	"testing"
)

func TestZapLog(t *testing.T) {
	ctx := gin.Context{}
	name := ctx.Query("name")
	WithContext(&ctx).Info("Test logger", zap.String("name", name))
	zlog.Info("test info log", zap.String("test", name))
}
