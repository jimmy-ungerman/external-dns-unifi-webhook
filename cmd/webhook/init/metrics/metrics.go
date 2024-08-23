package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// New metric for device count
	RecordCountGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_device_count",
		Help: "Current number of Records",
	})
)
