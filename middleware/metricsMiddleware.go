package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vamshireddy02/registry/metrics"
)

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		metrics.RequestCounter.WithLabelValues(c.Request.Method, c.Request.URL.Path).Inc()

		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		metrics.ResponseDurationHistogram.WithLabelValues(c.Request.Method, c.Request.URL.Path).Observe(duration)

		metrics.ResponseCounter.WithLabelValues(fmt.Sprintf("%d", c.Writer.Status())).Inc()
	}
}
