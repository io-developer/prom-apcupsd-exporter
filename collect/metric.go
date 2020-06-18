package collect

import "github.com/prometheus/client_golang/prometheus"

// Metric struct
type Metric struct {
	Gauge        prometheus.Gauge
	IsPermanent  bool
	DefaultValue float64
	OutputKey    string

	isRegistered bool
}

// Register method
func (m *Metric) Register() {
	if !m.isRegistered {
		m.isRegistered = true
		prometheus.MustRegister(m.Gauge)
	}
}

// Unregister method
func (m *Metric) Unregister() {
	if m.isRegistered {
		m.isRegistered = false
		prometheus.Unregister(m.Gauge)
	}
}
