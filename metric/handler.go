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

// StatusComponentHandler ..
type StatusComponentHandler struct {
	Handler
}

// Handle ..
func (h StatusComponentHandler) Handle(metric *Metric, output *apc.Output) {
	currentFlags := apc.ParseStatFlags(output.GetParsed("STATFLAG", ""))

	metric.Register()
	if gaugeVec, ok := metric.Collector.(*prometheus.GaugeVec); ok {
		for flagName, flagVal := range currentFlags {
			gaugeVec.WithLabelValues(flagName).Set(float64(flagVal))
		}
	}
}

// StatusTraceComponentHandler ..
type StatusTraceComponentHandler struct {
	Handler
}

// Handle ..
func (h StatusTraceComponentHandler) Handle(metric *Metric, output *apc.Output) {
	currentFlags := apc.ParseStatFlags(output.GetParsed("STATFLAG", ""))

	metric.Register()
	if summaryVec, ok := metric.Collector.(*prometheus.SummaryVec); ok {
		for flagName, flagVal := range currentFlags {
			if flagVal > 0 {
				summaryVec.WithLabelValues(flagName).Observe(float64(flagVal))
			}
		}
	}
}
