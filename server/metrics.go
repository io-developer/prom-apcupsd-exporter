package server

import (
	"local/apcupsd_exporter/metric"
	"net/http"

	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	collector   *metric.Collector
	promHandler = promhttp.Handler()
)

// RegisterMetricEndpoints ..
func RegisterMetricEndpoints(c *metric.Collector) {
	collector = c
	http.HandleFunc("/metrics", handleMetrics)
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	onComplete := make(chan bool)
	collector.Collect(metric.CollectOpts{
		PreventFlood: true,
		OnComplete:   onComplete,
	})
	if <-onComplete {
		level.Debug(Logger).Log("msg", "handleMetrics begin")
		promHandler.ServeHTTP(w, r)
		level.Debug(Logger).Log("msg", "handleMetrics end")
	}
}
