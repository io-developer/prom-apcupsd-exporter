package dto

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metric struct {
	InputSensivity                *prometheus.GaugeVec
	InputFrequency                prometheus.Gauge
	InputVoltage                  prometheus.Gauge
	InputVoltageMin               prometheus.Gauge
	InputVoltageMax               prometheus.Gauge
	InputVoltageNominal           prometheus.Gauge
	InputVoltageTransferLow       prometheus.Gauge
	InputVoltageTransferHigh      prometheus.Gauge
	OutputLoad                    prometheus.Gauge
	OutputAmps                    prometheus.Gauge
	OutputPowerNominal            prometheus.Gauge
	OutputPowerApparentNominal    prometheus.Gauge
	OutputVoltage                 prometheus.Gauge
	OutputVoltageNominal          prometheus.Gauge
	BatteryCharge                 prometheus.Gauge
	BatteryVoltage                prometheus.Gauge
	BatteryVoltageNominal         prometheus.Gauge
	BatteryExternalCount          prometheus.Gauge
	BatteryBadCount               prometheus.Gauge
	BatteryReplacedDate           prometheus.Gauge
	UpsManafacturedDate           prometheus.Gauge
	UpsDipSwitchFlag              prometheus.Gauge
	UpsReg1                       prometheus.Gauge
	UpsReg2                       prometheus.Gauge
	UpsReg3                       prometheus.Gauge
	UpsTimeleftSeconds            prometheus.Gauge
	UpsTimeleftSecondsLowBattery  prometheus.Gauge
	UpsTransferOnBatteryCount     *prometheus.GaugeVec
	UpsTransferOnBatteryReason    *prometheus.GaugeVec
	UpsOnBatterySeconds           prometheus.Gauge
	UpsOnBatterySecondsCumulative prometheus.Gauge
	UpsTransferOnBatteryDate      prometheus.Gauge
	UpsTransferOffBatteryDate     prometheus.Gauge
	UpsTurnOnDelaySeconds         prometheus.Gauge
	UpsTurnOnBatteryMin           prometheus.Gauge
	UpsTurnOffDelaySeconds        prometheus.Gauge
	UpsTempInternal               prometheus.Gauge
	UpsTempAmbient                prometheus.Gauge
	UpsHumidity                   prometheus.Gauge
	UpsAlarmMode                  *prometheus.GaugeVec
	UpsSelftestResult             *prometheus.GaugeVec
	UpsSelftestIntervalSeconds    prometheus.Gauge
	UpsCable                      *prometheus.GaugeVec
	UpsDriver                     *prometheus.GaugeVec
	UpsMode                       *prometheus.GaugeVec
	UpsStatus                     *prometheus.GaugeVec
	UpsStatusFlags                *prometheus.GaugeVec
	UpsStatusFlagCounters         *prometheus.GaugeVec
}

func NewMetric(constLabels prometheus.Labels) *Metric {
	return &Metric{
		InputSensivity: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name:        "apcupsd_input_sensitivity",
				Help:        "**SENSE** The sensitivity level of the UPS to line voltage fluctuations. 1='Low', 2='Medium', 3='High', 4='Auto Adjust', 5='Unknown'",
				ConstLabels: constLabels,
			}, []string{
				"name",
				"title",
				"value",
			},
		),
		InputFrequency: prometheus.NewGauge(prometheus.GaugeOpts{
			Name:        "apcupsd_input_frequency",
			Help:        "**LINEFREQ** Line frequency in hertz as given by the UPS.",
			ConstLabels: constLabels,
		}),
	}
}
