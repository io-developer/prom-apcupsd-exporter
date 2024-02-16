package old_server

import (
	"github.com/go-kit/kit/log"
	"github.com/io-developer/prom-apcupsd-exporter/pkg/old_metric"

	"github.com/prometheus/common/promlog"
)

var (
	logger    = promlog.New(&promlog.Config{})
	collector *old_metric.Collector
)

// Init ..
func Init(l log.Logger, c *old_metric.Collector) {
	logger = l
	collector = c

	metricsInit()
}
