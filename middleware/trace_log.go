package middleware

import (
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/megoo/logger/ctxzlog"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
)

func TraceLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 每个请求生成的请求traceId具有全局唯一性
		uid := uuid.NewV4()
		traceId := uid.String()
		ctxzlog.NewContext(ctx, zap.String("traceId", traceId))

		// 为日志添加请求的地址以及请求参数等信息
		ctxzlog.NewContext(ctx, zap.String("request.method", ctx.Request.Method))
		headers, _ := sonic.Marshal(ctx.Request.Header)
		ctxzlog.NewContext(ctx, zap.String("request.headers", string(headers)))
		ctxzlog.NewContext(ctx, zap.String("request.url", ctx.Request.URL.String()))

		// 将请求参数json序列化后添加进日志上下文
		if ctx.Request.Form == nil {
			ctx.Request.ParseMultipartForm(32 << 20)
		}
		form, _ := sonic.Marshal(ctx.Request.Form)
		ctxzlog.NewContext(ctx, zap.String("request.params", string(form)))

		ctx.Next()
	}
}
