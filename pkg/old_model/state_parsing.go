package old_model

import (
	"github.com/io-developer/prom-apcupsd-exporter/pkg/dto"
	"github.com/io-developer/prom-apcupsd-exporter/pkg/parsing"
)

var defaultState = &State{}

// NewStateFromOutput ..
func NewStateFromOutput(response *dto.ApcaccessResponse, def *State) State {
	if def == nil {
		def = defaultState
	}

	reader := parsing.NewApcaccessResponseReader(response)

	return State{
		// input
		InputSensivity: Sensivity{
			Text: reader.GetValueAt("SENSE", def.InputSensivity.Text),
			Type: reader.GetMappedAt("SENSE", SensivityTypes, def.InputSensivity.Type).(SensivityType),
		},
		InputFrequency:           reader.GetFloatAt("LINEFREQ", def.InputFrequency),
		InputVoltage:             reader.GetFloatAt("LINEV", def.InputVoltage),
		InputVoltageMin:          reader.GetFloatAt("MINLINEV", def.InputVoltageMin),
		InputVoltageMax:          reader.GetFloatAt("MAXLINEV", def.InputVoltageMax),
		InputVoltageNominal:      reader.GetFloatAt("NOMINV", def.InputVoltageNominal),
		InputVoltageTransferLow:  reader.GetFloatAt("LOTRANS", def.InputVoltageTransferLow),
		InputVoltageTransferHigh: reader.GetFloatAt("HITRANS", def.InputVoltageTransferHigh),

		// output
		OutputLoad:                 reader.GetFloatAt("LOADPCT", def.OutputLoad),
		OutputAmps:                 reader.GetFloatAt("OUTCURNT", def.OutputAmps),
		OutputPowerNominal:         reader.GetFloatAt("NOMPOWER", def.OutputPowerNominal),
		OutputPowerApparentNominal: reader.GetFloatAt("NOMAPNT", def.OutputPowerApparentNominal),
		OutputVoltage:              reader.GetFloatAt("OUTPUTV", def.OutputVoltage),
		OutputVoltageNominal:       reader.GetFloatAt("NOMOUTV", def.OutputVoltageNominal),

		// battery
		BatteryCharge:         reader.GetFloatAt("BCHARGE", def.BatteryCharge),
		BatteryVoltage:        reader.GetFloatAt("BATTV", def.BatteryVoltage),
		BatteryVoltageNominal: reader.GetFloatAt("NOMBATTV", def.BatteryVoltageNominal),
		BatteryExternalCount:  uint16(reader.GetUintAt("EXTBATTS", uint64(def.BatteryExternalCount))),
		BatteryBadCount:       uint16(reader.GetUintAt("BADBATTS", uint64(def.BatteryBadCount))),
		BatteryReplacedDate:   reader.GetTimeAt("BATTDATE", def.BatteryReplacedDate),

		// ups
		UpsManafacturedDate: reader.GetTimeAt("MANDATE", def.UpsManafacturedDate),
		UpsModel:            reader.GetValueAt("MODEL", def.UpsModel),
		UpsSerial:           reader.GetValueAt("SERIALNO", def.UpsSerial),
		UpsFirmware:         reader.GetValueAt("FIRMWARE", def.UpsFirmware),
		UpsName:             reader.GetValueAt("UPSNAME", def.UpsName),

		UpsStatus: NewStatus(
			reader.GetUintAt("STATFLAG", def.UpsStatus.Flag),
			reader.GetValueAt("STATUS", def.UpsStatus.Text),
		),

		UpsDipSwitchFlag: reader.GetUintAt("DIPSW", def.UpsDipSwitchFlag),

		UpsReg1: reader.GetUintAt("REG1", def.UpsReg1),
		UpsReg2: reader.GetUintAt("REG2", def.UpsReg2),
		UpsReg3: reader.GetUintAt("REG3", def.UpsReg3),

		UpsTimeleftSeconds:           reader.GetDurationSecondsAt("TIMELEFT", def.UpsTimeleftSeconds),
		UpsTimeleftSecondsLowBattery: reader.GetDurationSecondsAt("DLOWBATT", def.UpsTimeleftSecondsLowBattery),
		UpsTransferOnBatteryCount:    reader.GetUintAt("NUMXFERS", def.UpsTransferOnBatteryCount),

		UpsTransferOnBatteryReason: TransferOnbatteryReason{
			Text: reader.GetValueAt("LASTXFER", def.UpsTransferOnBatteryReason.Text),
			Type: reader.GetMappedAt("LASTXFER", TransferOnbatteryReasonTypes,
				def.UpsTransferOnBatteryReason.Type,
			).(TransferOnbatteryReasonType),
		},

		UpsTransferOnBatteryDate:  reader.GetTimeAt("XONBATT", def.UpsTransferOnBatteryDate),
		UpsTransferOffBatteryDate: reader.GetTimeAt("XOFFBATT", def.UpsTransferOnBatteryDate),

		UpsOnBatterySeconds:           reader.GetDurationSecondsAt("TONBATT", def.UpsOnBatterySeconds),
		UpsOnBatterySecondsCumulative: reader.GetDurationSecondsAt("CUMONBATT", def.UpsOnBatterySecondsCumulative),

		UpsTurnOffDelaySeconds: reader.GetDurationSecondsAt("DSHUTD", def.UpsTurnOffDelaySeconds),
		UpsTurnOnDelaySeconds:  reader.GetDurationSecondsAt("DWAKE", def.UpsTurnOnDelaySeconds),
		UpsTurnOnBatteryMin:    reader.GetFloatAt("RETPCT", def.UpsTurnOnBatteryMin),

		UpsTempInternal: reader.GetFloatAt("ITEMP", def.UpsTempInternal),
		UpsTempAmbient:  reader.GetFloatAt("AMBTEMP", def.UpsTempAmbient),
		UpsHumidity:     reader.GetFloatAt("HUMIDITY", def.UpsHumidity),

		UpsAlarmMode: AlarmMode{
			Text: reader.GetValueAt("ALARMDEL", def.UpsAlarmMode.Text),
			Type: reader.GetMappedAt("ALARMDEL", AlarmModeTypes,
				def.UpsAlarmMode.Type,
			).(AlarmModeType),
		},

		UpsSelftestResult: SelftestResult{
			Text: reader.GetValueAt("SELFTEST", def.UpsSelftestResult.Text),
			Type: reader.GetMappedAt("SELFTEST", SelftestResultTypes,
				def.UpsSelftestResult.Type,
			).(SelftestResultType),
		},
		UpsSelftestIntervalSeconds: reader.GetDurationSecondsAt("STESTI", def.UpsSelftestIntervalSeconds),

		UpsCable: Cable{
			Text: reader.GetValueAt("CABLE", def.UpsCable.Text),
			Type: reader.GetMappedAt("CABLE", CableTypes, def.UpsCable.Type).(CableType),
		},
		UpsDriver: Driver{
			Text: reader.GetValueAt("DRIVER", def.UpsDriver.Text),
			Type: reader.GetMappedAt("DRIVER", DriverTypes, def.UpsDriver.Type).(DriverType),
		},
		UpsMode: Mode{
			Text: reader.GetValueAt("UPSMODE", def.UpsMode.Text),
			Type: reader.GetMappedAt("UPSMODE", ModeTypes, def.UpsMode.Type).(ModeType),
		},

		// shutdown
		ShutdownBatteryMin:          reader.GetFloatAt("MBATTCHG", def.ShutdownBatteryMin),
		ShutdownTimeleftSecondsMin:  reader.GetDurationSecondsAt("MINTIMEL", def.ShutdownTimeleftSecondsMin),
		ShutdownOnBatterySecondsMax: reader.GetDurationSecondsAt("MAXTIME", def.ShutdownOnBatterySecondsMax),

		// apcupsd
		ApcupsdHost:      reader.GetValueAt("HOSTNAME", def.ApcupsdHost),
		ApcupsdVersion:   reader.GetValueAt("VERSION", def.ApcupsdVersion),
		ApcupsdStartTime: reader.GetTimeAt("STARTTIME", def.ApcupsdStartTime),
	}
}
