package server

import (
	"local/apcupsd_exporter/metric"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	promLog "github.com/prometheus/common/log"
)

func ListenAndServe(listenAddr string, upsAddr string, apcaccessPath string) {
	metric.Register()
	go metric.CollectMetricsLoop(apcaccessPath, upsAddr, 20*time.Second)

	promLog.Infoln("Starting exporter at", listenAddr)
	promLog.Infoln("Watching ups", upsAddr)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(listenAddr, nil)
}
