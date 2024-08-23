package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics holds all the Prometheus metrics for the application.
type Metrics struct {
	RecordsCreated        prometheus.Gauge
	TotalRequests         prometheus.Counter
	ResponseTimeHistogram prometheus.Histogram
}

// NewMetrics creates and registers all Prometheus metrics and returns a Metrics struct.
func NewMetrics(registry *prometheus.Registry) *Metrics {
	m := &Metrics{
		RecordsCreated: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "unifi_records_created",
			Help: "Current number of records created in Unifi",
		}),
		TotalRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: "unifi_requests_count",
			Help: "Total number of requests",
		}),
		ResponseTimeHistogram: promauto.NewHistogram(prometheus.HistogramOpts{
			Name:    "unifi_response_time_seconds",
			Help:    "Response time in seconds",
			Buckets: prometheus.DefBuckets,
		}),
	}
	registry.MustRegister(m.RecordsCreated)
	registry.MustRegister(m.TotalRequests)
	registry.MustRegister(m.ResponseTimeHistogram)

	return m
}
