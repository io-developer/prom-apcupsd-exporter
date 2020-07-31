package metric

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/io-developer/prom-apcupsd-exporter/apcupsd"
	"github.com/io-developer/prom-apcupsd-exporter/model"

	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

var defaultFactory = NewFactory()

// CollectorOtps ..
type CollectorOtps struct {
	ApcupsdAddr         string
	ApcaccessPath       string
	ApcaccessFloodLimit time.Duration
	CollectInterval     time.Duration
	ApcupsdStartSkip    time.Duration
	Factory             *Factory
	DefaultState        *model.State
}

// Collector ..
type Collector struct {
	opts         *CollectorOtps
	started      bool
	collectCh    chan CollectOpts
	currModel    *model.Model
	lastOutput   *apcupsd.Output
	lastOutputTs int64
	lastState    model.State
	metrics      []*Metric
}

// NewCollector ..
func NewCollector(opts CollectorOtps) *Collector {
	if opts.Factory == nil {
		opts.Factory = defaultFactory
	}
	return &Collector{
		opts:       &opts,
		collectCh:  make(chan CollectOpts),
		currModel:  model.NewModel(),
		lastOutput: apcupsd.NewOutput(""),
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

// GetModel method
func (c *Collector) GetModel() *model.Model {
	return c.currModel
}

// GetLastOutput method
func (c *Collector) GetLastOutput() *apcupsd.Output {
	return c.lastOutput
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

	c.updateOutput(opts)
	c.updateModel(opts)
	c.updateMetrics(opts)

	if opts.OnComplete != nil {
		opts.OnComplete <- true
	}
	level.Debug(Logger).Log("msg", "collect end <--")
}

func (c *Collector) updateOutput(opts CollectOpts) {
	if opts.SkipApcupsdParsing {
		level.Debug(Logger).Log("msg", "collect: skipping apcupsd output parsing")
		return
	}
	level.Debug(Logger).Log("msg", "collect: pending apcupsd output")

	ts := time.Now().UnixNano()
	if opts.PreventFlood && ts-c.lastOutputTs < int64(c.opts.ApcaccessFloodLimit) {
		level.Debug(Logger).Log("msg", "collect: apcupsd flood detected, skipping")
		return
	}
	c.lastOutputTs = ts

	cmdResult, err := exec.Command(c.opts.ApcaccessPath, "status", c.opts.ApcupsdAddr).Output()
	if err != nil {
		level.Error(Logger).Log(
			"msg", "apcaccess cmd exited with error",
			"err", err,
			"result", string(cmdResult),
		)
		cmdResult = []byte{}
	}

	c.lastOutput = apcupsd.NewOutput(string(cmdResult))
	c.lastOutput.Parse()

	c.lastState = c.parseState()
}

func (c *Collector) updateModel(opts CollectOpts) {
	level.Debug(Logger).Log("msg", "collect: updating model")

	dt := time.Now().Sub(c.lastState.ApcupsdStartTime)
	if dt < c.opts.ApcupsdStartSkip {
		level.Warn(Logger).Log(
			"msg", "skipping state update due to start delay",
			"dt", dt,
		)
		return
	}

	c.currModel.Update(c.lastState)

	for field, diff := range c.currModel.ChangedFields {
		level.Info(Logger).Log(
			"change", field,
			"old", fmt.Sprintf("%#v", diff[0]),
			"new", fmt.Sprintf("%#v", diff[1]),
		)
	}
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
	PreventFlood       bool
	SkipApcupsdParsing bool
	Signal             model.Signal
	OnComplete         chan bool
}
