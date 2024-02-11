package dto

import (
	"time"
)

type Ups struct {
	Information UpsInformation
	Settings    UpsSettings
	State       UpsState
}

type UpsInformation struct {
	Name             string
	Model            string
	Serial           string
	Firmware         string
	ManafacturedDate time.Time
	ApcupsdHost      string
	ApcupsdVersion   string
	ApcupsdStartTime time.Time
}

type UpsSettings struct {
	Mode                            UpsMode
	AlarmMode                       UpsAlarmMode
	Cable                           UpsCable
	Driver                          UpsDriver
	Sensivity                       UpsSensivity
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

type UpsState struct {
	Line            UpsLineState
	Output          UpsOutputState
	Battery         UpsBatteryState
	BatteryTransfer UpsBatteryTransferState
	Sensors         UpsSensorsState
	Status          UpsStatusState
	SelftestResult  UpsSelftestResult
}

type UpsLineState struct {
	Frequency      float64
	Voltage        float64
	VoltageMin     float64
	VoltageMax     float64
	VoltageNominal float64
}

type UpsOutputState struct {
	LoadLevel            float64
	PowerNominal         float64
	PowerApparentNominal float64
	Voltage              float64
	VoltageNominal       float64
	Amps                 float64
}

type UpsBatteryState struct {
	ChargeLevel    float64
	Voltage        float64
	VoltageNominal float64
	Timeleft       time.Duration
	ExternalCount  uint16
	BadCount       uint16
	ReplacedDate   time.Time
}

type UpsSensorsState struct {
	TempInternal float64
	TempAmbient  float64
	Humidity     float64
}

type UpsBatteryTransferState struct {
	OnBatteryTime           time.Duration
	OnBatteryTimeCumulative time.Duration
	OnBatteryCount          uint64
	OnBatteryLastDate       time.Time
	OnBatteryLastReason     TransferOnbatteryReason
	OffBatteryLastDate      time.Time
}

type UpsSensivity uint8

const (
	UPS_SENSIVITY_LOW     = UpsSensivity(1)
	UPS_SENSIVITY_MEDIUM  = UpsSensivity(2)
	UPS_SENSIVITY_HIGH    = UpsSensivity(3)
	UPS_SENSIVITY_AUTO    = UpsSensivity(4)
	UPS_SENSIVITY_UNKNOWN = UpsSensivity(5)
)

type TransferOnbatteryReason uint8

const (
	UPS_BATTERY_TRANSFER_REASON__NA         = TransferOnbatteryReason(0)
	UPS_BATTERY_TRANSFER_REASON__NONE       = TransferOnbatteryReason(1)
	UPS_BATTERY_TRANSFER_REASON__SELFTEST   = TransferOnbatteryReason(2)
	UPS_BATTERY_TRANSFER_REASON__FORCED     = TransferOnbatteryReason(3)
	UPS_BATTERY_TRANSFER_REASON__UNDERVOLT  = TransferOnbatteryReason(4)
	UPS_BATTERY_TRANSFER_REASON__OVERVOLT   = TransferOnbatteryReason(5)
	UPS_BATTERY_TRANSFER_REASON__RIPPLE     = TransferOnbatteryReason(6)
	UPS_BATTERY_TRANSFER_REASON__NOTCHSPIKE = TransferOnbatteryReason(7)
	UPS_BATTERY_TRANSFER_REASON__FREQ       = TransferOnbatteryReason(8)
	UPS_BATTERY_TRANSFER_REASON__UNKNOWN    = TransferOnbatteryReason(9)
)

type UpsAlarmMode uint8

const (
	UPS_ALARM_MODE__NA          = UpsAlarmMode(0)
	UPS_ALARM_MODE__NONE        = UpsAlarmMode(1)
	UPS_ALARM_MODE__ALWAYS      = UpsAlarmMode(2)
	UPS_ALARM_MODE__5_SEC       = UpsAlarmMode(3)
	UPS_ALARM_MODE__30_SEC      = UpsAlarmMode(4)
	UPS_ALARM_MODE__LOW_BATTERY = UpsAlarmMode(5)
)

type UpsSelftestResult uint8

const (
	UPS_SELFTEST_RESULT__NA         = UpsSelftestResult(0)
	UPS_SELFTEST_RESULT__NONE       = UpsSelftestResult(1)
	UPS_SELFTEST_RESULT__PASSED     = UpsSelftestResult(2)
	UPS_SELFTEST_RESULT__FAILCAP    = UpsSelftestResult(3)
	UPS_SELFTEST_RESULT__FAILED     = UpsSelftestResult(4)
	UPS_SELFTEST_RESULT__INPROGRESS = UpsSelftestResult(5)
	UPS_SELFTEST_RESULT__WARNING    = UpsSelftestResult(6)
	UPS_SELFTEST_RESULT__UNKNOWN    = UpsSelftestResult(7)
)

type UpsCable uint8

const (
	UPS_CABLE__NO_CABLE      = UpsCable(0)
	UPS_CABLE__CUSTOM_SIMPLE = UpsCable(1)
	UPS_CABLE__APC_940_0119A = UpsCable(2)
	UPS_CABLE__APC_940_0127A = UpsCable(3)
	UPS_CABLE__APC_940_0128A = UpsCable(4)
	UPS_CABLE__APC_940_0020B = UpsCable(5)
	UPS_CABLE__APC_940_0020C = UpsCable(6)
	UPS_CABLE__APC_940_0023A = UpsCable(7)
	UPS_CABLE__MAM           = UpsCable(8)
	UPS_CABLE__APC_940_0095A = UpsCable(9)
	UPS_CABLE__APC_940_0095B = UpsCable(10)
	UPS_CABLE__APC_940_0095C = UpsCable(11)
	UPS_CABLE__CUSTOM_SMART  = UpsCable(12)
	UPS_CABLE__APC_940_0024B = UpsCable(121)
	UPS_CABLE__APC_940_0024C = UpsCable(122)
	UPS_CABLE__APC_940_1524C = UpsCable(123)
	UPS_CABLE__APC_940_0024G = UpsCable(124)
	UPS_CABLE__APC_940_0625A = UpsCable(125)
	UPS_CABLE__ETHERNET      = UpsCable(13)
	UPS_CABLE__USB           = UpsCable(14)
)

type UpsDriver uint8

const (
	UPS_DRIVER__NO       = UpsDriver(0)
	UPS_DRIVER__DUMB     = UpsDriver(1)
	UPS_DRIVER__APCSMART = UpsDriver(2)
	UPS_DRIVER__USB      = UpsDriver(3)
	UPS_DRIVER__NETWORK  = UpsDriver(4)
	UPS_DRIVER__TEST     = UpsDriver(5)
	UPS_DRIVER__PCNET    = UpsDriver(6)
	UPS_DRIVER__SNMPLITE = UpsDriver(7)
	UPS_DRIVER__MODBUS   = UpsDriver(8)
)

type UpsMode uint8

const (
	UPS_MODE__INVALID      = UpsMode(0)
	UPS_MODE__STAND_ALONE  = UpsMode(1)
	UPS_MODE__SHARE_SLAVE  = UpsMode(2)
	UPS_MODE__SHARE_MASTER = UpsMode(3)
)

type UpsStatusState struct {
	Flag         UpsStatusFlag
	Flags        UpsStatusFlagState
	FlagCounters UpsStatusFlagCounters
	Reg1         uint64
	Reg2         uint64
	Reg3         uint64
}

type UpsStatusFlag uint64

const (
	UPS_STATUS_FLAG__NA                       = UpsStatusFlag(0x00000000)
	UPS_STATUS_FLAG__CALIBRATION              = UpsStatusFlag(0x00000001)
	UPS_STATUS_FLAG__TRIM                     = UpsStatusFlag(0x00000002)
	UPS_STATUS_FLAG__BOOST                    = UpsStatusFlag(0x00000004)
	UPS_STATUS_FLAG__ONLINE                   = UpsStatusFlag(0x00000008)
	UPS_STATUS_FLAG__ON_BATTERY               = UpsStatusFlag(0x00000010)
	UPS_STATUS_FLAG__OVERLOAD                 = UpsStatusFlag(0x00000020)
	UPS_STATUS_FLAG__BATTERY_LOW              = UpsStatusFlag(0x00000040)
	UPS_STATUS_FLAG__REPLACE_BATTERY          = UpsStatusFlag(0x00000080)
	UPS_STATUS_FLAG__COMMUNICATION_LOST       = UpsStatusFlag(0x00000100)
	UPS_STATUS_FLAG__SHUTDOWN_IN_PROGRESS     = UpsStatusFlag(0x00000200)
	UPS_STATUS_FLAG__SLAVE                    = UpsStatusFlag(0x00000400)
	UPS_STATUS_FLAG__SLAVE_DOWN               = UpsStatusFlag(0x00000800)
	UPS_STATUS_FLAG__ON_BATTERY_MESSAGE       = UpsStatusFlag(0x00020000)
	UPS_STATUS_FLAG__FASTPOLL                 = UpsStatusFlag(0x00040000)
	UPS_STATUS_FLAG__SHUTDOWN_LOAD_LEVEL      = UpsStatusFlag(0x00080000)
	UPS_STATUS_FLAG__SHUTDOWN_ON_BATTERY_TIME = UpsStatusFlag(0x00100000)
	UPS_STATUS_FLAG__SHUTDOWN_TIMELEFT        = UpsStatusFlag(0x00200000)
	UPS_STATUS_FLAG__SHUTDOWN_EMERGENCY       = UpsStatusFlag(0x00400000)
	UPS_STATUS_FLAG__SHUTDOWN_REMOTE          = UpsStatusFlag(0x00800000)
	UPS_STATUS_FLAG__PLUGGED                  = UpsStatusFlag(0x01000000)
	UPS_STATUS_FLAG__BATTERYPRECENT           = UpsStatusFlag(0x04000000)
)

type UpsStatusFlagState struct {
	Calibration          bool
	Trim                 bool
	Boost                bool
	Online               bool
	OnBattery            bool
	Overload             bool
	BatteryLow           bool
	ReplaceBattery       bool
	CommunicationLost    bool
	ShutdownInProgress   bool
	Slave                bool
	SlaveDown            bool
	OnBatteryMessage     bool
	FastPoll             bool
	ShutdownBatteryLevel bool
	ShutdownBatteryTime  bool
	ShutdownTimeLeft     bool
	ShutdownEmergency    bool
	ShutdownRemote       bool
	Plugged              bool
	BatteryPrecent       bool
}

type UpsStatusFlagCounters struct {
	Calibration          uint64
	Trim                 uint64
	Boost                uint64
	Online               uint64
	OnBattery            uint64
	Overload             uint64
	BatteryLow           uint64
	ReplaceBattery       uint64
	CommunicationLost    uint64
	ShutdownInProgress   uint64
	Slave                uint64
	SlaveDown            uint64
	OnBatteryMessage     uint64
	FastPoll             uint64
	ShutdownBatteryLevel uint64
	ShutdownBatteryTime  uint64
	ShutdownTimeLeft     uint64
	ShutdownEmergency    uint64
	ShutdownRemote       uint64
	Plugged              uint64
	BatteryPrecent       uint64
}
