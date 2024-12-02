package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var (
	MethodErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "parser_error_count",
			Help: "For error count of each method",
		}, []string{"method"},
	)
	MethodRequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "parser_request_count",
			Help: "For counting rpc and error rate",
		}, []string{"method"},
	)
	MethodRequestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "parser_request_latency",
			Help:    "For counting latency for methods",
			Buckets: prometheus.DefBuckets,
		}, []string{"method"},
	)
)

func init() {
	prometheus.MustRegister(
		MethodErrorCount,
		MethodRequestCount,
		MethodRequestLatency,
	)
}

func SinceSeconds(started time.Time) float64 {
	return float64(time.Since(started)) / float64(time.Second)
}
