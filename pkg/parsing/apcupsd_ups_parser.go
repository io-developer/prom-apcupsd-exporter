package parsing

import (
	"time"

	"github.com/io-developer/prom-apcupsd-exporter/pkg/dto/apcupsd"
)

type ApcupsdUpsParser struct {
	reader *ApcaccessResponseReader
}

func NewApcupsdUpsParser(reader *ApcaccessResponseReader) *ApcupsdUpsParser {
	return &ApcupsdUpsParser{
		reader: reader,
	}
}

func (p *ApcupsdUpsParser) Parse() (apcupsd.Ups, error) {
	ups := apcupsd.Ups{
		HardwareInfo:    p.parseHardwareInfo(),
		Config:          p.parseConfig(),
		LineIn:          p.parseLineIn(),
		LineOut:         p.parseLineOut(),
		Battery:         p.parseBattery(),
		BatteryTransfer: p.parseBatteryTransfer(),
		Selftest:        p.parseSelftest(),
		Sensors:         p.parseSensors(),
		Status:          p.parseStatus(),
	}
	return ups, nil
}

func (p *ApcupsdUpsParser) parseHardwareInfo() apcupsd.UpsHardwareInfo {
	return apcupsd.UpsHardwareInfo{
		Model:            p.reader.GetValueAt("MODEL", ""),
		Serial:           p.reader.GetValueAt("SERIALNO", ""),
		Firmware:         p.reader.GetValueAt("FIRMWARE", ""),
		ManafacturedDate: p.reader.GetTimeAt("MANDATE", time.UnixMilli(0)),
	}
}

func (p *ApcupsdUpsParser) parseConfig() apcupsd.UpsConfig {
	return apcupsd.UpsConfig{
		Name:                    p.reader.GetValueAt("UPSNAME", ""),
		Host:                    p.reader.GetValueAt("HOSTNAME", ""),
		Mode:                    p.parseMode(),
		AlarmMode:               p.parseAlarmMode(),
		Cable:                   p.parseCable(),
		Driver:                  p.parseDriver(),
		DipSwitch:               p.reader.GetUintAt("DIPSW", 0),
		ShutdownDelay:           p.reader.GetDurationAt("DSHUTD", 0),
		ShutdownBatteryLevel:    p.reader.GetFloatAt("MBATTCHG", 0),
		ShutdownBatteryTime:     p.reader.GetDurationAt("MAXTIME", 0),
		ShutdownBatteryTimeLeft: p.reader.GetDurationAt("MINTIMEL", 0),
		WakeupDelay:             p.reader.GetDurationAt("DWAKE", 0),
		WakeupBatteryLevel:      p.reader.GetFloatAt("RETPCT", 0),
		StartTime:               p.reader.GetTimeAt("STARTTIME", time.UnixMilli(0)),
		Version:                 p.reader.GetValueAt("VERSION", ""),
	}
}

func (p *ApcupsdUpsParser) parseLineIn() apcupsd.UpsLineIn {
	return apcupsd.UpsLineIn{
		Sensivity:      p.parseSensivity(),
		Frequency:      p.reader.GetFloatAt("LINEFREQ", 0),
		Voltage:        p.reader.GetFloatAt("LINEV", 0),
		VoltageMin:     p.reader.GetFloatAt("MINLINEV", 0),
		VoltageMax:     p.reader.GetFloatAt("MAXLINEV", 0),
		VoltageNominal: p.reader.GetFloatAt("NOMINV", 0),
	}
}

func (p *ApcupsdUpsParser) parseLineOut() apcupsd.UpsLineOut {
	return apcupsd.UpsLineOut{
		LoadLevel:            p.reader.GetFloatAt("LOADPCT", 0),
		Voltage:              p.reader.GetFloatAt("OUTPUTV", 0),
		VoltageNominal:       p.reader.GetFloatAt("NOMOUTV", 0),
		Amps:                 p.reader.GetFloatAt("OUTCURNT", 0),
		PowerNominal:         p.reader.GetFloatAt("NOMPOWER", 0),
		PowerApparentNominal: p.reader.GetFloatAt("NOMAPNT", 0),
	}
}

