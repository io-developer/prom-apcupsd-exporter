package server

import (
	"local/apcupsd_exporter/metric"

	"github.com/go-kit/kit/log"

	"github.com/prometheus/common/promlog"
)

var (
	logger    = promlog.New(&promlog.Config{})
	collector *metric.Collector
)

// Init ..
func Init(l log.Logger, c *metric.Collector) {
	logger = l
	collector = c

	metricsInit()
	wsInit()
	signalsInit()
}
