package main

import (
	"flag"
	"local/apcupsd_exporter/metric"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	promLog "github.com/prometheus/common/log"
)

func main() {
	listenAddr := flag.String("listen", "0.0.0.0:8001", "ip:port")
	upsAddr := flag.String("ups", "127.0.0.1:3551", "apcupsd host:port")
	apcaccessPath := flag.String("apcaccess", "/sbin/apcaccess", "apcaccess path")
	flag.Parse()

	metric.RegisterPermanents()
	go metric.CollectLoop(*apcaccessPath, *upsAddr, 5*time.Second)

	promLog.Infoln("Starting exporter at", *listenAddr)
	promLog.Infoln("Watching ups", *upsAddr)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(*listenAddr, nil)
}
