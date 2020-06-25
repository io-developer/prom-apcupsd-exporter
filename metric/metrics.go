package metric

import (
	"fmt"
	"local/apcupsd_exporter/model"
	"sort"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

/*
input.sensitivity: high
input.transfer.high: 253
input.transfer.low: 200
input.voltage: 228.9
output.current: 0.00
output.frequency: 50.0
output.voltage: 228.9
output.voltage.nominal: 230.0



DATE     : 2020-06-18 16:56:59 +0300
HOSTNAME : home-nas
VERSION  : 3.14.14 (31 May 2016) debian
UPSNAME  : HomeSrv
STARTTIME: 2020-06-18 02:00:02 +0300
MODEL    : Back-UPS CS 500
STATUS   : ONLINE
SERIALNO : 4B1802P05216
FIRMWARE : 808.q10 .I USB FW:q


CABLE    : USB Cable
DRIVER   : USB UPS Driver
UPSMODE  : Stand Alone
LINEV    : 226.0 Volts
LOADPCT  : 7.0 Percent
BCHARGE  : 100.0 Percent
TIMELEFT : 64.5 Minutes
MBATTCHG : 1 Percent
MINTIMEL : -1 Minutes
MAXTIME  : 0 Seconds
OUTPUTV  : 230.0 Volts
LINEFREQ : 50.0 Hz
NOMPOWER : 300 Watts
NOMINV   : 230 Volts
NOMOUTV  : 230 Volts
TONBATT  : 0 Seconds
CUMONBATT: 0 Seconds
XOFFBATT : N/A  date time
XONBATT     date time
SENSE    : Low|High
LOTRANS  : 180.0 Volts
HITRANS  : 260.0 Volts
BATTV    : 13.7 Volts
NOMBATTV : 12.0 Volts
ITEMP    : 29.2 C
BATTDATE : 2018-11-29
MANDATE  : 2018-01-09
DSHUTD   : 180 Seconds
DWAKE    : 0 Seconds
**DLOWBATT**
    The remaining runtime below which the UPS
    sends the low battery signal. At this point apcupsd will force an
	immediate emergency shutdown.
NUMXFERS : 0
RETPCT   : 0.0 Percent
**HUMIDITY**
	The humidity as measured by the UPS.
**AMBTEMP**
    The ambient temperature as measured by the UPS.
**EXTBATTS**
    The number of external batteries as
    defined by the user. A correct number here helps the UPS compute
    the remaining runtime more accurately.
**BADBATTS**
	The number of bad battery packs.
ALARMDEL : No alarm|30 Seconds
SELFTEST : NO
STATFLAG : 0x05000008
STESTI   : None|14 days
LASTXFER : Unacceptable line voltage changes
| No transfers since turnon
| Automatic or explicit self test

*/

// Metrics declare
var Metrics = []*Metric{

	// Input
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_sensitivity",
			Help: "**SENSE** The sensitivity level of the UPS to line voltage fluctuations. " +
				typesToDesc(model.SensivityTypes),
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.InputSensivity.Type)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_frequency",
			Help: "**LINEFREQ** Line frequency in hertz as given by the UPS.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.InputFrequency
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage",
			Help: "**LINEV** The current line voltage as returned by the UPS.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.InputVoltage
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_min",
			Help: "**MINLINEV** The minimum line voltage since the UPS was started, as returned by the UPS",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.InputVoltageMin
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_max",
			Help: "**MAXLINEV** The maximum line voltage since the UPS was started, as reported by the UPS",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.InputVoltageMax
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_nominal",
			Help: "**NOMINV** The input voltage that the UPS is configured to expect.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.InputVoltageNominal
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_transfer_low",
			Help: "**LOTRANS** The line voltage below which the UPS will switch to batteries.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.InputVoltageTransferLow
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_transfer_high",
			Help: "**HITRANS** The line voltage above which the UPS will switch to batteries.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.InputVoltageTransferHigh
		},
	},

	// Output
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_load",
			Help: "**LOADPCT** The percentage of load capacity as estimated by the UPS.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.OutputLoad
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_amps",
			Help: "**OUTCURNT** Amps",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.OutputAmps
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_power_nominal",
			Help: "**NOMPOWER** The maximum power in Watts that the UPS is designed to supply.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.OutputPowerNominal
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_voltage",
			Help: "**OUTPUTV** The voltage the UPS is supplying to your equipment",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.OutputVoltage
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_voltage_nominal",
			Help: "**NOMOUTV** The output voltage that the UPS will attempt to supply when on battery power.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.OutputVoltageNominal
		},
	},

	// Battery
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_charge",
			Help: "**BCHARGE** The percentage charge on the batteries.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.BatteryCharge
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_voltage",
			Help: "**BATTV** Battery voltage as supplied by the UPS.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.BatteryVoltage
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_voltage_nominal",
			Help: "**NOMBATTV** The nominal battery voltage.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.BatteryVoltageNominal
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_external",
			Help: "**EXTBATTS** The number of external batteries as defined by the user. A correct number here helps the UPS compute the remaining runtime more accurately.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.BatteryExternalCount)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_bad",
			Help: "**BADBATTS** The number of bad battery packs.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.BatteryBadCount)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_replaced_timestamp",
			Help: "**BATTDATE** The date that batteries were last replaced.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.BatteryReplacedDate.Unix())
		},
	},

	// Ups
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_manafactured_timestamp",
			Help: "**MANDATE** The date the UPS was manufactured.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsManafacturedDate.Unix())
		},
	},
	{
		IsPermanent: true,
		Collector: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "apcupsd_ups_status",
			Help: "Current status vec labeled by flag. Value 0 or single flag. Flags: " +
				typedFlagsToDescFmt(model.StatusFlags, "0x%08x='%s'"),
		}, []string{"flag"}),
		HandlerFunc: func(m *Metric, model *model.Model) {
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
		}, []string{"flag"}),
		HandlerFunc: func(m *Metric, model *model.Model) {
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
		}, []string{"flag"}),
		HandlerFunc: func(m *Metric, model *model.Model) {
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
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsStatus.Flag)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_dip_switch_flag",
			Help: "**DIPSW** The current dip switch settings on UPSes that have them.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsDipSwitchFlag)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_reg1",
			Help: "**REG1** The value from the UPS fault register 1.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsReg1)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_reg2",
			Help: "**REG2** The value from the UPS fault register 2.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsReg2)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_reg3",
			Help: "**REG3** The value from the UPS fault register 3.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsReg3)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_timeleft",
			Help: "**TIMELEFT** (seconds) The remaining runtime left on batteries as estimated by the UPS.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsTimeleftSeconds)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_timeleft_low_battery",
			Help: "**DLOWBATT** (seconds) The remaining runtime below which the UPS sends the low battery signal. At this point apcupsd will force an immediate emergency shutdown.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsTimeleftSecondsLowBattery)
		},
	},
	{
		Collector: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "apcupsd_ups_transfer_onbattery",
			Help: "**NUMXFERS** The number of transfers to batteries since apcupsd startup.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsTransferOnBatteryCount)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery_reason",
			Help: "**LASTXFER** The reason for the last transfer to batteries." +
				typesToDesc(model.TransferOnbatteryReasonTypes),
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsTransferOnBatteryReason.Type)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery_time",
			Help: "**TONBATT** Time in seconds currently on batteries",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsOnBatterySeconds)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery_time_cumulative",
			Help: "Total (cumulative) time on batteries in seconds since apcupsd startup.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsOnBatterySecondsCumulative)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery_timestamp",
			Help: "**XONBATT** Time and date of last transfer to batteries",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsTransferOnBatteryDate.Unix())
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_offbattery_timestamp",
			Help: "Time and date of last transfer from batteries",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsTransferOffBatteryDate.Unix())
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_turnon_delay",
			Help: "**DWAKE** (seconds) The amount of time the UPS will wait before restoring power to your equipment after a power off condition when the power is restored.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsTurnOnDelaySeconds)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_turnon_battery_min",
			Help: "	**RETPCT** The percentage charge that the batteries must have after a power off condition before the UPS will restore power to your equipment.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.UpsTurnOnBatteryMin
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_turnoff_delay",
			Help: "**DSHUTD** (seconds) The grace delay that the UPS gives after receiving a power down command from apcupsd before it powers off your equipment.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsTurnOffDelaySeconds)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_temp_internal",
			Help: "**ITEMP** (Celsius) Internal UPS temperature as supplied by the UPS.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.UpsTempInternal
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_temp_ambient",
			Help: "**AMBTEMP** The ambient temperature as measured by the UPS.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.UpsTempAmbient
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_humidity",
			Help: "**HUMIDITY** The humidity as measured by the UPS.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.UpsHumidity
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_alarm_mode",
			Help: "**ALARMDEL** The delay period for the UPS alarm." +
				typesToDesc(model.AlarmModeTypes),
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsAlarmMode.Type)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_selftest_result",
			Help: "**SELFTEST** The results of the last self test, and may have the following values." +
				typesToDesc(model.SelftestResultTypes),
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsSelftestResult.Type)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_selftest_interval",
			Help: "**STESTI** The interval in seconds between automatic self tests.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsSelftestIntervalSeconds)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_cable",
			Help: "**CABLE** The cable as specified in the configuration file ('UPSCABLE')." +
				typesToDescFmt(model.CableTypes, "% 3d='%s'"),
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsCable.Type)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_driver",
			Help: "**DRIVER** type." +
				typesToDesc(model.DriverTypes),
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsDriver.Type)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_mode",
			Help: "**UPSMODE** The mode in which apcupsd is operating as specified in the configuration file ('UPSMODE'). " + typesToDesc(model.ModeTypes),
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.UpsMode.Type)
		},
	},

	// Shutdown
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_shutdown_battery_min",
			Help: "**MBATTCHG** If the battery charge percentage (BCHARGE) drops below this value, apcupsd will  shutdown your system.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return model.State.ShutdownBatteryMin
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_shutdown_timeleft_min",
			Help: "**MINTIMEL** (seconds) apcupsd will shutdown your system if the remaining runtime equals or is below this point.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.ShutdownTimeleftSecondsMin)
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_shutdown_onbattery_time_max",
			Help: "**MAXTIME** (seconds) apcupsd will shutdown your system if the time on batteries exceeds this value. A value of zero disables the feature.",
		}),
		ValFunc: func(m *Metric, model *model.Model) float64 {
			return float64(model.State.ShutdownOnBatterySecondsMax)
		},
	},
}

// RegisterPermanents registering permanents
func RegisterPermanents() {
	for _, m := range Metrics {
		if m.IsPermanent {
			m.Register()
		}
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
