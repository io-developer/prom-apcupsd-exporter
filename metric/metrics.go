package metric

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
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
			Help: "**SENSE** The sensitivity level of the UPS to line voltage fluctuations." +
				" Low=1, Medium=2, High=3, 'Auto Adjust'=4, Unknown=5",
		}),
		Handler: DefaultHandler{
			ApcKey: "SENSE",
			ValueMap: map[string]float64{
				"Low":         1,
				"Medium":      2,
				"High":        3,
				"Auto Adjust": 4,
				"Unknown":     5,
			},
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_frequency",
			Help: "**LINEFREQ** Line frequency in hertz as given by the UPS.",
		}),
		Handler: NewDefaultHandler("LINEFREQ"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage",
			Help: "**LINEV** The current line voltage as returned by the UPS.",
		}),
		Handler: NewDefaultHandler("LINEV"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_min",
			Help: "**MINLINEV** The minimum line voltage since the UPS was started, as returned by the UPS",
		}),
		Handler: NewDefaultHandler("MINLINEV"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_max",
			Help: "**MAXLINEV** The maximum line voltage since the UPS was started, as reported by the UPS",
		}),
		Handler: NewDefaultHandler("MAXLINEV"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_nominal",
			Help: "**NOMINV** The input voltage that the UPS is configured to expect.",
		}),
		Handler: NewDefaultHandler("NOMINV"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_transfer_low",
			Help: "**LOTRANS** The line voltage below which the UPS will switch to batteries.",
		}),
		Handler: NewDefaultHandler("LOTRANS"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_transfer_high",
			Help: "**HITRANS** The line voltage above which the UPS will switch to batteries.",
		}),
		Handler: NewDefaultHandler("HITRANS"),
	},

	// Output
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_load",
			Help: "**LOADPCT** The percentage of load capacity as estimated by the UPS.",
		}),
		Handler: NewDefaultHandler("LOADPCT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_amps",
			Help: "**OUTCURNT** Amps",
		}),
		Handler: NewDefaultHandler("OUTCURNT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_power_nominal",
			Help: "**NOMPOWER** The maximum power in Watts that the UPS is designed to supply.",
		}),
		Handler: NewDefaultHandler("NOMPOWER"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_voltage",
			Help: "**OUTPUTV** The voltage the UPS is supplying to your equipment",
		}),
		Handler: NewDefaultHandler("OUTPUTV"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_voltage_nominal",
			Help: "**NOMOUTV** The output voltage that the UPS will attempt to supply when on battery power.",
		}),
		Handler: NewDefaultHandler("NOMOUTV"),
	},

	// Battery
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_charge",
			Help: "**BCHARGE** The percentage charge on the batteries.",
		}),
		Handler: NewDefaultHandler("BCHARGE"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_voltage",
			Help: "**BATTV** Battery voltage as supplied by the UPS.",
		}),
		Handler: NewDefaultHandler("BATTV"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_voltage_nominal",
			Help: "**NOMBATTV** The nominal battery voltage.",
		}),
		Handler: NewDefaultHandler("NOMBATTV"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_external",
			Help: "**EXTBATTS** The number of external batteries as defined by the user. A correct number here helps the UPS compute the remaining runtime more accurately.",
		}),
		Handler: NewDefaultHandler("EXTBATTS"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_bad",
			Help: "**BADBATTS** The number of bad battery packs.",
		}),
		Handler: NewDefaultHandler("BADBATTS"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_replaced_timestamp",
			Help: "**BATTDATE** The date that batteries were last replaced.",
		}),
		Handler: NewDefaultHandler("BATTDATE"),
	},

	// Ups
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_manafactured_timestamp",
			Help: "**MANDATE** The date the UPS was manufactured.",
		}),
		Handler: NewDefaultHandler("MANDATE"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_status_flags",
			Help: "**STATFLAG** Current status flags",
		}),
		Handler: NewDefaultHandler("STATFLAG"),
	},
	{
		Collector: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "apcupsd_ups_status",
			Help: "**STATFLAG** Labeled current status flags",
		}, []string{"flag"}),
		Handler: StatusComponentHandler{},
	},
	{
		Collector: prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Name:       "apcupsd_ups_status_trace_1m",
			Help:       "Labeled flags active for last 1 minute ({quantile=1} > 0)",
			MaxAge:     1 * time.Minute,
			AgeBuckets: 2,
			Objectives: map[float64]float64{1: 0},
		}, []string{"flag"}),
		Handler: StatusTraceComponentHandler{},
	},
	{
		Collector: prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Name:       "apcupsd_ups_status_trace_3m",
			Help:       "Labeled flags active for last 3 minutes ({quantile=1} > 0)",
			MaxAge:     3 * time.Minute,
			AgeBuckets: 2,
			Objectives: map[float64]float64{1: 0},
		}, []string{"flag"}),
		Handler: StatusTraceComponentHandler{},
	},
	{
		Collector: prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Name:       "apcupsd_ups_status_trace_10m",
			Help:       "Labeled flags active for last 10 minutes ({quantile=1} > 0)",
			MaxAge:     10 * time.Minute,
			AgeBuckets: 2,
			Objectives: map[float64]float64{1: 0},
		}, []string{"flag"}),
		Handler: StatusTraceComponentHandler{},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_dip_switch_flag",
			Help: "**DIPSW** The current dip switch settings on UPSes that have them.",
		}),
		Handler: NewDefaultHandler("DIPSW"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_reg1",
			Help: "**REG1** The value from the UPS fault register 1.",
		}),
		Handler: NewDefaultHandler("REG1"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_reg2",
			Help: "**REG2** The value from the UPS fault register 2.",
		}),
		Handler: NewDefaultHandler("REG2"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_reg3",
			Help: "**REG3** The value from the UPS fault register 3.",
		}),
		Handler: NewDefaultHandler("REG3"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_timeleft",
			Help: "**TIMELEFT** (seconds) The remaining runtime left on batteries as estimated by the UPS.",
		}),
		Handler: NewDefaultHandler("TIMELEFT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_timeleft_low_battery",
			Help: "**DLOWBATT** (seconds) The remaining runtime below which the UPS sends the low battery signal. At this point apcupsd will force an immediate emergency shutdown.",
		}),
		Handler: NewDefaultHandler("DLOWBATT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery",
			Help: "**NUMXFERS** The number of transfers to batteries since apcupsd startup.",
		}),
		Handler: NewDefaultHandler("NUMXFERS"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery_reason",
			Help: "**LASTXFER** The reason for the last transfer to batteries." +
				" 'No transfers since turnon'=1," +
				" 'Automatic or explicit self test'=2," +
				" 'Forced by software'=3," +
				" 'Low line voltage'=4," +
				" 'High line voltage'=5," +
				" 'Unacceptable line voltage changes'=6," +
				" 'Line voltage notch or spike'=7," +
				" 'Input frequency out of range'=8," +
				" 'UNKNOWN EVENT'=9",
		}),
		Handler: DefaultHandler{
			ApcKey: "LASTXFER",
			ValueMap: map[string]float64{
				"No transfers since turnon":         1,
				"Automatic or explicit self test":   2,
				"Forced by software":                3,
				"Low line voltage":                  4,
				"High line voltage":                 5,
				"Unacceptable line voltage changes": 6,
				"Line voltage notch or spike":       7,
				"Input frequency out of range":      8,
				"UNKNOWN EVENT":                     9,
			},
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery_time",
			Help: "**TONBATT** Time in seconds currently on batteries",
		}),
		Handler: NewDefaultHandler("TONBATT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery_time_cumulative",
			Help: "Total (cumulative) time on batteries in seconds since apcupsd startup.",
		}),
		Handler: NewDefaultHandler("CUMONBATT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery_timestamp",
			Help: "**XONBATT** Time and date of last transfer to batteries",
		}),
		Handler: NewDefaultHandler("XONBATT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_offbattery_timestamp",
			Help: "Time and date of last transfer from batteries",
		}),
		Handler: NewDefaultHandler("XOFFBATT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_turnon_delay",
			Help: "**DWAKE** (seconds) The amount of time the UPS will wait before restoring power to your equipment after a power off condition when the power is restored.",
		}),
		Handler: NewDefaultHandler("DWAKE"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_turnon_battery_min",
			Help: "	**RETPCT** The percentage charge that the batteries must have after a power off condition before the UPS will restore power to your equipment.",
		}),
		Handler: NewDefaultHandler("RETPCT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_turnoff_delay",
			Help: "**DSHUTD** (seconds) The grace delay that the UPS gives after receiving a power down command from apcupsd before it powers off your equipment.",
		}),
		Handler: NewDefaultHandler("DSHUTD"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_temp_internal",
			Help: "**ITEMP** (Celsius) Internal UPS temperature as supplied by the UPS.",
		}),
		Handler: NewDefaultHandler("ITEMP"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_temp_ambient",
			Help: "**AMBTEMP** The ambient temperature as measured by the UPS.",
		}),
		Handler: NewDefaultHandler("AMBTEMP"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_humidity",
			Help: "**HUMIDITY** The humidity as measured by the UPS.",
		}),
		Handler: NewDefaultHandler("HUMIDITY"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_alarm_mode",
			Help: "**ALARMDEL** The delay period for the UPS alarm.\n" +
				"'No alarm'=1, 'Always'=2, '5 Seconds'=3, '30 Seconds'=4, 'Low Battery'=5",
		}),
		Handler: DefaultHandler{
			ApcKey: "ALARMDEL",
			ValueMap: map[string]float64{
				"No alarm":    1,
				"Always":      2,
				"5 Seconds":   3,
				"5":           3,
				"30 Seconds":  4,
				"30":          4,
				"Low Battery": 5,
			},
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_selftest_result",
			Help: "**SELFTEST** The results of the last self test, and may have the following values." +
				" NO=1 No results i.e. no self test performed in the last 5 minutes," +
				" OK=2 self test indicates good battery," +
				" BT=3 self test failed due to insufficient battery capacity," +
				" NG=4 self test failed due to overload," +
				" IP=5 INPROGRESS," +
				" WN=6 WARNING," +
				" ??=7 UNKNOWN",
		}),
		Handler: DefaultHandler{
			ApcKey: "SELFTEST",
			ValueMap: map[string]float64{
				"NO": 1,
				"OK": 2,
				"BT": 3,
				"NG": 4,
				"IP": 5,
				"WN": 6,
				"??": 7,
			},
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_selftest_interval",
			Help: "**STESTI** The interval in seconds between automatic self tests.",
		}),
		Handler: NewDefaultHandler("STESTI"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_cable",
			Help: "**CABLE** The cable as specified in the configuration file ('UPSCABLE')." +
				" 'Custom Cable Simple'=1," +
				" 'APC Cable 940-0119A'=2," +
				" 'APC Cable 940-0127A'=3," +
				" 'APC Cable 940-0128A'=4," +
				" 'APC Cable 940-0020B'=5," +
				" 'APC Cable 940-0020C'=6," +
				" 'APC Cable 940-0023A'=7," +
				" 'MAM Cable 04-02-2000'=8," +
				" 'APC Cable 940-0095A'=9," +
				" 'APC Cable 940-0095B'=10," +
				" 'APC Cable 940-0095C'=11," +
				" 'Custom Cable Smart'=12," +
				" 'APC Cable 940-0024B'=121," +
				" 'APC Cable 940-0024C'=122," +
				" 'APC Cable 940-1524C'=123," +
				" 'APC Cable 940-0024G'=124," +
				" 'APC Cable 940-0625A'=125," +
				" 'Ethernet Link'=13," +
				" 'USB Cable'=14",
		}),
		Handler: DefaultHandler{
			ApcKey: "CABLE",
			ValueMap: map[string]float64{
				"Custom Cable Simple":  1,
				"APC Cable 940-0119A":  2,
				"APC Cable 940-0127A":  3,
				"APC Cable 940-0128A":  4,
				"APC Cable 940-0020B":  5,
				"APC Cable 940-0020C":  6,
				"APC Cable 940-0023A":  7,
				"MAM Cable 04-02-2000": 8,
				"APC Cable 940-0095A":  9,
				"APC Cable 940-0095B":  10,
				"APC Cable 940-0095C":  11,
				"Custom Cable Smart":   12,
				"APC Cable 940-0024B":  121,
				"APC Cable 940-0024C":  122,
				"APC Cable 940-1524C":  123,
				"APC Cable 940-0024G":  124,
				"APC Cable 940-0625A":  125,
				"Ethernet Link":        13,
				"USB Cable":            14,
			},
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_driver",
			Help: "**DRIVER** type." +
				" 'DUMB UPS Driver'=1," +
				" 'APC Smart UPS (any)'=2," +
				" 'USB UPS Driver'=3," +
				" 'NETWORK UPS Driver'=4," +
				" 'TEST UPS Driver'=5," +
				" 'PCNET UPS Driver'=6," +
				" 'SNMP UPS Driver'=7," +
				" 'MODBUS UPS Driver'=8",
		}),
		Handler: DefaultHandler{
			ApcKey: "DRIVER",
			ValueMap: map[string]float64{
				"DUMB UPS Driver":     1,
				"APC Smart UPS (any)": 2,
				"USB UPS Driver":      3,
				"NETWORK UPS Driver":  4,
				"TEST UPS Driver":     5,
				"PCNET UPS Driver":    6,
				"SNMP UPS Driver":     7,
				"MODBUS UPS Driver":   8,
			},
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_mode",
			Help: "**UPSMODE** The mode in which apcupsd is operating as specified in the configuration file ('UPSMODE')." +
				" 'Stand Alone'=1," +
				" 'ShareUPS Slave'=2," +
				" 'ShareUPS Master'=3",
		}),
		Handler: DefaultHandler{
			ApcKey: "UPSMODE",
			ValueMap: map[string]float64{
				"Stand Alone":     1,
				"ShareUPS Slave":  2,
				"ShareUPS Master": 3,
			},
		},
	},

	// Shutdown
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_shutdown_battery_min",
			Help: "**MBATTCHG** If the battery charge percentage (BCHARGE) drops below this value, apcupsd will  shutdown your system.",
		}),
		Handler: NewDefaultHandler("MBATTCHG"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_shutdown_timeleft_min",
			Help: "**MINTIMEL** (seconds) apcupsd will shutdown your system if the remaining runtime equals or is below this point.",
		}),
		Handler: NewDefaultHandler("MINTIMEL"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_shutdown_onbattery_time_max",
			Help: "**MAXTIME** (seconds) apcupsd will shutdown your system if the time on batteries exceeds this value. A value of zero disables the feature.",
		}),
		Handler: NewDefaultHandler("MAXTIME"),
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
