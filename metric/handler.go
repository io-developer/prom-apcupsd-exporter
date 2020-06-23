package metric

import (
	"local/apcupsd_exporter/apc"
	"local/apcupsd_exporter/model"
	"math"

	"github.com/prometheus/client_golang/prometheus"
)

// Handler ..
type Handler interface {
	Handle(m *Metric, o *apc.Output)
}

// DefaultHandler ..
type DefaultHandler struct {
	Handler
	ApcKey   string
	ValueMap map[string]float64
}

// NewDefaultHandler ..
func NewDefaultHandler(apcKey string) DefaultHandler {
	return DefaultHandler{ApcKey: apcKey}
}

// Handle ..
func (h DefaultHandler) Handle(metric *Metric, output *apc.Output) {
	raw := output.GetParsed(h.ApcKey, "")
	if raw == "" && !metric.IsPermanent {
		metric.Unregister()
		return
	}

	val := metric.DefaultValue
	if h.ValueMap != nil {
		if mapped, exists := h.ValueMap[raw]; exists {
			val = mapped
		}
	} else if parsed, err := apc.ParseCommon(raw); err == nil {
		val = parsed
	}

	metric.Register()
	if gauge, ok := metric.Collector.(prometheus.Gauge); ok {
		gauge.Set(val)
	}
}

// StatusHandler ..
type StatusHandler struct {
	Handler
}

// Handle ..
func (h StatusHandler) Handle(metric *Metric, output *apc.Output) {
	currentFlags := apc.ParseFlags(output.GetParsed("STATFLAG", ""), model.StatusFlags)

	metric.Register()
	if gaugeVec, ok := metric.Collector.(*prometheus.GaugeVec); ok {
		for flagName, flagVal := range currentFlags {
			gaugeVec.WithLabelValues(flagName).Set(math.Min(1.0, float64(flagVal)))
		}
	}
}

// StatusHexHandler ..
type StatusHexHandler struct {
	Handler
}

// Handle ..
func (h StatusHexHandler) Handle(metric *Metric, output *apc.Output) {
	currentFlags := apc.ParseFlags(output.GetParsed("STATFLAG", ""), model.StatusFlags)

	metric.Register()
	if gaugeVec, ok := metric.Collector.(*prometheus.GaugeVec); ok {
		for flagName, flagVal := range currentFlags {
			gaugeVec.WithLabelValues(flagName).Set(float64(flagVal))
		}
	}
}
