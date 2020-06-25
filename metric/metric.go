package metric

import (
	"local/apcupsd_exporter/model"
	"math"

	"github.com/prometheus/client_golang/prometheus"
)

// Metric struct
type Metric struct {
	Collector    prometheus.Collector
	Handler      Handler
	HandlerFunc  func(m *Metric, model *model.Model)
	ValFunc      func(m *Metric, model *model.Model) float64
	DefaultValue float64
	IsPermanent  bool
	isRegistered bool
}

// Register method
func (m *Metric) Register() {
	if !m.isRegistered {
		m.isRegistered = true
		prometheus.MustRegister(m.Collector)
	}
}

// Unregister method
func (m *Metric) Unregister() {
	if m.isRegistered {
		m.isRegistered = false
		prometheus.Unregister(m.Collector)
	}
}

// UpdateCollector method
func (m *Metric) UpdateCollector(val float64) {
	if math.IsNaN(val) {
		m.Unregister()
	} else {
		m.Register()
	}
	if gauge, ok := m.Collector.(prometheus.Gauge); ok {
		gauge.Set(val)
	}
}
