package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/io-developer/prom-apcupsd-exporter/pkg/metric"
	"github.com/io-developer/prom-apcupsd-exporter/pkg/model"
	"github.com/io-developer/prom-apcupsd-exporter/pkg/server"

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
	logLevel                 string
	listenAddr               string
	apcupsdAddr              string
	apcaccessCmd             string
	apcaccessPath            string
	apcaccessFloodLimit      time.Duration
	apcaccessErrorIgnoreTime time.Duration
	apcupsdStartSkip         time.Duration
	collectInterval          time.Duration
	defaultState             *model.State
}

func parseArgs() cliArgs {
	logLevel := flag.String("logLevel", "info", "Options: debug, info, warn, error")
	listen := flag.String("listen", "0.0.0.0:8001", "ip:port")
	apcupsd := flag.String("apcupsd", "127.0.0.1:3551", "apcupsd host:port")
	apcaccessCmd := flag.String("apcaccessCmd", "/sbin/apcaccess -h 127.0.0.1:3551", "apcaccess cmd")
	apcaccess := flag.String("apcaccess", "/sbin/apcaccess", "apcaccess path")
	floodLimit := flag.Float64("floodLimit", 0.5, "Min time delta between apcaccess calls in seconds")
	errorIgnoreTime := flag.Float64("errorIgnoreTime", 120, "Max time in seconds to ignore apcaccess read errors")
	apcupsdStartSkip := flag.Float64("apcupsdStartSkip", 60, "Ignore first N sec agetr apcupsd start")
	collectInterval := flag.Float64("collectInterval", 30, "Base Collect loop interval in seconds")
	defStateJSON := flag.String("defaultModelState", "",
		"JSON of default values of model state.\n"+
			"For example: '{\"OutputPowerNominal\": 100500}' returns metric "+
			"'apcupsd_output_power_nominal 100500' if no value was parsed in apcaccess output",
	)
	flag.Parse()

	args := cliArgs{
		logLevel:                 *logLevel,
		listenAddr:               *listen,
		apcupsdAddr:              *apcupsd,
		apcaccessCmd:             *apcaccessCmd,
		apcaccessPath:            *apcaccess,
		apcaccessFloodLimit:      time.Duration(*floodLimit * float64(time.Second)),
		apcaccessErrorIgnoreTime: time.Duration(*errorIgnoreTime * float64(time.Second)),
		apcupsdStartSkip:         time.Duration(*apcupsdStartSkip * float64(time.Second)),
		collectInterval:          time.Duration(*collectInterval * float64(time.Second)),
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

	level.Debug(logger).Log("msg", fmt.Sprintf("Parsed cli args:\n %#v\n\n", args))

	collector := metric.NewCollector(metric.CollectorOtps{
		ApcupsdAddr:              args.apcupsdAddr,
		ApcaccessCmd:             args.apcaccessCmd,
		ApcaccessPath:            args.apcaccessPath,
		ApcaccessFloodLimit:      args.apcaccessFloodLimit,
		ApcaccessErrorIgnoreTime: args.apcaccessErrorIgnoreTime,
		ApcupsdStartSkip:         args.apcupsdStartSkip,
		CollectInterval:          args.collectInterval,
		DefaultState:             args.defaultState,
	})
	collector.Start()

	server.Init(logger, collector)

	logger.Log("msg", fmt.Sprintf("Starting exporter at %s\n\n", args.listenAddr))

	if err := http.ListenAndServe(args.listenAddr, nil); err != nil {
		level.Error(logger).Log("msg", "Can't start server", "err", err)
	}
}
