package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"local/apcupsd_exporter/metric"
	"local/apcupsd_exporter/model"
	"local/apcupsd_exporter/server"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/common/promlog"
)

var logger log.Logger

func createLogger(logLevel string) log.Logger {
	l := &promlog.AllowedLevel{}
	l.Set(logLevel)
	return promlog.New(&promlog.Config{Level: l})
}

type cliArgs struct {
	logLevel            string
	listenAddr          string
	apcupsdAddr         string
	apcaccessPath       string
	apcaccessFloodLimit time.Duration
	collectInterval     time.Duration
	defaultState        *model.State
}

func parseArgs() cliArgs {
	loglevel := flag.String("loglevel", "info", "Options: debug, info, warn, error")
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
		logLevel:            *loglevel,
		listenAddr:          *listen,
		apcupsdAddr:         *apcupsd,
		apcaccessPath:       *apcaccess,
		apcaccessFloodLimit: time.Duration(*floodlimit * float64(time.Second)),
		collectInterval:     time.Duration(*collectinterval * float64(time.Second)),
	}

	if *defStateJSON != "" {
		args.defaultState = &model.State{}
		if err := json.Unmarshal([]byte(*defStateJSON), args.defaultState); err != nil {
			args.defaultState = nil
		}
	}

	return args
}

func main() {
	args := parseArgs()

	logger = createLogger(args.logLevel)
	metric.Logger = logger
	server.Logger = logger

	level.Debug(logger).Log("msg", fmt.Sprintf("Parsed cli args:\n %#v\n\n", args))

	collector := metric.NewCollector(metric.CollectorOtps{
		ApcupsdAddr:         args.apcupsdAddr,
		ApcaccessPath:       args.apcaccessPath,
		ApcaccessFloodLimit: args.apcaccessFloodLimit,
		CollectInterval:     args.collectInterval,
		DefaultState:        args.defaultState,
	})
	collector.Start()

	server.MetricsRegister(collector)
	server.WsRegister(collector)

	logger.Log("msg", fmt.Sprintf("Starting exporter at %s\n\n", args.listenAddr))

	if err := http.ListenAndServe(args.listenAddr, nil); err != nil {
		level.Error(logger).Log("msg", "Can't start server", "err", err)
	}
}
