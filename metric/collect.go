package metric

import (
	"local/apcupsd_exporter/apc"
	"os/exec"
	"time"

	promLog "github.com/prometheus/common/log"
)

// CollectLoop ..
func CollectLoop(upscBinary string, addr string, interval time.Duration) {
	for {
		Collect(upscBinary, addr)
		time.Sleep(interval)
	}
}

// Collect ..
func Collect(apcaccessPath string, addr string) {
	promLog.Infoln("CollectMetrics()")

	cmdResult, err := exec.Command(apcaccessPath, "status", addr, "-u").Output()
	if err != nil {
		promLog.Fatal(err)
	}

	output := apc.NewOutput(string(cmdResult))
	output.Parse()

	promLog.Infof("output.Parsed %+#v", output.Parsed)

	for _, metric := range Metrics {
		metric.Handler.Handle(metric, output)
	}
}
