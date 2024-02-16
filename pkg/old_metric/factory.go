package old_metric

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/io-developer/prom-apcupsd-exporter/pkg/old_model"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

// Factory of default metrics
type Factory struct {
	constLabels prometheus.Labels
	currMetrics []*Metric
}

// NewFactory constructor
func NewFactory() *Factory {
	return &Factory{
		constLabels: prometheus.Labels{},
		currMetrics: nil,
	}
}

// SetConstLabels method
func (f *Factory) SetConstLabels(labels prometheus.Labels) {
	if labels == nil {
		return
	}
	if len(labels) != len(f.constLabels) || !reflect.DeepEqual(labels, f.constLabels) {
		f.currMetrics = nil
	}
	f.constLabels = labels
}

// GetMetrics returns metrics with actual labels
func (f *Factory) GetMetrics() (metrics []*Metric, changed bool) {
	if f.currMetrics == nil {
		f.currMetrics = f.createMetrics()
		changed = true
	}
	return f.currMetrics, changed
}

func (f *Factory) createMetrics() []*Metric {
	return []*Metric{

		// Input
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_input_sensitivity",
				Help: "**SENSE** The sensitivity level of the UPS to line voltage fluctuations. " +
					typesToDesc(old_model.SensivityTypes),

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.InputSensivity.Type)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_input_frequency",
				Help: "**LINEFREQ** Line frequency in hertz as given by the UPS.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.InputFrequency
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_input_voltage",
				Help: "**LINEV** The current line voltage as returned by the UPS.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.InputVoltage
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_input_voltage_min",
				Help: "**MINLINEV** The minimum line voltage since the UPS was started, as returned by the UPS",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.InputVoltageMin
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_input_voltage_max",
				Help: "**MAXLINEV** The maximum line voltage since the UPS was started, as reported by the UPS",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.InputVoltageMax
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_input_voltage_nominal",
				Help: "**NOMINV** The input voltage that the UPS is configured to expect.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.InputVoltageNominal
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_input_voltage_transfer_low",
				Help: "**LOTRANS** The line voltage below which the UPS will switch to batteries.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.InputVoltageTransferLow
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_input_voltage_transfer_high",
				Help: "**HITRANS** The line voltage above which the UPS will switch to batteries.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.InputVoltageTransferHigh
			},
		},

		// Output
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_output_load",
				Help: "**LOADPCT** The percentage of load capacity as estimated by the UPS.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.OutputLoad
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_output_amps",
				Help: "**OUTCURNT** Amps",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.OutputAmps
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_output_power_nominal",
				Help: "**NOMPOWER** The maximum power in Watts that the UPS is designed to supply.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.OutputPowerNominal
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_output_power_apparent_nominal",
				Help: "**NOMAPNT** VA.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.OutputPowerApparentNominal
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_output_voltage",
				Help: "**OUTPUTV** The voltage the UPS is supplying to your equipment",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.OutputVoltage
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_output_voltage_nominal",
				Help: "**NOMOUTV** The output voltage that the UPS will attempt to supply when on battery power.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.OutputVoltageNominal
			},
		},

		// Battery
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_battery_charge",
				Help: "**BCHARGE** The percentage charge on the batteries.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.BatteryCharge
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_battery_voltage",
				Help: "**BATTV** Battery voltage as supplied by the UPS.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.BatteryVoltage
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_battery_voltage_nominal",
				Help: "**NOMBATTV** The nominal battery voltage.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.BatteryVoltageNominal
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_battery_externals",
				Help: "**EXTBATTS** The number of external batteries as defined by the user. A correct number here helps the UPS compute the remaining runtime more accurately.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.BatteryExternalCount)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_battery_bads",
				Help: "**BADBATTS** The number of bad battery packs.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.BatteryBadCount)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_battery_replaced",
				Help: "**BATTDATE** The date that batteries were last replaced.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.BatteryReplacedDate.Unix())
			},
		},

		// Ups
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_manafactured",
				Help: "**MANDATE** The date the UPS was manufactured.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsManafacturedDate.Unix())
			},
		},
		{
			IsPermanent: true,
			Collector: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: "apcupsd_ups_status",
				Help: "Current status vec labeled by flag. Value 0 or single flag. Flags: " +
					typedFlagsToDescFmt(old_model.StatusFlags, "0x%08x='%s'"),

				ConstLabels: f.constLabels,
			}, []string{"flag"}),
			HandlerFunc: func(m *Metric, model *old_model.Model) {
				gaugeVec := m.Collector.(*prometheus.GaugeVec)
				for name, val := range model.State.UpsStatus.GetFlags() {
					gaugeVec.WithLabelValues(name).Set(float64(val))
				}
			},
		},
		{
			IsPermanent: true,
			Collector: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: "apcupsd_ups_status01",
				Help: "Current status vec labeled by flag. Value 0 or 1. See flag names in status description",

				ConstLabels: f.constLabels,
			}, []string{"flag"}),
			HandlerFunc: func(m *Metric, model *old_model.Model) {
				gaugeVec := m.Collector.(*prometheus.GaugeVec)
				for name, val := range model.State.UpsStatus.GetNormedFlags(false) {
					gaugeVec.WithLabelValues(name).Set(float64(val))
				}
			},
		},
		{
			IsPermanent: true,
			Collector: prometheus.NewCounterVec(prometheus.CounterOpts{
				Name: "apcupsd_ups_status_changes",
				Help: "Number of status changes per flag ('flag' label).",

				ConstLabels: f.constLabels,
			}, []string{"flag"}),
			HandlerFunc: func(m *Metric, model *old_model.Model) {
				counterVec := m.Collector.(*prometheus.CounterVec)
				for name, val := range model.State.UpsStatus.FlagChangeCounts {
					counter := counterVec.WithLabelValues(name)
					curData := &dto.Metric{}
					counter.Write(curData)
					delta := int64(val) - int64(*curData.Counter.Value)
					if delta > 0 {
						counter.Add(float64(delta))
					}
				}
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_status_flag",
				Help: "**STATFLAG** Current status flag (summary flags value). See flags in status description",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsStatus.Flag)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_dip_switch_flag",
				Help: "**DIPSW** The current dip switch settings on UPSes that have them.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsDipSwitchFlag)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_reg1",
				Help: "**REG1** The value from the UPS fault register 1.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsReg1)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_reg2",
				Help: "**REG2** The value from the UPS fault register 2.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsReg2)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_reg3",
				Help: "**REG3** The value from the UPS fault register 3.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsReg3)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_timeleft",
				Help: "**TIMELEFT** (seconds) The remaining runtime left on batteries as estimated by the UPS.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsTimeleftSeconds)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_timeleft_low_battery",
				Help: "**DLOWBATT** (seconds) The remaining runtime below which the UPS sends the low battery signal. At this point apcupsd will force an immediate emergency shutdown.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsTimeleftSecondsLowBattery)
			},
		},
		{
			Collector: prometheus.NewCounter(prometheus.CounterOpts{
				Name: "apcupsd_ups_transfer_onbattery",
				Help: "**NUMXFERS** The number of transfers to batteries since apcupsd startup.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsTransferOnBatteryCount)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_transfer_onbattery_reason",
				Help: "**LASTXFER** The reason for the last transfer to batteries." +
					typesToDesc(old_model.TransferOnbatteryReasonTypes),

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsTransferOnBatteryReason.Type)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_transfer_onbattery_time",
				Help: "**TONBATT** Time in seconds currently on batteries",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsOnBatterySeconds)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_transfer_onbattery_time_cumulative",
				Help: "Total (cumulative) time on batteries in seconds since apcupsd startup.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsOnBatterySecondsCumulative)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_transfer_onbattery_timestamp",
				Help: "**XONBATT** Time and date of last transfer to batteries",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsTransferOnBatteryDate.Unix())
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_transfer_offbattery_timestamp",
				Help: "Time and date of last transfer from batteries",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsTransferOffBatteryDate.Unix())
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_turnon_delay",
				Help: "**DWAKE** (seconds) The amount of time the UPS will wait before restoring power to your equipment after a power off condition when the power is restored.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsTurnOnDelaySeconds)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_turnon_battery_min",
				Help: "	**RETPCT** The percentage charge that the batteries must have after a power off condition before the UPS will restore power to your equipment.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.UpsTurnOnBatteryMin
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_turnoff_delay",
				Help: "**DSHUTD** (seconds) The grace delay that the UPS gives after receiving a power down command from apcupsd before it powers off your equipment.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsTurnOffDelaySeconds)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_temp_internal",
				Help: "**ITEMP** (Celsius) Internal UPS temperature as supplied by the UPS.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.UpsTempInternal
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_temp_ambient",
				Help: "**AMBTEMP** The ambient temperature as measured by the UPS.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.UpsTempAmbient
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_humidity",
				Help: "**HUMIDITY** The humidity as measured by the UPS.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.UpsHumidity
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_alarm_mode",
				Help: "**ALARMDEL** The delay period for the UPS alarm." +
					typesToDesc(old_model.AlarmModeTypes),

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsAlarmMode.Type)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_selftest_result",
				Help: "**SELFTEST** The results of the last self test, and may have the following values." +
					typesToDesc(old_model.SelftestResultTypes),

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsSelftestResult.Type)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_selftest_interval",
				Help: "**STESTI** The interval in seconds between automatic self tests.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsSelftestIntervalSeconds)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_cable",
				Help: "**CABLE** The cable as specified in the configuration file ('UPSCABLE')." +
					typesToDescFmt(old_model.CableTypes, "% 3d='%s'"),

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsCable.Type)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_driver",
				Help: "**DRIVER** type." +
					typesToDesc(old_model.DriverTypes),

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsDriver.Type)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_ups_mode",
				Help: "**UPSMODE** The mode in which apcupsd is operating as specified in the configuration file ('UPSMODE'). " + typesToDesc(old_model.ModeTypes),

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.UpsMode.Type)
			},
		},

		// Shutdown
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_shutdown_battery_min",
				Help: "**MBATTCHG** If the battery charge percentage (BCHARGE) drops below this value, apcupsd will  shutdown your system.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return model.State.ShutdownBatteryMin
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_shutdown_timeleft_min",
				Help: "**MINTIMEL** (seconds) apcupsd will shutdown your system if the remaining runtime equals or is below this point.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.ShutdownTimeleftSecondsMin)
			},
		},
		{
			Collector: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "apcupsd_shutdown_onbattery_time_max",
				Help: "**MAXTIME** (seconds) apcupsd will shutdown your system if the time on batteries exceeds this value. A value of zero disables the feature.",

				ConstLabels: f.constLabels,
			}),
			ValFunc: func(m *Metric, model *old_model.Model) float64 {
				return float64(model.State.ShutdownOnBatterySecondsMax)
			},
		},
	}
}

func typesToDesc(types map[string]interface{}) string {
	return typesToDescFmt(types, "%d='%s'")
}

func typesToDescFmt(types map[string]interface{}, f string) string {
	descs := []string{}
	used := map[uint64]bool{}
	for name, v := range types {
		val, _ := strconv.ParseUint(fmt.Sprintf("%v", v), 10, 64)
		if _, exists := used[val]; !exists {
			used[val] = true
			descs = append(descs, fmt.Sprintf(f, val, name))
		}
	}
	sort.Strings(descs)
	return strings.Join(descs, ", ")
}

func typedFlagsToDescFmt(typedFlags map[string]uint64, f string) string {
	descs := []string{}
	for name, flag := range typedFlags {
		descs = append(descs, fmt.Sprintf(f, flag, name))
	}
	sort.Strings(descs)
	return strings.Join(descs, ", ")
}
