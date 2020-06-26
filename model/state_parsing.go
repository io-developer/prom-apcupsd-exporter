package model

import (
	"local/apcupsd_exporter/apcupsd"
	"time"
)

// NewStateFromOutput ..
func NewStateFromOutput(o *apcupsd.Output) *State {
	return &State{
		// input
		InputSensivity: Sensivity{
			Text: o.Get("SENSE", ""),
			Type: o.GetMapped("SENSE", SensivityTypes,
				SensivityType(0),
			).(SensivityType),
		},
		InputFrequency:           o.GetFloat("LINEFREQ", 0),
		InputVoltage:             o.GetFloat("LINEV", 0),
		InputVoltageMin:          o.GetFloat("MINLINEV", 0),
		InputVoltageMax:          o.GetFloat("MAXLINEV", 0),
		InputVoltageNominal:      o.GetFloat("NOMINV", 0),
		InputVoltageTransferLow:  o.GetFloat("LOTRANS", 0),
		InputVoltageTransferHigh: o.GetFloat("HITRANS", 0),

		// output
		OutputLoad:           o.GetFloat("LOADPCT", 0),
		OutputAmps:           o.GetFloat("OUTCURNT", 0),
		OutputPowerNominal:   o.GetFloat("NOMPOWER", 0),
		OutputVoltage:        o.GetFloat("OUTPUTV", 0),
		OutputVoltageNominal: o.GetFloat("NOMOUTV", 0),

		// battery
		BatteryCharge:         o.GetFloat("BCHARGE", 0),
		BatteryVoltage:        o.GetFloat("BATTV", 0),
		BatteryVoltageNominal: o.GetFloat("NOMBATTV", 0),
		BatteryExternalCount:  uint16(o.GetUint("EXTBATTS", 0)),
		BatteryBadCount:       uint16(o.GetUint("BADBATTS", 0)),
		BatteryReplacedDate:   o.GetTime("BATTDATE", time.Time{}),

		// ups
		UpsManafacturedDate: o.GetTime("MANDATE", time.Time{}),
		UpsModel:            o.Get("MODEL", ""),
		UpsSerial:           o.Get("SERIALNO", ""),
		UpsFirmware:         o.Get("FIRMWARE", ""),
		UpsName:             o.Get("UPSNAME", ""),

		UpsStatus: NewStatus(
			o.GetUint("STATFLAG", 0),
			o.Get("STATUS", ""),
		),

		UpsDipSwitchFlag: o.GetUint("DIPSW", 0),

		UpsReg1: o.GetUint("REG1", 0),
		UpsReg2: o.GetUint("REG2", 0),
		UpsReg3: o.GetUint("REG3", 0),

		UpsTimeleftSeconds:           o.GetSeconds("TIMELEFT", 0),
		UpsTimeleftSecondsLowBattery: o.GetSeconds("DLOWBATT", 0),
		UpsTransferOnBatteryCount:    o.GetUint("NUMXFERS", 0),

		UpsTransferOnBatteryReason: TransferOnbatteryReason{
			Text: o.Get("LASTXFER", ""),
			Type: o.GetMapped("LASTXFER", TransferOnbatteryReasonTypes,
				TransferOnbatteryReasonType(0),
			).(TransferOnbatteryReasonType),
		},

		UpsTransferOnBatteryDate:  o.GetTime("XONBATT", time.Time{}),
		UpsTransferOffBatteryDate: o.GetTime("XOFFBATT", time.Time{}),

		UpsOnBatterySeconds:           o.GetSeconds("TONBATT", 0),
		UpsOnBatterySecondsCumulative: o.GetSeconds("CUMONBATT", 0),

		UpsTurnOffDelaySeconds: o.GetSeconds("DSHUTD", 0),
		UpsTurnOnDelaySeconds:  o.GetSeconds("DWAKE", 0),
		UpsTurnOnBatteryMin:    o.GetFloat("RETPCT", 0),

		UpsTempInternal: o.GetFloat("ITEMP", 0),
		UpsTempAmbient:  o.GetFloat("AMBTEMP", 0),
		UpsHumidity:     o.GetFloat("HUMIDITY", 0),

		UpsAlarmMode: AlarmMode{
			Text: o.Get("ALARMDEL", ""),
			Type: o.GetMapped("ALARMDEL", AlarmModeTypes,
				AlarmModeType(0),
			).(AlarmModeType),
		},

		UpsSelftestResult: SelftestResult{
			Text: o.Get("SELFTEST", ""),
			Type: o.GetMapped("SELFTEST", SelftestResultTypes,
				SelftestResultType(0),
			).(SelftestResultType),
		},
		UpsSelftestIntervalSeconds: o.GetSeconds("STESTI", 0),

		UpsCable: Cable{
			Text: o.Get("CABLE", ""),
			Type: o.GetMapped("CABLE", CableTypes, CableType(0)).(CableType),
		},
		UpsDriver: Driver{
			Text: o.Get("DRIVER", ""),
			Type: o.GetMapped("DRIVER", DriverTypes, DriverType(0)).(DriverType),
		},
		UpsMode: Mode{
			Text: o.Get("UPSMODE", ""),
			Type: o.GetMapped("UPSMODE", ModeTypes, ModeType(0)).(ModeType),
		},

		// shutdown
		ShutdownBatteryMin:          o.GetFloat("MBATTCHG", 0),
		ShutdownTimeleftSecondsMin:  o.GetSeconds("MINTIMEL", 0),
		ShutdownOnBatterySecondsMax: o.GetSeconds("MAXTIME", 0),

		// apcupsd
		ApcupsdHost:    o.Get("HOSTNAME", ""),
		ApcupsdVersion: o.Get("VERSION", ""),
	}
}
