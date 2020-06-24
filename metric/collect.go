package metric

import (
	"local/apcupsd_exporter/apc"
	"local/apcupsd_exporter/model"
	"os/exec"
	"time"

	promLog "github.com/prometheus/common/log"
)

// Vars ..
var (
	Model              *model.Model
	ApcaccessPath      string
	ApcupsdAddr        string
	CollectMinInterval time.Duration
)

var lastCollectNanoTs = int64(0)

func checkInterval() bool {
	nowNanoTs := time.Now().UnixNano()
	res := nowNanoTs-lastCollectNanoTs >= int64(CollectMinInterval)
	if res {
		lastCollectNanoTs = nowNanoTs
	}
	return res
}

// Collect ..
func Collect() {
	if !checkInterval() {
		return
	}

	promLog.Infoln("Collect..")

	cmdResult, err := exec.Command(ApcaccessPath, "status", ApcupsdAddr).Output()
	if err != nil {
		promLog.Fatal(err)
	}

	output := apc.NewOutput(string(cmdResult))
	output.Parse()

	promLog.Infoln("Output parsed")

	Model.Update(model.NewStateFromOutput(output))

	for field, diff := range Model.ChangedFields {
		promLog.Infof("Changed '%s'\n  OLD: %#v\n  NEW: %#v\n", field, diff[0], diff[1])
	}

	for _, metric := range Metrics {
		metric.Handler.Handle(metric, output)
	}

	promLog.Infoln("Metrics handled")
}

// CollectLoop ..
func CollectLoop(interval time.Duration) {
	for {
		Collect()
		time.Sleep(interval)
	}
}
