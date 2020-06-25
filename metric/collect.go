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
	Channel             chan Opts
	ApcupsdAddr         string
	ApcaccessPath       string
	ApcaccessFloodLimit time.Duration
	CurrentOutput       *apc.Output
	CurrentModel        *model.Model
)

// Opts ..
type Opts struct {
	PreventFlood bool
	OnComplete   chan bool
}

// Collect ..
func Collect(c chan Opts) {
	for true {
		opts, ok := <-c
		if !ok {
			return
		}

		if !opts.PreventFlood || checkFloodInterval() {
			promLog.Infoln("parseApcOutput..")
			parseOutput()

			promLog.Infoln("updateModel..")
			updateModel()
			logModelChanges()

			promLog.Infoln("updateMetrics..")
			updateMetrics()

			promLog.Infoln("Collect done")
		}

		if opts.OnComplete != nil {
			opts.OnComplete <- true
		}
	}
}

// CollectLoop ..
func CollectLoop(interval time.Duration) {
	for {
		Channel <- Opts{PreventFlood: true}
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
	cmdResult, err := exec.Command(ApcaccessPath, "status", ApcupsdAddr).Output()
	if err != nil {
		promLog.Fatal(err)
	}
	CurrentOutput = apc.NewOutput(string(cmdResult))
	CurrentOutput.Parse()
}

func updateModel() {
	CurrentModel.Update(model.NewStateFromOutput(CurrentOutput))
}

func logModelChanges() {
	for field, diff := range CurrentModel.ChangedFields {
		promLog.Infof("Changed '%s'\n  OLD: %#v\n  NEW: %#v\n", field, diff[0], diff[1])
	}
}

func updateMetrics() {
	for _, metric := range Metrics {
		metric.Handler.Handle(metric, CurrentOutput)
	}
}
