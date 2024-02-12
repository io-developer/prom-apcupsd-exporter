package apcupsd

import (
	"time"
)

type State struct {
	Information     Information
	Config          Config
	Input           InputState
	Output          OutputState
	Battery         BatteryState
	BatteryTransfer BatteryTransferState
	Sensors         SensorsState
	Status          StatusState
	SelftestResult  SelftestResult
}

type Information struct {
	Model            string
	Serial           string
	Firmware         string
	ManafacturedDate time.Time
	Version          string
	StartTime        time.Time
}

type Config struct {
	Name                            string
	Host                            string
	Mode                            Mode
	AlarmMode                       AlarmMode
	Cable                           Cable
	Driver                          Driver
	Sensivity                       Sensivity
	DipSwitch                       uint64
	SelftestInterval                time.Duration
	ShutdownDelay                   time.Duration
	ShutdownBatteryLevel            float64
	ShutdownBatteryTime             time.Duration
	ShutdownTimeLeft                time.Duration
	WakeupDelay                     time.Duration
	WakeupBatteryLevel              float64
	LowBatteryTimeLeft              time.Duration
	BatteryTransferInputVoltageLow  float64
	BatteryTransferInputVoltageHigh float64
}

type InputState struct {
	Frequency      float64
	Voltage        float64
	VoltageMin     float64
	VoltageMax     float64
	VoltageNominal float64
}

type OutputState struct {
	Load                 float64
	PowerNominal         float64
	PowerApparentNominal float64
	Voltage              float64
	VoltageNominal       float64
	Amps                 float64
}

type BatteryState struct {
	Charge         float64
	Voltage        float64
	VoltageNominal float64
	Timeleft       time.Duration
	ExternalCount  uint16
	BadCount       uint16
	ReplacedDate   time.Time
}

type BatteryTransferState struct {
	Time           time.Duration
	TimeCumulative time.Duration
	Count          uint64
	LastReason     BatteryTransferReason
	LastDateOn     time.Time
	LastDateOff    time.Time
}

type SensorsState struct {
	TempInternal float64
	TempAmbient  float64
	Humidity     float64
}
