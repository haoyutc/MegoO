package metrics

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/megoo/logger/zlog"
	"go.uber.org/zap"
	"testing"
)

func TestProm(t *testing.T) {

	zlog.Info("Init metrics server")
	// create Prometheus server and middleware
	gooseProm := echo.New()
	gooseProm.HideBanner = true
	prom := prometheus.NewPrometheus(SubSystemName, nil)
	// Scrape metrics from main Server
	//s.echoHttpServer.Use(prom.HandlerFunc)
	// SetUp metrics endpoint at another server port 9966
	prom.SetMetricsPath(gooseProm)

	go func() {
		if err := gooseProm.Start(":9966"); err != nil {
			zlog.Error("Middleware metrics start failed", zap.Error(err))
		}
	}()
}
