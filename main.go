package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	promLog "github.com/prometheus/common/log"
)

// Common
var (
	logger = log.New(os.Stdout, "", 0)
)

// Regex
var (
	batteryChargeRegex = regexp.MustCompile(`(?:battery[.]charge:(?:\s)(.*))`)
	upsStatusRegex     = regexp.MustCompile(`(?:ups[.]status:(?:\s)(.*))`)
)

// NUT Gauges
var (
	batteryCharge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "battery_charge",
		Help: "Current battery charge (percent)",
	})

	upsStatus = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "ups_status",
		Help: "Current UPS Status (0=Calibration, 1=SmartTrim, 2=SmartBoost, 3=Online, 4=OnBattery, 5=Overloaded, 6=LowBattery, 7=ReplaceBattery, 8=OnBypass, 9=Off, 10=Charging, 11=Discharging)",
	})
)

func registerMetrics() {
	prometheus.MustRegister(batteryCharge)
	prometheus.MustRegister(upsStatus)
}

func pollMetricsLoop(upscBinary string, addr string, interval time.Duration) {
	for {
		pollMetrics(upscBinary, addr)
		time.Sleep(interval)
	}
}

func pollMetrics(upscBinary string, addr string) {
	promLog.Infoln("pollMetrics()")

	upsOutput, err := exec.Command(upscBinary, "status", addr, "-u").Output()
	kvMap := parseOutput(string(upsOutput))

	promLog.Infof("kvMap %#v", kvMap)

	if err != nil {
		promLog.Fatal(err)
	}

	if batteryChargeRegex.FindAllStringSubmatch(string(upsOutput), -1) == nil {
		prometheus.Unregister(batteryCharge)
	} else {
		batteryChargeValue, _ := strconv.ParseFloat(batteryChargeRegex.FindAllStringSubmatch(string(upsOutput), -1)[0][1], 64)
		batteryCharge.Set(batteryChargeValue)
	}
}

func parseOutput(str string) map[string]string {
	dict := map[string]string{}
	for _, line := range strings.Split(str, "\n") {
		slice := strings.SplitN(line, ":", 2)
		if len(slice) == 2 {
			k := strings.Trim(slice[0], " \t")
			v := strings.Trim(slice[1], " \t")
			dict[k] = v
		}
	}
	return dict
}

func main() {
	listenArg := flag.String("listen", "0.0.0.0:8001", "ip:port")
	addrArg := flag.String("apcaddr", "127.0.0.1:3551", "apcupsd host:port")
	upscArg := flag.String("apcaccess", "/sbin/apcaccess", "apcaccess path")
	flag.Parse()

	registerMetrics()
	go pollMetricsLoop(*upscArg, *addrArg, 20*time.Second)

	promLog.Infoln("Starting exporter at", *listenArg)
	promLog.Infoln("Watching ups", *addrArg)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(*listenArg, nil)
}
