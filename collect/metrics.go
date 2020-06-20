package collect

import "github.com/prometheus/client_golang/prometheus"

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
CABLE    : USB Cable
DRIVER   : USB UPS Driver
UPSMODE  : Stand Alone
STARTTIME: 2020-06-18 02:00:02 +0300
MODEL    : Back-UPS CS 500
STATUS   : ONLINE

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



ALARMDEL : No alarm|30 Seconds

LASTXFER : Unacceptable line voltage changes
| No transfers since turnon
| Automatic or explicit self test


SELFTEST : NO
STESTI   : None|14 days
STATFLAG : 0x05000008
SERIALNO : 4B1802P05216
FIRMWARE : 808.q10 .I USB FW:q

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
*/

// Metrics declare
var Metrics = []*Metric{

	// Input
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_sensitivity",
			Help: "**SENSE** The sensitivity level of the UPS to line voltage fluctuations.\n Unknown=0, Low=1, Medium=2, High=3, 'Auto Adjust'=4",
		}),
		OutputKey: "SENSE",
		Type:      "valueMap",
		ValueMap: map[string]float64{
			"Unknown":     0,
			"Low":         1,
			"Medium":      2,
			"High":        3,
			"Auto Adjust": 4,
		},
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_frequency",
			Help: "**LINEFREQ** Line frequency in hertz as given by the UPS.",
		}),
		OutputKey: "LINEFREQ",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage",
			Help: "**LINEV** The current line voltage as returned by the UPS.",
		}),
		OutputKey: "LINEV",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_min",
			Help: "**MINLINEV** The minimum line voltage since the UPS was started, as returned by the UPS",
		}),
		OutputKey: "MINLINEV",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_max",
			Help: "**MAXLINEV** The maximum line voltage since the UPS was started, as reported by the UPS",
		}),
		OutputKey: "MAXLINEV",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_nominal",
			Help: "**NOMINV** The input voltage that the UPS is configured to expect.",
		}),
		OutputKey: "NOMINV",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_transfer_low",
			Help: "**LOTRANS** The line voltage below which the UPS will switch to batteries.",
		}),
		OutputKey: "LOTRANS",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_transfer_high",
			Help: "**HITRANS** The line voltage above which the UPS will switch to batteries.",
		}),
		OutputKey: "HITRANS",
	},

	// Output
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_load",
			Help: "**LOADPCT** The percentage of load capacity as estimated by the UPS.",
		}),
		OutputKey: "LOADPCT",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_power_nominal",
			Help: "**NOMPOWER** The maximum power in Watts that the UPS is designed to supply.",
		}),
		OutputKey: "NOMPOWER",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_voltage",
			Help: "**OUTPUTV** The voltage the UPS is supplying to your equipment",
		}),
		OutputKey: "OUTPUTV",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_voltage_nominal",
			Help: "**NOMOUTV** The output voltage that the UPS will attempt to supply when on battery power.",
		}),
		OutputKey: "NOMOUTV",
	},

	// Battery
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_charge",
			Help: "**BCHARGE** The percentage charge on the batteries.",
		}),
		OutputKey: "BCHARGE",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_voltage",
			Help: "**BATTV** Battery voltage as supplied by the UPS.",
		}),
		OutputKey: "BATTV",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_voltage_nominal",
			Help: "**NOMBATTV** The nominal battery voltage.",
		}),
		OutputKey: "NOMBATTV",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_replaced_timestamp",
			Help: "**BATTDATE** The date that batteries were last replaced.",
		}),
		OutputKey: "BATTDATE",
		Type:      "date",
	},

	// Ups
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_manafactured_timestamp",
			Help: "**MANDATE** The date the UPS was manufactured.",
		}),
		OutputKey: "MANDATE",
		Type:      "date",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_status_flag",
			Help: "**STATFLAG** Status flag. English version is given by STATUS.",
		}),
		OutputKey: "STATFLAG",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_dip_switch_flag",
			Help: "**DIPSW** The current dip switch settings on UPSes that have them.",
		}),
		OutputKey: "DIPSW",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_reg1",
			Help: "**REG1** The value from the UPS fault register 1.",
		}),
		OutputKey: "REG1",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_reg2",
			Help: "**REG2** The value from the UPS fault register 2.",
		}),
		OutputKey: "REG2",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_reg3",
			Help: "**REG3** The value from the UPS fault register 3.",
		}),
		OutputKey: "REG3",
	},

	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_temp_internal",
			Help: "**ITEMP** (Celsius) Internal UPS temperature as supplied by the UPS.",
		}),
		OutputKey: "ITEMP",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_timeleft",
			Help: "**TIMELEFT** (seconds) The remaining runtime left on batteries as estimated by the UPS.",
		}),
		OutputKey: "TIMELEFT",
		Type:      "minutes",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_timeleft_low_battery",
			Help: "**DLOWBATT** (seconds) The remaining runtime below which the UPS sends the low battery signal. At this point apcupsd will force an immediate emergency shutdown.",
		}),
		OutputKey: "DLOWBATT",
		Type:      "minutes",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery_count",
			Help: "**NUMXFERS** The number of transfers to batteries since apcupsd startup.",
		}),
		OutputKey: "NUMXFERS",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery_time",
			Help: "**TONBATT** Time in seconds currently on batteries",
		}),
		OutputKey: "TONBATT",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery_time_cumulative",
			Help: "Total (cumulative) time on batteries in seconds since apcupsd startup.",
		}),
		OutputKey: "CUMONBATT",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery_timestamp",
			Help: "**XONBATT** Time and date of last transfer to batteries",
		}),
		OutputKey: "XONBATT",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_offbattery_timestamp",
			Help: "Time and date of last transfer from batteries",
		}),
		OutputKey: "XOFFBATT",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_turnon_delay",
			Help: "**DWAKE** (seconds) The amount of time the UPS will wait before restoring power to your equipment after a power off condition when the power is restored.",
		}),
		OutputKey: "DWAKE",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_turnon_battery_min",
			Help: "	**RETPCT** The percentage charge that the batteries must have after a power off condition before the UPS will restore power to your equipment.",
		}),
		OutputKey: "RETPCT",
	},

	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_turnoff_delay",
			Help: "**DSHUTD** (seconds) The grace delay that the UPS gives after receiving a power down command from apcupsd before it powers off your equipment.",
		}),
		OutputKey: "DSHUTD",
	},

	// Shutdown
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_shutdown_battery_min",
			Help: "**MBATTCHG** If the battery charge percentage (BCHARGE) drops below this value, apcupsd will  shutdown your system.",
		}),
		OutputKey: "MBATTCHG",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_shutdown_timeleft_min",
			Help: "**MINTIMEL** (seconds) apcupsd will shutdown your system if the remaining runtime equals or is below this point.",
		}),
		OutputKey: "MINTIMEL",
		Type:      "minutes",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_shutdown_onbattery_time_max",
			Help: "**MAXTIME** (seconds) apcupsd will shutdown your system if the time on batteries exceeds this value. A value of zero disables the feature.",
		}),
		OutputKey: "MAXTIME",
		Type:      "minutes",
	},
}

// MetricsRegister registering permanents
func MetricsRegister() {
	for _, m := range Metrics {
		if m.IsPermanent {
			m.Gauge.Set(m.DefaultValue)
			m.Register()
		}
	}
}
