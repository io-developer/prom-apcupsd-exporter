package main

import (
	"flag"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
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

func startMetrics(upscBinary string, addr string) {
	for {
		upsOutput, err := exec.Command(upscBinary, "status", addr, "-u").Output()

		log.Infoln("upsOutput", upsOutput)

		if err != nil {
			log.Fatal(err)
		}

		if batteryChargeRegex.FindAllStringSubmatch(string(upsOutput), -1) == nil {
			prometheus.Unregister(batteryCharge)
		} else {
			batteryChargeValue, _ := strconv.ParseFloat(batteryChargeRegex.FindAllStringSubmatch(string(upsOutput), -1)[0][1], 64)
			batteryCharge.Set(batteryChargeValue)
		}

		time.Sleep(10 * time.Second)
	}
}

func main() {
	listenArg := flag.String("listen", "0.0.0.0:8001", "ip:port")
	addrArg := flag.String("apcaddr", "127.0.0.1:3551", "apcupsd host:port")
	upscArg := flag.String("apcaccess", "/sbin/apcaccess", "apcaccess path")
	flag.Parse()

	registerMetrics()

	log.Infoln("Starting NUT exporter on ups", *addrArg)
	http.Handle("/metrics", promhttp.Handler())

	log.Infoln("Exporter started at ", *listenArg)
	http.ListenAndServe(*listenArg, nil)

	go startMetrics(*upscArg, *addrArg)
}
