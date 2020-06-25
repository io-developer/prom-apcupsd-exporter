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
	ApcupsdAddr         = "127.0.0.1:3551"
	ApcaccessPath       = "/sbin/apcaccess"
	ApcaccessFloodLimit = time.Duration(0)
	CurrentOutput       *apc.Output
	CurrentModel        = model.NewModel()
	CollectChan         = make(chan CollectOpts)
)

// CollectOpts ..
type CollectOpts struct {
	PreventFlood bool
	OnComplete   chan bool
}

// Collect ..
func Collect(c chan CollectOpts) {
	for true {
		opts, ok := <-c
		if !ok {
			return
		}

		if !opts.PreventFlood || checkFloodInterval() {
			parseOutput()
			updateModel()
			logModelChanges()
			updateMetrics()
		}

		if opts.OnComplete != nil {
			opts.OnComplete <- true
		}
	}
}

// CollectLoop ..
func CollectLoop(interval time.Duration) {
	for {
		CollectChan <- CollectOpts{PreventFlood: true}
		time.Sleep(interval)
	}
}

var lastCollectNanoTs = int64(0)

func checkFloodInterval() bool {
	nowNanoTs := time.Now().UnixNano()
	res := nowNanoTs-lastCollectNanoTs >= int64(ApcaccessFloodLimit)
	if res {
		lastCollectNanoTs = nowNanoTs
	}
	return res
}

func parseOutput() {
	promLog.Infoln("parsing apcupsd output..")

	cmdResult, err := exec.Command(ApcaccessPath, "status", ApcupsdAddr).Output()
	if err != nil {
		promLog.Errorln("apcaccess exited with error")
		promLog.Errorln("  Error:", err.Error())
		promLog.Errorln("  Result:", string(cmdResult))
		cmdResult = []byte{}
	}
	CurrentOutput = apc.NewOutput(string(cmdResult))
	CurrentOutput.Parse()
}

func updateModel() {
	promLog.Infoln("updating model..")
	CurrentModel.Update(model.NewStateFromOutput(CurrentOutput))
}

func logModelChanges() {
	for field, diff := range CurrentModel.ChangedFields {
		promLog.Infof("field changed '%s'\n  OLD: %#v\n  NEW: %#v\n", field, diff[0], diff[1])
	}
}

func updateMetrics() {
	promLog.Infoln("updating metrics..")
	for _, metric := range Metrics {
		if metric.HandlerFunc != nil {
			metric.HandlerFunc(metric, CurrentModel)
		} else if metric.ValFunc != nil {
			metric.UpdateCollector(metric.ValFunc(metric, CurrentModel))
		}
	}
	promLog.Infoln("metrics updated")
}
