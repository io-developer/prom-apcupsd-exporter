package apcupsd

type StatusState struct {
	Flag         StatusFlag
	Flags        StatusFlagState
	FlagCounters StatusFlagCounters
	Reg1         uint64
	Reg2         uint64
	Reg3         uint64
}

type StatusFlag uint64

const (
	STATUS_FLAG__NA                       = StatusFlag(0x00000000)
	STATUS_FLAG__CALIBRATION              = StatusFlag(0x00000001)
	STATUS_FLAG__TRIM                     = StatusFlag(0x00000002)
	STATUS_FLAG__BOOST                    = StatusFlag(0x00000004)
	STATUS_FLAG__ONLINE                   = StatusFlag(0x00000008)
	STATUS_FLAG__ON_BATTERY               = StatusFlag(0x00000010)
	STATUS_FLAG__OVERLOAD                 = StatusFlag(0x00000020)
	STATUS_FLAG__BATTERY_LOW              = StatusFlag(0x00000040)
	STATUS_FLAG__REPLACE_BATTERY          = StatusFlag(0x00000080)
	STATUS_FLAG__COMMUNICATION_LOST       = StatusFlag(0x00000100)
	STATUS_FLAG__SHUTDOWN_IN_PROGRESS     = StatusFlag(0x00000200)
	STATUS_FLAG__SLAVE                    = StatusFlag(0x00000400)
	STATUS_FLAG__SLAVE_DOWN               = StatusFlag(0x00000800)
	STATUS_FLAG__ON_BATTERY_MESSAGE       = StatusFlag(0x00020000)
	STATUS_FLAG__FASTPOLL                 = StatusFlag(0x00040000)
	STATUS_FLAG__SHUTDOWN_LOAD_LEVEL      = StatusFlag(0x00080000)
	STATUS_FLAG__SHUTDOWN_ON_BATTERY_TIME = StatusFlag(0x00100000)
	STATUS_FLAG__SHUTDOWN_TIMELEFT        = StatusFlag(0x00200000)
	STATUS_FLAG__SHUTDOWN_EMERGENCY       = StatusFlag(0x00400000)
	STATUS_FLAG__SHUTDOWN_REMOTE          = StatusFlag(0x00800000)
	STATUS_FLAG__PLUGGED                  = StatusFlag(0x01000000)
	STATUS_FLAG__BATTERYPRECENT           = StatusFlag(0x04000000)
)

type StatusFlagState struct {
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

type StatusFlagCounters struct {
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
