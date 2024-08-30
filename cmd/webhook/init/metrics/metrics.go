package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	Records  prometheus.Gauge
	Registry *prometheus.Registry
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
		Registry: registry,
	}
	// Register metrics with Prometheus registry.
	registry.MustRegister(m.Records)

	return m
}
