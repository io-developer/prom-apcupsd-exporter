package main

import (
	"flag"
	"local/apcupsd_exporter/metric"
	"local/apcupsd_exporter/model"
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
	apcaccessFloodLimit := flag.Float64("apcaccessFloodLimit", 0.5,
		"Min time delta between Collect calls in seconds. Prevents /metrics flooding",
	)
	collectLoopInterval := flag.Float64("collectLoopInterval", 10,
		"Base Collect loop interval in seconds",
	)
	flag.Parse()

	promLog.Infoln("Apcupsd server addr:", *upsAddr)
	promLog.Infoln("Apcaccess bin path:", *apcaccessPath)
	promLog.Infoln("Min interval between apcaccess calls:", *apcaccessFloodLimit, "sec")
	promLog.Infoln("Loop interval between collect calls:", *collectLoopInterval, "sec")

	metric.Channel = make(chan metric.Opts)
	metric.CurrentModel = model.NewModel()
	metric.ApcupsdAddr = *upsAddr
	metric.ApcaccessPath = *apcaccessPath
	metric.ApcaccessFloodLimit = time.Duration(*apcaccessFloodLimit * float64(time.Second))

	promHandler := promhttp.Handler()
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		onComplete := make(chan bool)
		metric.Channel <- metric.Opts{
			PreventFlood: true,
			OnComplete:   onComplete,
		}
		if <-onComplete {
			promLog.Infoln("ServeHTTP start")
			promHandler.ServeHTTP(w, r)
			promLog.Infoln("ServeHTTP end")
		}
	})

	metric.RegisterPermanents()

	go metric.Collect(metric.Channel)
	go metric.CollectLoop(time.Duration(*collectLoopInterval * float64(time.Second)))

	promLog.Infoln("Starting exporter at", *listenAddr)
	promLog.Infoln("")
	http.ListenAndServe(*listenAddr, nil)
}
