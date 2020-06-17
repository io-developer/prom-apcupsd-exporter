package metric

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

	promLog.Infof("output %+#v", output)

	val, exists := output.GetFloat64("foo")
	promLog.Infoln("get foo", val, "exists", exists)

	val, exists = output.GetFloat64("LINEV")
	promLog.Infoln("get LINEV", val, "exists", exists)

	//	BatteryCharge.Set(output.Get())
}