func (p *ApcupsdUpsParser) parseBattery() apcupsd.UpsBattery {
	return apcupsd.UpsBattery{
		ChargeLevel:        p.reader.GetFloatAt("BCHARGE", 0),
		Voltage:            p.reader.GetFloatAt("BATTV", 0),
		VoltageNominal:     p.reader.GetFloatAt("NOMBATTV", 0),
		TimeLeft:           p.reader.GetDurationAt("TIMELEFT", 0),
		LowLevelAtTimeLeft: p.reader.GetDurationAt("DLOWBATT", 0),
		ExternalCount:      uint16(p.reader.GetUintAt("EXTBATTS", 0)),
		BadCount:           uint16(p.reader.GetUintAt("BADBATTS", 0)),
		ReplacedDate:       p.reader.GetTimeAt("BATTDATE", time.UnixMilli(0)),
	}
}

func (p *ApcupsdUpsParser) parseBatteryTransfer() apcupsd.UpsBatteryTransfer {
	return apcupsd.UpsBatteryTransfer{
		Time:               p.reader.GetDurationAt("TONBATT", 0),
		TimeCumulative:     p.reader.GetDurationAt("CUMONBATT", 0),
		Count:              p.reader.GetUintAt("NUMXFERS", 0),
		LastReason:         p.parseBatteryTransferReason(),
		LastDateOn:         p.reader.GetTimeAt("XONBATT", time.UnixMilli(0)),
		LastDateOff:        p.reader.GetTimeAt("XOFFBATT", time.UnixMilli(0)),
		LineInVoltageUnder: p.reader.GetFloatAt("LOTRANS", 0),
		LineInVoltageOver:  p.reader.GetFloatAt("HITRANS", 0),
	}
}

func (p *ApcupsdUpsParser) parseSelftest() apcupsd.UpsSelftest {
	return apcupsd.UpsSelftest{
		LastResult: p.parseSelftestResult(),
		Interval:   p.reader.GetDurationAt("STESTI", 0),
	}
}

func (p *ApcupsdUpsParser) parseSensors() apcupsd.UpsSensors {
	return apcupsd.UpsSensors{
		TempInternal: p.reader.GetFloatAt("ITEMP", 0),
		TempAmbient:  p.reader.GetFloatAt("AMBTEMP", 0),
		Humidity:     p.reader.GetFloatAt("HUMIDITY", 0),
	}
}

func (p *ApcupsdUpsParser) parseStatus() apcupsd.UpsStatus {
	flag := apcupsd.StatusFlag(p.reader.GetUintAt("STATFLAG", 0))
	return apcupsd.UpsStatus{
		Text:      p.reader.GetValueAt("STATUS", ""),
		Flag:      flag,
		FlagState: p.parseStatusFlagState(flag),
		Reg1:      p.reader.GetUintAt("REG1", 0),
		Reg2:      p.reader.GetUintAt("REG2", 0),
		Reg3:      p.reader.GetUintAt("REG3", 0),
	}
}

