package metric

import (
	"fmt"
	"time"

	"github.com/io-developer/prom-apcupsd-exporter/pkg/cmd"
	"github.com/io-developer/prom-apcupsd-exporter/pkg/dto"
	"github.com/io-developer/prom-apcupsd-exporter/pkg/model"
	"github.com/io-developer/prom-apcupsd-exporter/pkg/parsing"

	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

var defaultFactory = NewFactory()

// CollectorOtps ..
type CollectorOtps struct {
	ApcupsdAddr              string
	ApcaccessCmd             string
	ApcaccessPath            string
	ApcaccessFloodLimit      time.Duration
	ApcaccessErrorIgnoreTime time.Duration
	CollectInterval          time.Duration
	ApcupsdStartSkip         time.Duration
	Factory                  *Factory
	DefaultState             *model.State
}

// Collector ..
type Collector struct {
	opts                *CollectorOtps
	started             bool
	collectCh           chan CollectOpts
	currModel           *model.Model
	lastOutput          *dto.ApcupsdResponse
	lastOutputTs        int64
	lastSuccessOutputTs int64
	lastState           model.State
	metrics             []*Metric
}

// NewCollector ..
func NewCollector(opts CollectorOtps) *Collector {
	if opts.Factory == nil {
		opts.Factory = defaultFactory
	}
	return &Collector{
		opts:      &opts,
		collectCh: make(chan CollectOpts),
		currModel: model.NewModel(),
		lastOutput: &dto.ApcupsdResponse{
			Output:    "",
			KeyValues: make(map[string]string),
		},
	}
}

// Start method
func (c *Collector) Start() {
	if !c.started {
		c.started = true

		go c.listenCollect()
		go c.loopCollect()
	}
}

// GetFactory method
func (c *Collector) GetFactory() *Factory {
	return c.opts.Factory
}

// Collect method
func (c *Collector) Collect(opts CollectOpts) {
	level.Debug(Logger).Log("msg", "collect requested")

	c.collectCh <- opts
}

func (c *Collector) loopCollect() {
	for {
		c.Collect(CollectOpts{
			PreventFlood: true,
		})
		time.Sleep(c.opts.CollectInterval)
	}
}

func (c *Collector) listenCollect() {
	for {
		if opts, ok := <-c.collectCh; ok {
			c.collect(opts)
		} else {
			return
		}
	}
}

func (c *Collector) collect(opts CollectOpts) {
	level.Debug(Logger).Log("msg", "collect begin -->")

	if !c.updateOutput(opts) {
		level.Debug(Logger).Log("msg", "collect updateOutput failed <--")
		if opts.OnComplete != nil {
			opts.OnComplete <- true
		}
		return
	}
	if !c.updateModel(opts) {
		level.Debug(Logger).Log("msg", "collect updateModel failed <--")
		if opts.OnComplete != nil {
			opts.OnComplete <- true
		}
		return
	}

	c.updateMetrics(opts)

	if opts.OnComplete != nil {
		opts.OnComplete <- true
	}
	level.Debug(Logger).Log("msg", "collect end <--")
}

func (c *Collector) updateOutput(opts CollectOpts) bool {
	level.Debug(Logger).Log("msg", "collect: pending apcupsd output")

	ts := time.Now().UnixNano()
	if opts.PreventFlood && ts-c.lastOutputTs < int64(c.opts.ApcaccessFloodLimit) {
		level.Debug(Logger).Log("msg", "collect: apcupsd flood detected, skipping")
		return false
	}
	c.lastOutputTs = ts

	shell := cmd.NewShell()
	stdout, stderr, exitCode, err := shell.Exec(c.opts.ApcaccessCmd)
	if err != nil || exitCode != 0 {
		//log.Printf("[ERROR] apcaccess: %#v\n%#v\n", err, stderr)
		level.Error(Logger).Log(
			"msg", "apcaccess cmd exited with error",
			"cmd", c.opts.ApcaccessCmd,
			"err", err,
			"stderr", string(stderr),
		)
		if ts-c.lastSuccessOutputTs <= int64(c.opts.ApcaccessErrorIgnoreTime) {
			level.Warn(Logger).Log(
				"msg", "Ignoring bad exit code for a while...",
				"err", err,
			)
		}
		return false
	}

	c.lastSuccessOutputTs = ts

	resp, err := parsing.NewApcupsdParser().ParseApcaccessOutput(string(stdout))
	if err != nil {
		level.Error(Logger).Log(
			"msg", "apcaccess response parsing error",
			"err", err,
			"result", string(stdout),
		)
	}

	c.lastOutput = resp
	c.lastState = c.parseState()

	return true
}

func (c *Collector) updateModel(opts CollectOpts) bool {
	level.Debug(Logger).Log("msg", "collect: updating model")

	dt := time.Now().Sub(c.lastState.ApcupsdStartTime)
	if dt < c.opts.ApcupsdStartSkip {
		level.Warn(Logger).Log(
			"msg", "skipping state update due to start delay",
			"dt", dt,
		)
		return false
	}

	c.currModel.Update(c.lastState)

	for field, diff := range c.currModel.ChangedFields {
		level.Info(Logger).Log(
			"change", field,
			"old", fmt.Sprintf("%#v", diff[0]),
			"new", fmt.Sprintf("%#v", diff[1]),
		)
	}

	return true
}

func (c *Collector) parseState() model.State {
	prev := c.currModel.PrevState
	state := model.NewStateFromOutput(c.lastOutput, c.opts.DefaultState)

	if state.UpsTransferOffBatteryDate.Sub(prev.UpsTransferOffBatteryDate) < 0 {
		state.UpsTransferOffBatteryDate = prev.UpsTransferOffBatteryDate
	}
	if state.UpsTransferOnBatteryDate.Sub(prev.UpsTransferOnBatteryDate) < 0 {
		state.UpsTransferOnBatteryDate = prev.UpsTransferOnBatteryDate
	}

	return state
}

func (c *Collector) updateMetrics(opts CollectOpts) {
	level.Debug(Logger).Log("msg", "collect: updating metrics")

	state := c.currModel.State
	c.GetFactory().SetConstLabels(prometheus.Labels{
		"ups_serial": state.UpsSerial,
		"ups_model":  state.UpsModel,
		"ups_name":   state.UpsName,
	})

	metrics, metricsChanged := c.opts.Factory.GetMetrics()
	if metricsChanged {
		level.Debug(Logger).Log("msg", "collect: metrics changed, rebuilding..")

		for _, metric := range c.metrics {
			metric.Unregister()
		}
	}

	c.metrics = metrics
	for _, metric := range c.metrics {
		metric.Update(c.currModel)
	}
}

// CollectOpts ..
type CollectOpts struct {
	PreventFlood bool
	OnComplete   chan bool
}
