package metric

import (
	"local/apcupsd_exporter/apc"

	"github.com/prometheus/client_golang/prometheus"
)

// Handler ..
type Handler interface {
	Handle(m *Metric, o *apc.Output)
}

// DefaultHandler ..
type DefaultHandler struct {
	Handler
	ApcKey string
	Mapper Mapper
}

// NewDefaultHandler ..
func NewDefaultHandler(apcKey string) DefaultHandler {
	return DefaultHandler{
		ApcKey: apcKey,
		Mapper: DefaultMapper{},
	}
}

// Handle ..
func (h DefaultHandler) Handle(metric *Metric, output *apc.Output) {
	raw, exists := output.Parsed[h.ApcKey]
	if !exists && !metric.IsPermanent {
		metric.Unregister()
		return
	}
	val := h.Mapper.Map(raw, metric.DefaultValue)
	metric.Register()
	if gauge, ok := metric.Collector.(prometheus.Gauge); ok {
		gauge.Set(val)
	}
}
