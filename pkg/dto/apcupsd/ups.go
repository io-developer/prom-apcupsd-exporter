package apcupsd

import (
	"time"
)

type Ups struct {
	Information     UpsInformation
	Config          UpsConfig
	LineIn          UpsLineIn
	LineOut         UpsLineOut
	Battery         UpsBattery
	BatteryTransfer UpsBatteryTransfer
	Selftest        UpsSelftest
	Sensors         UpsSensors
	Status          UpsStatus
}

type UpsInformation struct {
	Model            string
	Serial           string
	Firmware         string
	ManafacturedDate time.Time
	Version          string
	StartTime        time.Time
}

type UpsConfig struct {
	Name                    string
	Host                    string
	Mode                    Mode
	AlarmMode               AlarmMode
	Cable                   Cable
	Driver                  Driver
	DipSwitch               uint64
	ShutdownDelay           time.Duration
	ShutdownBatteryLevel    float64
	ShutdownBatteryTime     time.Duration
	ShutdownBatteryTimeLeft time.Duration
	WakeupDelay             time.Duration
	WakeupBatteryLevel      float64
	LowBatteryTimeLeft      time.Duration
}

type UpsLineIn struct {
	Sensivity      Sensivity
	Frequency      float64
	Voltage        float64
	VoltageMin     float64
	VoltageMax     float64
	VoltageNominal float64
}

type UpsLineOut struct {
	LoadLevel            float64
	Voltage              float64
	VoltageNominal       float64
	Amps                 float64
	PowerNominal         float64
	PowerApparentNominal float64
}

type UpsBattery struct {
	ChargeLevel    float64
	Voltage        float64
	VoltageNominal float64
	Timeeft        time.Duration
	ExternalCount  uint16
	BadCount       uint16
	ReplacedDate   time.Time
}

type UpsBatteryTransfer struct {
	Time               time.Duration
	TimeCumulative     time.Duration
	Count              uint64
	LastReason         BatteryTransferReason
	LastDateOn         time.Time
	LastDateOff        time.Time
	LineInVoltageUnder float64
	LineInVoltageOver  float64
}

type UpsSelftest struct {
	LastResult SelftestResult
	Interval   time.Duration
}

type UpsSensors struct {
	TempInternal float64
	TempAmbient  float64
	Humidity     float64
}

type UpsStatus struct {
	Flag      StatusFlag
	FlagState UpsStatusFlagState
	Reg1      uint64
	Reg2      uint64
	Reg3      uint64
}

type UpsStatusFlagState struct {
	Calibration             bool
	Trim                    bool
	Boost                   bool
	Online                  bool
	OnBattery               bool
	Overload                bool
	BatteryLow              bool
	ReplaceBattery          bool
	CommunicationLost       bool
	Shutdown                bool
	Slave                   bool
	SlaveDown               bool
	OnBatteryMessageSent    bool
	FastPoll                bool
	ShutdownBatteryLevel    bool
	ShutdownBatteryTime     bool
	ShutdownBatteryTimeLeft bool
	ShutdownEmergency       bool
	ShutdownRemote          bool
	Plugged                 bool
	BatteryPresent          bool
}
