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

	promLog.Infof("\n Model.PrevState:\n %#v \n", Model.PrevState)
	promLog.Infof("\n Model.State:\n %#v \n", Model.State)
	promLog.Infof("\n Model.ChangedFields:\n %#v \n\n", Model.ChangedFields)

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
