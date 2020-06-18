package server

import (
	"local/apcupsd_exporter/collect"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	promLog "github.com/prometheus/common/log"
)

// ListenAndServe ..
func ListenAndServe(listenAddr string, upsAddr string, apcaccessPath string) {
	collect.MetricsRegister()
	go collect.CollectMetricsLoop(apcaccessPath, upsAddr, 5*time.Second)

	promLog.Infoln("Starting exporter at", listenAddr)
	promLog.Infoln("Watching ups", upsAddr)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(listenAddr, nil)
}