func (p *ApcupsdUpsParser) parseStatusFlagState(flag apcupsd.StatusFlag) apcupsd.UpsStatusFlagState {
	return apcupsd.UpsStatusFlagState{
		Calibration:             flag&apcupsd.STATUS_FLAG__CALIBRATION != 0,
		Trim:                    flag&apcupsd.STATUS_FLAG__TRIM != 0,
		Boost:                   flag&apcupsd.STATUS_FLAG__BOOST != 0,
		Online:                  flag&apcupsd.STATUS_FLAG__ONLINE != 0,
		OnBattery:               flag&apcupsd.STATUS_FLAG__ON_BATTERY != 0,
		Overload:                flag&apcupsd.STATUS_FLAG__OVERLOAD != 0,
		BatteryLow:              flag&apcupsd.STATUS_FLAG__BATTERY_LOW != 0,
		ReplaceBattery:          flag&apcupsd.STATUS_FLAG__REPLACE_BATTERY != 0,
		CommunicationLost:       flag&apcupsd.STATUS_FLAG__COMMUNICATION_LOST != 0,
		Shutdown:                flag&apcupsd.STATUS_FLAG__SHUTDOWN != 0,
		Slave:                   flag&apcupsd.STATUS_FLAG__SLAVE != 0,
		SlaveDown:               flag&apcupsd.STATUS_FLAG__SLAVE_DOWN != 0,
		OnBatteryMessageSent:    flag&apcupsd.STATUS_FLAG__ON_BATTERY_MESSAGE_SENT != 0,
		FastPoll:                flag&apcupsd.STATUS_FLAG__FAST_POLL != 0,
		ShutdownBatteryLevel:    flag&apcupsd.STATUS_FLAG__SHUTDOWN_BATTERY_LEVEL != 0,
		ShutdownBatteryTime:     flag&apcupsd.STATUS_FLAG__SHUTDOWN_BATTERY_TIME != 0,
		ShutdownBatteryTimeLeft: flag&apcupsd.STATUS_FLAG__SHUTDOWN_BATTERY_TIME_LEFT != 0,
		ShutdownEmergency:       flag&apcupsd.STATUS_FLAG__SHUTDOWN_EMERGENCY != 0,
		ShutdownRemote:          flag&apcupsd.STATUS_FLAG__SHUTDOWN_REMOTE != 0,
		Plugged:                 flag&apcupsd.STATUS_FLAG__PLUGGED != 0,
		BatteryPresent:          flag&apcupsd.STATUS_FLAG__BATTERY_PRESENT != 0,
	}
}

func (p *ApcupsdUpsParser) parseSensivity() apcupsd.Sensivity {
	val := p.reader.GetValueAt("SENSE", "")
	res, exists := apcupsd.SENSIVITY_FROM_NAME[val]
	if exists {
		return res
	}
	return apcupsd.SENSIVITY__NA
}

func (p *ApcupsdUpsParser) parseCable() apcupsd.Cable {
	val := p.reader.GetValueAt("CABLE", "")
	res, exists := apcupsd.CABLE_FROM_NAME[val]
	if exists {
		return res
	}
	return apcupsd.CABLE__NO_CABLE
}

func (p *ApcupsdUpsParser) parseDriver() apcupsd.Driver {
	val := p.reader.GetValueAt("DRIVER", "")
	res, exists := apcupsd.DRIVER_FROM_NAME[val]
	if exists {
		return res
	}
	return apcupsd.DRIVER__NA
}

func (p *ApcupsdUpsParser) parseMode() apcupsd.Mode {
	val := p.reader.GetValueAt("UPSMODE", "")
	res, exists := apcupsd.MODE_FROM_NAME[val]
	if exists {
		return res
	}
	return apcupsd.MODE__NA
}

func (p *ApcupsdUpsParser) parseAlarmMode() apcupsd.AlarmMode {
	val := p.reader.GetValueAt("ALARMDEL", "")
	res, exists := apcupsd.ALARM_MODE_FROM_NAME[val]
	if exists {
		return res
	}
	return apcupsd.ALARM_MODE__NA
}

func (p *ApcupsdUpsParser) parseSelftestResult() apcupsd.SelftestResult {
	val := p.reader.GetValueAt("SELFTEST", "")
	res, exists := apcupsd.SELFTEST_RESULT_FROM_RAW[val]
	if exists {
		return res
	}
	return apcupsd.SELFTEST_RESULT__NA
}

func (p *ApcupsdUpsParser) parseBatteryTransferReason() apcupsd.BatteryTransferReason {
	val := p.reader.GetValueAt("LASTXFER", "")
	res, exists := apcupsd.BATTERY_TRANSFER_REASON_FROM_NAME[val]
	if exists {
		return res
	}
	return apcupsd.BATTERY_TRANSFER_REASON__NA
}
