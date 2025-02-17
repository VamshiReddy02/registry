package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "path"},
	)

	ResponseDurationHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_duration_seconds",
			Help:    "Histogram of response durations for HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	ResponseCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_responses_total",
			Help: "Total number of HTTP responses.",
		},
		[]string{"status"},
	)
)

func init() {
	prometheus.MustRegister(RequestCounter)
	prometheus.MustRegister(ResponseDurationHistogram)
	prometheus.MustRegister(ResponseCounter)
}
