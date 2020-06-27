package main

import (
	"encoding/json"
	"flag"
	"local/apcupsd_exporter/metric"
	"local/apcupsd_exporter/model"
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
	defaultState        *model.State
}

func parseArgs() cliArgs {
	listen := flag.String("listen", "0.0.0.0:8001", "ip:port")
	apcupsd := flag.String("apcupsd", "127.0.0.1:3551", "apcupsd host:port")
	apcaccess := flag.String("apcaccess", "/sbin/apcaccess", "apcaccess path")
	floodlimit := flag.Float64("floodlimit", 0.5, "Min time delta between apcaccess calls in seconds")
	collectinterval := flag.Float64("collectinterval", 10, "Base Collect loop interval in seconds")
	defStateJSON := flag.String("default_model_state", "",
		"JSON of default values of model state.\n"+
			"For example: '{\"OutputPowerNominal\": 100500}' returns metric "+
			"'apcupsd_output_power_nominal 100500' if no value was parsed in apcaccess output",
	)
	flag.Parse()

	args := cliArgs{
		listenAddr:          *listen,
		apcupsdAddr:         *apcupsd,
		apcaccessPath:       *apcaccess,
		apcaccessFloodLimit: time.Duration(*floodlimit * float64(time.Second)),
		collectInterval:     time.Duration(*collectinterval * float64(time.Second)),
	}

	if *defStateJSON != "" {
		args.defaultState = &model.State{}
		if err := json.Unmarshal([]byte(*defStateJSON), args.defaultState); err != nil {
			promLog.Errorln("Error on parsing 'default_model_state':", err)
			args.defaultState = nil
		}
	}

	promLog.Infof("Parsed cli args:\n %#v\n\n", args)

	return args
}

func main() {
	args := parseArgs()

	collector := metric.NewCollector(metric.CollectorOtps{
		ApcupsdAddr:         args.apcupsdAddr,
		ApcaccessPath:       args.apcaccessPath,
		ApcaccessFloodLimit: args.apcaccessFloodLimit,
		CollectInterval:     args.collectInterval,
		DefaultState:        args.defaultState,
	})
	collector.Start()

	server.RegisterMetricEndpoints(collector)
	server.RegisterWsEndpoints(collector)

	promLog.Infof("Starting exporter at %s\n\n", args.listenAddr)
	if err := http.ListenAndServe(args.listenAddr, nil); err != nil {
		promLog.Fatalln("Cant start server: ", err.Error())
	}
}
