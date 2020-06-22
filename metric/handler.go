package metric

import (
	"local/apcupsd_exporter/apc"

	"github.com/prometheus/client_golang/prometheus"
	promLog "github.com/prometheus/common/log"
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

// StatusComponentHandler ..
type StatusComponentHandler struct {
	Handler
}

// Handle ..
func (h StatusComponentHandler) Handle(metric *Metric, output *apc.Output) {
	currentFlags := map[string]uint64{}
	if raw, exists := output.Parsed["STATFLAG"]; exists {
		currentFlags = apc.ParseStatFlags(raw)
		promLog.Infoln("  currentFlags", currentFlags)
	}

	metric.Register()
	if gaugeVec, ok := metric.Collector.(*prometheus.GaugeVec); ok {
		for flagName, flagVal := range currentFlags {
			gaugeVec.WithLabelValues(flagName).Set(float64(flagVal))
		}
	}
}
