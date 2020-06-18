package collect

import (
	"local/apcupsd_exporter/apc"
	"os/exec"
	"regexp"
	"time"

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

func CollectMetrics(apcaccessPath string, addr string) {
	promLog.Infoln("CollectMetrics()")

	cmdResult, err := exec.Command(apcaccessPath, "status", addr, "-u").Output()
	if err != nil {
		promLog.Fatal(err)
	}

	output := apc.NewOutput(string(cmdResult))
	output.Parse()

	promLog.Infof("output.Parsed %+#v", output.Parsed)

	for _, metric := range Metrics {
		val, exists := output.GetFloat64(metric.OutputKey)
		if exists {
			metric.Gauge.Set(val)
			metric.Register()
		} else if metric.IsPermanent {
			metric.Gauge.Set(metric.DefaultValue)
		} else {
			metric.Unregister()
		}
	}
}
