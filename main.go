package main

import (
	"flag"
	"local/apcupsd_exporter/metric"
	"local/apcupsd_exporter/server"
	"net/http"
	"time"

	promLog "github.com/prometheus/common/log"
)

type cliArgs struct {
	listenAddr          string
	apcupsdAddr         string
	apcaccessPath       string
	apcaccessFloodLimit time.Duration
	collectInterval     time.Duration
}

func parseArgs() cliArgs {
	listen := flag.String("listen", "0.0.0.0:8001", "ip:port")
	apcupsd := flag.String("apcupsd", "127.0.0.1:3551", "apcupsd host:port")
	apcaccess := flag.String("apcaccess", "/sbin/apcaccess", "apcaccess path")
	floodlimit := flag.Float64("floodlimit", 0.5, "Min time delta between apcaccess calls in seconds")
	collectinterval := flag.Float64("collectinterval", 10, "Base Collect loop interval in seconds")
	flag.Parse()

	args := cliArgs{
		listenAddr:          *listen,
		apcupsdAddr:         *apcupsd,
		apcaccessPath:       *apcaccess,
		apcaccessFloodLimit: time.Duration(*floodlimit * float64(time.Second)),
		collectInterval:     time.Duration(*collectinterval * float64(time.Second)),
	}

	promLog.Infof("Parsed cli args:\n %#v\n\n", args)

	return args
}

func main() {
	args := parseArgs()

	metric.ApcupsdAddr = args.apcupsdAddr
	metric.ApcaccessPath = args.apcaccessPath
	metric.ApcaccessFloodLimit = args.apcaccessFloodLimit

	metric.RegisterPermanents()
	go metric.Collect(metric.CollectChan)
	go metric.CollectLoop(args.collectInterval)

	http.HandleFunc("/metrics", server.HandleMetrics)
	http.HandleFunc("/ws", server.HandleWs)

	promLog.Infof("Starting exporter at %s\n\n", args.listenAddr)
	if err := http.ListenAndServe(args.listenAddr, nil); err != nil {
		promLog.Fatalln("Cant start server: ", err.Error())
	}
}
