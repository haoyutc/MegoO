package metrics

import (
	"github.com/megoo/logger/zlog"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"strconv"
	"time"
)

const SubSystemName = "MegoO"

// Custom metrics
//	counter, counter_vec, gauge, gauge_vec,
//	histogram, histogram_vec, summary, summary_vec

var apiOpsProcessed = prometheus.NewCounterVec(prometheus.CounterOpts{
	Subsystem: SubSystemName,
	Name:      "processed_api_ops_total",
	Help:      "The total number of processed api events",
}, []string{"http_code"})

func ObserveApiOpsWithHttpCode(httpCode int) {
	apiOpsProcessed.WithLabelValues(strconv.Itoa(httpCode)).Inc()
}

var taskStatusProcessed = prometheus.NewCounterVec(prometheus.CounterOpts{
	Subsystem: SubSystemName,
	Name:      "processed_task_ops_total",
	Help:      "The total number of processed task events",
}, []string{"status"})

func ObserveTaskStatus(status string) {
	taskStatusProcessed.WithLabelValues(status).Inc()
}

var unArchiveOpsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
	Subsystem: SubSystemName,
	Name:      "unarchive_ops_count",
	Help:      "The total number of processed unarchive events",
}, []string{"status"})

func ObserveUnArchiveProcessed(status string) {
	unArchiveOpsCounter.WithLabelValues(status).Inc()
}

var moduleTimeHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Subsystem: SubSystemName,
	Name:      "module_time_used",
	Help:      "This is my histogram",
}, []string{"module"})

func ObserveModuleTimeCost(moduleName string) func() {
	start := time.Now()
	return func() {
		moduleTimeHistogram.WithLabelValues(moduleName).Observe(time.Since(start).Seconds())
	}
}

var metrics = []prometheus.Collector{
	apiOpsProcessed,
	taskStatusProcessed,
	moduleTimeHistogram,
	unArchiveOpsCounter,
}

// 初始化自定义指标

func init() {
	// 注册自定义指标
	RegistryCustomPromCollector(metrics)
}

func RegistryCustomPromCollector(cols []prometheus.Collector) {
	for _, col := range cols {
		if err := prometheus.Register(col); err != nil {
			zlog.Error("This metrics could not be registered in Prometheus", zap.Error(err))
		}
	}
}
