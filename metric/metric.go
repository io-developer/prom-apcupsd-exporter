package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Metric struct
type Metric struct {
	Collector    prometheus.Collector
	Handler      Handler
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
