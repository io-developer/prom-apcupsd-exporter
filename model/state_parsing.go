package model

import (
	"local/apcupsd_exporter/apc"
	"time"
)

// NewStateFromOutput ..
func NewStateFromOutput(o *apc.Output) *State {
	return &State{
		// input
		InputSensivity: Sensivity{
			Type: o.GetMapped("SENSE", SensivityTypes, 0).(SensivityType),
			Text: o.Get("SENSE", ""),
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
		UpsManafacturedDate:          o.GetTime("MANDATE", time.Time{}),
		UpsStatus:                    NewStatus(o.GetUint("STATFLAG", 0), o.Get("STATUS", "")),
		UpsDipSwitchFlag:             o.GetUint("DIPSW", 0),
		UpsReg1:                      o.GetUint("REG1", 0),
		UpsReg2:                      o.GetUint("REG2", 0),
		UpsReg3:                      o.GetUint("REG3", 0),
		UpsTimeleftSeconds:           o.GetSeconds("TIMELEFT", 0),
		UpsTimeleftSecondsLowBattery: o.GetSeconds("DLOWBATT", 0),
		UpsTransferOnBatteryCount:    o.GetUint("NUMXFERS", 0),
		UpsTransferOnBatteryReason: TransferOnbatteryReason{
			Type: o.GetMapped("LASTXFER", TransferOnbatteryReasonTypes, 0).(TransferOnbatteryReasonType),
			Text: o.Get("LASTXFER", ""),
		},
		UpsTransferOnBatteryDate:      o.GetTime("XONBATT", time.Time{}),
		UpsTransferOffBatteryDate:     o.GetTime("XOFFBATT", time.Time{}),
		UpsOnBatterySeconds:           o.GetSeconds("TONBATT", 0),
		UpsOnBatterySecondsCumulative: o.GetSeconds("CUMONBATT", 0),
		UpsTurnOffDelaySeconds:        o.GetSeconds("DSHUTD", 0),
		UpsTurnOnDelaySeconds:         o.GetSeconds("DWAKE", 0),
		UpsTurnOnBatteryMin:           o.GetFloat("RETPCT", 0),
		UpsTempInternal:               o.GetFloat("ITEMP", 0),
		UpsTempAmbient:                o.GetFloat("AMBTEMP", 0),
		UpsHumidity:                   o.GetFloat("HUMIDITY", 0),
		UpsAlarmMode: AlarmMode{
			Type: o.GetMapped("ALARMDEL", AlarmModeTypes, 0).(AlarmModeType),
			Text: o.Get("ALARMDEL", ""),
		},
		UpsSelftestResult: SelftestResult{
			Type: o.GetMapped("SELFTEST", SelftestResultTypes, 0).(SelftestResultType),
			Text: o.Get("SELFTEST", ""),
		},
		UpsSelftestIntervalSeconds: o.GetSeconds("STESTI", 0),
		UpsCable: Cable{
			Type: o.GetMapped("CABLE", CableTypes, 0).(CableType),
			Text: o.Get("CABLE", ""),
		},
		UpsDriver: Driver{
			Type: o.GetMapped("DRIVER", DriverTypes, 0).(DriverType),
			Text: o.Get("DRIVER", ""),
		},
		UpsMode: Mode{
			Type: o.GetMapped("UPSMODE", ModeTypes, 0).(ModeType),
			Text: o.Get("UPSMODE", ""),
		},

		// shutdown
		ShutdownBatteryMin:          o.GetFloat("MBATTCHG", 0),
		ShutdownTimeleftSecondsMin:  o.GetSeconds("MINTIMEL", 0),
		ShutdownOnBatterySecondsMax: o.GetSeconds("MAXTIME", 0),
	}
}
