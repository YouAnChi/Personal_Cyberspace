package metrics

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// HTTPRequestDuration 请求延迟直方图
	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP请求延迟统计",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
		},
		[]string{"method", "path", "status"},
	)

	// HTTPRequestTotal 请求总数计数器
	HTTPRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "HTTP请求总数",
		},
		[]string{"method", "path", "status"},
	)

	// DBOperationDuration 数据库操作延迟直方图
	DBOperationDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_operation_duration_seconds",
			Help:    "数据库操作延迟统计",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5},
		},
		[]string{"operation"},
	)

	// DBOperationTotal 数据库操作总数计数器
	DBOperationTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "db_operations_total",
			Help: "数据库操作总数",
		},
		[]string{"operation"},
	)
)

// InitMetrics 初始化指标收集器
func InitMetrics() {
	// 注册指标收集器
	prometheus.MustRegister(HTTPRequestDuration)
	prometheus.MustRegister(HTTPRequestTotal)
	prometheus.MustRegister(DBOperationDuration)
	prometheus.MustRegister(DBOperationTotal)
}

// MetricsMiddleware Gin中间件，用于收集HTTP请求指标
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 处理请求
		c.Next()

		// 记录请求延迟
		duration := time.Since(start).Seconds()
		status := c.Writer.Status()

		// 更新指标
		HTTPRequestDuration.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			string(rune(status)),
		).Observe(duration)

		HTTPRequestTotal.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			string(rune(status)),
		).Inc()
	}
}

// PrometheusHandler 返回Prometheus指标数据的处理函数
func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
