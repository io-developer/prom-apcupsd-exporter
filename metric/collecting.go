package metric

import (
	"local/apcupsd_exporter/apc"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	promLog "github.com/prometheus/common/log"
)

// Regex
var (
	batteryChargeRegex = regexp.MustCompile(`(?:battery[.]charge:(?:\s)(.*))`)
	upsStatusRegex     = regexp.MustCompile(`(?:ups[.]status:(?:\s)(.*))`)
)

func CollectMetricsLoop(upscBinary string, addr string, interval time.Duration) {
	for {
		CollectMetrics(upscBinary, addr)
		time.Sleep(interval)
	}
}

func CollectMetrics(upscBinary string, addr string) {
	promLog.Infoln("CollectMetrics()")

	upsOutput, err := exec.Command(upscBinary, "status", addr, "-u").Output()
	kvMap := apc.ParseOutput(string(upsOutput))

	promLog.Infof("kvMap %#v", kvMap)

	if err != nil {
		promLog.Fatal(err)
	}

	if batteryChargeRegex.FindAllStringSubmatch(string(upsOutput), -1) == nil {
		prometheus.Unregister(BatteryCharge)
	} else {
		batteryChargeValue, _ := strconv.ParseFloat(batteryChargeRegex.FindAllStringSubmatch(string(upsOutput), -1)[0][1], 64)
		BatteryCharge.Set(batteryChargeValue)
	}
}
