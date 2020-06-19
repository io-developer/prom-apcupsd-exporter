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
XOFFBATT : N/A
SENSE    : Low|High
LOTRANS  : 180.0 Volts
HITRANS  : 260.0 Volts
BATTV    : 13.7 Volts
NOMBATTV : 12.0 Volts
ITEMP    : 29.2 C
BATTDATE : 2018-11-29
MANDATE  : 2018-01-09

DWAKE    : 0 Seconds
DSHUTD   : 180 Seconds
RETPCT   : 0.0 Percent
ALARMDEL : No alarm|30 Seconds

LASTXFER : Unacceptable line voltage changes
| No transfers since turnon

NUMXFERS : 0

SELFTEST : NO
STESTI   : None|14 days
STATFLAG : 0x05000008
SERIALNO : 4B1802P05216
FIRMWARE : 808.q10 .I USB FW:q


*/

// Metrics declare
var Metrics = []*Metric{

	// Battery
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_charge",
			Help: "Current battery charge (percent)",
		}),
		OutputKey: "BCHARGE",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_voltage",
			Help: "",
		}),
		OutputKey: "BATTV",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_voltage_nominal",
			Help: "",
		}),
		OutputKey: "NOMBATTV",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_datetime",
			Help: "ts",
		}),
		OutputKey: "BATTDATE",
		Type:      "date",
	},

	// Input
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_sensitivity",
			Help: "",
		}),
		OutputKey: "SENSE",
		Type:      "valueMap",
		ValueMap: map[string]float64{
			"Low":    0,
			"Normal": 1,
			"High":   2,
		},
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage",
			Help: "",
		}),
		OutputKey: "LINEV",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_nominal",
			Help: "",
		}),
		OutputKey: "NOMINV",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_frequency",
			Help: "",
		}),
		OutputKey: "LINEFREQ",
	},

	// Output
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_load",
			Help: "",
		}),
		OutputKey: "LOADPCT",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_power_nominal",
			Help: "",
		}),
		OutputKey: "NOMPOWER",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_voltage",
			Help: "",
		}),
		OutputKey: "OUTPUTV",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_voltage_nominal",
			Help: "",
		}),
		OutputKey: "NOMOUTV",
	},

	// Ups
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_manafactured_datetime",
			Help: "ts",
		}),
		OutputKey: "MANDATE",
		Type:      "date",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_temp_internal",
			Help: "",
		}),
		OutputKey: "ITEMP",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_voltage_low",
			Help: "",
		}),
		OutputKey: "LOTRANS",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_voltage_high",
			Help: "",
		}),
		OutputKey: "HITRANS",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_timeleft",
			Help: "seconds",
		}),
		OutputKey: "TIMELEFT",
		Type:      "minutes",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_onbattery",
			Help: "seconds",
		}),
		OutputKey: "TONBATT",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_onbattery_cumulative",
			Help: "seconds",
		}),
		OutputKey: "CUMONBATT",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_offbattery_datetime",
			Help: "ts",
		}),
		OutputKey: "CUMONBATT",
	},

	// Shutdown
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_shutdown_min_charge",
			Help: "",
		}),
		OutputKey: "MBATTCHG",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_shutdown_min_timeleft",
			Help: "seconds",
		}),
		OutputKey: "MINTIMEL",
		Type:      "minutes",
	},
	{
		Gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_shutdown_max_time",
			Help: "seconds",
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
