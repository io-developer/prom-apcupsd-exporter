package server

import (
	"github.com/io-developer/prom-apcupsd-exporter/metric"

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
}
