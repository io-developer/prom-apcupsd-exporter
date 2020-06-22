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
	listenAddr := flag.String("listen", "0.0.0.0:8001",
		"ip:port",
	)
	upsAddr := flag.String("ups", "127.0.0.1:3551",
		"apcupsd host:port",
	)
	apcaccessPath := flag.String("apcaccess", "/sbin/apcaccess",
		"apcaccess path",
	)
	collectMinInterval := flag.Float64("collectMinInterval", 0.5,
		"Min time delta between Collect calls in seconds. Prevents /metrics flooding",
	)
	collectLoopInterval := flag.Float64("collectLoopInterval", 10,
		"Base Collect loop interval in seconds",
	)
	flag.Parse()

	promLog.Infoln("Apcupsd server addr:", *upsAddr)
	promLog.Infoln("Apcaccess bin path:", *apcaccessPath)
	promLog.Infoln("Min interval between collect calls:", *collectMinInterval, "sec")
	promLog.Infoln("Loop interval between collect calls:", *collectLoopInterval, "sec")

	metric.ApcaccessPath = *apcaccessPath
	metric.ApcupsdAddr = *upsAddr
	metric.CollectMinInterval = time.Duration(*collectMinInterval * float64(time.Second))

	promHandler := promhttp.Handler()
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		metric.Collect()
		promHandler.ServeHTTP(w, r)
	})

	metric.RegisterPermanents()
	go metric.CollectLoop(time.Duration(*collectLoopInterval * float64(time.Second)))

	promLog.Infoln("Starting exporter at", *listenAddr)
	promLog.Infoln("")
	http.ListenAndServe(*listenAddr, nil)
}
