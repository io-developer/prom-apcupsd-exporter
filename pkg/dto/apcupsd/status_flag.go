package apcupsd

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
