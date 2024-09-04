package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	Records  prometheus.Gauge
	Registry *prometheus.Registry
	Duration *prometheus.HistogramVec
}

func NewMetrics() *Metrics {
	registry := prometheus.NewRegistry()

	// Create Prometheus metrics.
	m := &Metrics{
		Records: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "external_dns_unifi_webhook",
			Name:      "records_created",
			Help:      "Amount of records created.",
		},
		),
		Duration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "external_dns_unifi_webhook",
			Name:      "request_duration_seconds",
			Help:      "Duration of the request",
			Buckets:   []float64{0.1, 0.15, 0.2, 0.25, 0.3},
		}, []string{"status", "route", "method"},
		),
		Registry: registry,
	}
	// Register metrics with Prometheus registry.
	registry.MustRegister(m.Records, m.Duration)

	return m
}
