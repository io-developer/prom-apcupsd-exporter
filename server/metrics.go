package server

import (
	"local/apcupsd_exporter/metric"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	promLog "github.com/prometheus/common/log"
)

var promHandler = promhttp.Handler()

func HandleMetrics(w http.ResponseWriter, r *http.Request) {
	onComplete := make(chan bool)
	metric.CollectChan <- metric.CollectOpts{
		PreventFlood: true,
		OnComplete:   onComplete,
	}
	if <-onComplete {
		promLog.Infoln("ServeHTTP start")
		promHandler.ServeHTTP(w, r)
		promLog.Infoln("ServeHTTP end")
	}
}
