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
