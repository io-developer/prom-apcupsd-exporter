package server

import (
	"local/apcupsd_exporter/metric"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	promLog "github.com/prometheus/common/log"
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
		promLog.Infoln("ServeHTTP start")
		promHandler.ServeHTTP(w, r)
		promLog.Infoln("ServeHTTP end")
	}
}
