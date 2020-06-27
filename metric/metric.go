package metric

import (
	"local/apcupsd_exporter/model"
	"math"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

// Metric struct
type Metric struct {
	Collector    prometheus.Collector
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
		prometheus.Register(m.Collector)
	}
}

// Unregister method
func (m *Metric) Unregister() {
	if m.isRegistered {
		m.isRegistered = false
		prometheus.Unregister(m.Collector)
	}
}

// Update method
func (m *Metric) Update(curModel *model.Model) {
	if m.IsPermanent {
		m.Register()
	}
	if m.HandlerFunc != nil {
		m.HandlerFunc(m, curModel)
	} else if m.ValFunc != nil {
		m.UpdateCollector(m.ValFunc(m, curModel))
	}
}

// UpdateCollector method
func (m *Metric) UpdateCollector(val float64) {
	if !m.IsPermanent && math.IsNaN(val) {
		m.Unregister()
	} else {
		m.Register()
	}

	if gauge, ok := m.Collector.(prometheus.Gauge); ok {
		gauge.Set(val)
	} else if counter, ok := m.Collector.(prometheus.Counter); ok {
		curData := &dto.Metric{}
		counter.Write(curData)
		delta := val - *curData.Counter.Value
		if int64(delta) > 0 {
			counter.Add(delta)
		}
	}
}
