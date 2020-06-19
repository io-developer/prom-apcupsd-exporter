package collect

import (
	"local/apcupsd_exporter/apc"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

// Metric struct
type Metric struct {
	Gauge        prometheus.Gauge
	IsPermanent  bool
	DefaultValue float64
	OutputKey    string
	Type         string // "", date, minutes, valueMap
	ValueMap     map[string]float64

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

// GetOutputValue method
func (m *Metric) GetOutputValue(output *apc.Output) (num float64, exists bool) {
	str, exists := output.Parsed[m.OutputKey]
	if !exists {
		return m.DefaultValue, false
	}
	if m.Type == "" {
		val, _ := strconv.ParseFloat(str, 64)
		return val, true
	}
	if m.Type == "minutes" {
		val, _ := strconv.ParseFloat(str, 64)
		return val * 60, true
	}
	if m.Type == "date" {
		return -1, true
	}
	if m.Type == "valueMap" && m.ValueMap != nil {
		mapped, exists := m.ValueMap[str]
		if exists {
			return mapped, true
		}
	}
	return m.DefaultValue, true
}
