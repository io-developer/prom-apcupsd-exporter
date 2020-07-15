package model

import (
	"local/apcupsd_exporter/apcupsd"
)

var defaultState = &State{}

// NewStateFromOutput ..
func NewStateFromOutput(o *apcupsd.Output, def *State) State {
	if def == nil {
		def = defaultState
	}
	return State{
		// input
		InputSensivity: Sensivity{
			Text: o.Get("SENSE", def.InputSensivity.Text),
			Type: o.GetMapped("SENSE", SensivityTypes, def.InputSensivity.Type).(SensivityType),
		},
		InputFrequency:           o.GetFloat("LINEFREQ", def.InputFrequency),
		InputVoltage:             o.GetFloat("LINEV", def.InputVoltage),
		InputVoltageMin:          o.GetFloat("MINLINEV", def.InputVoltageMin),
		InputVoltageMax:          o.GetFloat("MAXLINEV", def.InputVoltageMax),
		InputVoltageNominal:      o.GetFloat("NOMINV", def.InputVoltageNominal),
		InputVoltageTransferLow:  o.GetFloat("LOTRANS", def.InputVoltageTransferLow),
		InputVoltageTransferHigh: o.GetFloat("HITRANS", def.InputVoltageTransferHigh),

		// output
		OutputLoad:                 o.GetFloat("LOADPCT", def.OutputLoad),
		OutputAmps:                 o.GetFloat("OUTCURNT", def.OutputAmps),
		OutputPowerNominal:         o.GetFloat("NOMPOWER", def.OutputPowerNominal),
		OutputPowerApparentNominal: o.GetFloat("NOMAPNT", def.OutputPowerApparentNominal),
		OutputVoltage:              o.GetFloat("OUTPUTV", def.OutputVoltage),
		OutputVoltageNominal:       o.GetFloat("NOMOUTV", def.OutputVoltageNominal),

		// battery
		BatteryCharge:         o.GetFloat("BCHARGE", def.BatteryCharge),
		BatteryVoltage:        o.GetFloat("BATTV", def.BatteryVoltage),
		BatteryVoltageNominal: o.GetFloat("NOMBATTV", def.BatteryVoltageNominal),
		BatteryExternalCount:  uint16(o.GetUint("EXTBATTS", uint64(def.BatteryExternalCount))),
		BatteryBadCount:       uint16(o.GetUint("BADBATTS", uint64(def.BatteryBadCount))),
		BatteryReplacedDate:   o.GetTime("BATTDATE", def.BatteryReplacedDate),

		// ups
		UpsManafacturedDate: o.GetTime("MANDATE", def.UpsManafacturedDate),
		UpsModel:            o.Get("MODEL", def.UpsModel),
		UpsSerial:           o.Get("SERIALNO", def.UpsSerial),
		UpsFirmware:         o.Get("FIRMWARE", def.UpsFirmware),
		UpsName:             o.Get("UPSNAME", def.UpsName),

		UpsStatus: NewStatus(
			o.GetUint("STATFLAG", def.UpsStatus.Flag),
			o.Get("STATUS", def.UpsStatus.Text),
		),

		UpsDipSwitchFlag: o.GetUint("DIPSW", def.UpsDipSwitchFlag),

		UpsReg1: o.GetUint("REG1", def.UpsReg1),
		UpsReg2: o.GetUint("REG2", def.UpsReg2),
		UpsReg3: o.GetUint("REG3", def.UpsReg3),

		UpsTimeleftSeconds:           o.GetSeconds("TIMELEFT", def.UpsTimeleftSeconds),
		UpsTimeleftSecondsLowBattery: o.GetSeconds("DLOWBATT", def.UpsTimeleftSecondsLowBattery),
		UpsTransferOnBatteryCount:    o.GetUint("NUMXFERS", def.UpsTransferOnBatteryCount),

		UpsTransferOnBatteryReason: TransferOnbatteryReason{
			Text: o.Get("LASTXFER", def.UpsTransferOnBatteryReason.Text),
			Type: o.GetMapped("LASTXFER", TransferOnbatteryReasonTypes,
				def.UpsTransferOnBatteryReason.Type,
			).(TransferOnbatteryReasonType),
		},

		UpsTransferOnBatteryDate:  o.GetTime("XONBATT", def.UpsTransferOnBatteryDate),
		UpsTransferOffBatteryDate: o.GetTime("XOFFBATT", def.UpsTransferOnBatteryDate),

		UpsOnBatterySeconds:           o.GetSeconds("TONBATT", def.UpsOnBatterySeconds),
		UpsOnBatterySecondsCumulative: o.GetSeconds("CUMONBATT", def.UpsOnBatterySecondsCumulative),

		UpsTurnOffDelaySeconds: o.GetSeconds("DSHUTD", def.UpsTurnOffDelaySeconds),
		UpsTurnOnDelaySeconds:  o.GetSeconds("DWAKE", def.UpsTurnOnDelaySeconds),
		UpsTurnOnBatteryMin:    o.GetFloat("RETPCT", def.UpsTurnOnBatteryMin),

		UpsTempInternal: o.GetFloat("ITEMP", def.UpsTempInternal),
		UpsTempAmbient:  o.GetFloat("AMBTEMP", def.UpsTempAmbient),
		UpsHumidity:     o.GetFloat("HUMIDITY", def.UpsHumidity),

		UpsAlarmMode: AlarmMode{
			Text: o.Get("ALARMDEL", def.UpsAlarmMode.Text),
			Type: o.GetMapped("ALARMDEL", AlarmModeTypes,
				def.UpsAlarmMode.Type,
			).(AlarmModeType),
		},

		UpsSelftestResult: SelftestResult{
			Text: o.Get("SELFTEST", def.UpsSelftestResult.Text),
			Type: o.GetMapped("SELFTEST", SelftestResultTypes,
				def.UpsSelftestResult.Type,
			).(SelftestResultType),
		},
		UpsSelftestIntervalSeconds: o.GetSeconds("STESTI", def.UpsSelftestIntervalSeconds),

		UpsCable: Cable{
			Text: o.Get("CABLE", def.UpsCable.Text),
			Type: o.GetMapped("CABLE", CableTypes, def.UpsCable.Type).(CableType),
		},
		UpsDriver: Driver{
			Text: o.Get("DRIVER", def.UpsDriver.Text),
			Type: o.GetMapped("DRIVER", DriverTypes, def.UpsDriver.Type).(DriverType),
		},
		UpsMode: Mode{
			Text: o.Get("UPSMODE", def.UpsMode.Text),
			Type: o.GetMapped("UPSMODE", ModeTypes, def.UpsMode.Type).(ModeType),
		},

		// shutdown
		ShutdownBatteryMin:          o.GetFloat("MBATTCHG", def.ShutdownBatteryMin),
		ShutdownTimeleftSecondsMin:  o.GetSeconds("MINTIMEL", def.ShutdownTimeleftSecondsMin),
		ShutdownOnBatterySecondsMax: o.GetSeconds("MAXTIME", def.ShutdownOnBatterySecondsMax),

		// apcupsd
		ApcupsdHost:      o.Get("HOSTNAME", def.ApcupsdHost),
		ApcupsdVersion:   o.Get("VERSION", def.ApcupsdVersion),
		ApcupsdStartTime: o.GetTime("STARTTIME", def.ApcupsdStartTime),
	}
}
