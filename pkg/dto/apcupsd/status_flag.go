package apcupsd

type StatusFlag uint64

const (
	STATUS_FLAG__NA                         = StatusFlag(0x00000000)
	STATUS_FLAG__CALIBRATION                = StatusFlag(0x00000001)
	STATUS_FLAG__TRIM                       = StatusFlag(0x00000002)
	STATUS_FLAG__BOOST                      = StatusFlag(0x00000004)
	STATUS_FLAG__ONLINE                     = StatusFlag(0x00000008)
	STATUS_FLAG__ON_BATTERY                 = StatusFlag(0x00000010)
	STATUS_FLAG__OVERLOAD                   = StatusFlag(0x00000020)
	STATUS_FLAG__BATTERY_LOW                = StatusFlag(0x00000040)
	STATUS_FLAG__REPLACE_BATTERY            = StatusFlag(0x00000080)
	STATUS_FLAG__COMMUNICATION_LOST         = StatusFlag(0x00000100)
	STATUS_FLAG__SHUTDOWN                   = StatusFlag(0x00000200)
	STATUS_FLAG__SLAVE                      = StatusFlag(0x00000400)
	STATUS_FLAG__SLAVE_DOWN                 = StatusFlag(0x00000800)
	STATUS_FLAG__ON_BATTERY_MESSAGE_SENT    = StatusFlag(0x00020000)
	STATUS_FLAG__FAST_POLL                  = StatusFlag(0x00040000)
	STATUS_FLAG__SHUTDOWN_BATTERY_LEVEL     = StatusFlag(0x00080000)
	STATUS_FLAG__SHUTDOWN_BATTERY_TIME      = StatusFlag(0x00100000)
	STATUS_FLAG__SHUTDOWN_BATTERY_TIME_LEFT = StatusFlag(0x00200000)
	STATUS_FLAG__SHUTDOWN_EMERGENCY         = StatusFlag(0x00400000)
	STATUS_FLAG__SHUTDOWN_REMOTE            = StatusFlag(0x00800000)
	STATUS_FLAG__PLUGGED                    = StatusFlag(0x01000000)
	STATUS_FLAG__BATTERY_PRESENT            = StatusFlag(0x04000000)
)

var STATUS_FLAG_TO_NAME = map[StatusFlag]string{
	STATUS_FLAG__NA:                         "NA",
	STATUS_FLAG__CALIBRATION:                "Calibration",
	STATUS_FLAG__TRIM:                       "Smart trim",
	STATUS_FLAG__BOOST:                      "Smart boost",
	STATUS_FLAG__ONLINE:                     "Online",
	STATUS_FLAG__ON_BATTERY:                 "On battery",
	STATUS_FLAG__OVERLOAD:                   "Overloaded",
	STATUS_FLAG__BATTERY_LOW:                "Battery low",
	STATUS_FLAG__REPLACE_BATTERY:            "Replace battery",
	STATUS_FLAG__COMMUNICATION_LOST:         "Communication lost",
	STATUS_FLAG__SHUTDOWN:                   "Shutdown in progress",
	STATUS_FLAG__SLAVE:                      "Slave",
	STATUS_FLAG__SLAVE_DOWN:                 "Slave down",
	STATUS_FLAG__ON_BATTERY_MESSAGE_SENT:    "On battery message sent",
	STATUS_FLAG__FAST_POLL:                  "Fast poll",
	STATUS_FLAG__SHUTDOWN_BATTERY_LEVEL:     "Shutdown due to battery charge left",
	STATUS_FLAG__SHUTDOWN_BATTERY_TIME:      "Shutdown due to on battery time",
	STATUS_FLAG__SHUTDOWN_BATTERY_TIME_LEFT: "Shutdown due to battery time left",
	STATUS_FLAG__SHUTDOWN_EMERGENCY:         "Shutdown due to battery power has failed",
	STATUS_FLAG__SHUTDOWN_REMOTE:            "Shutdown remote",
	STATUS_FLAG__PLUGGED:                    "UPS plugged",
	STATUS_FLAG__BATTERY_PRESENT:            "Battery connected",
}

var STATUS_FLAG_TO_SHORT_NAME = map[StatusFlag]string{
	STATUS_FLAG__NA:                         "NA",
	STATUS_FLAG__CALIBRATION:                "calibration",
	STATUS_FLAG__TRIM:                       "trim",
	STATUS_FLAG__BOOST:                      "boost",
	STATUS_FLAG__ONLINE:                     "online",
	STATUS_FLAG__ON_BATTERY:                 "onbatt",
	STATUS_FLAG__OVERLOAD:                   "overload",
	STATUS_FLAG__BATTERY_LOW:                "battlow",
	STATUS_FLAG__REPLACE_BATTERY:            "replacebatt",
	STATUS_FLAG__COMMUNICATION_LOST:         "commlost",
	STATUS_FLAG__SHUTDOWN:                   "shutdown",
	STATUS_FLAG__SLAVE:                      "slave",
	STATUS_FLAG__SLAVE_DOWN:                 "slavedown",
	STATUS_FLAG__ON_BATTERY_MESSAGE_SENT:    "onbatt_msg",
	STATUS_FLAG__FAST_POLL:                  "fastpoll",
	STATUS_FLAG__SHUTDOWN_BATTERY_LEVEL:     "shut_load",
	STATUS_FLAG__SHUTDOWN_BATTERY_TIME:      "shut_btime",
	STATUS_FLAG__SHUTDOWN_BATTERY_TIME_LEFT: "shut_ltime",
	STATUS_FLAG__SHUTDOWN_EMERGENCY:         "shut_emerg",
	STATUS_FLAG__SHUTDOWN_REMOTE:            "shut_remote",
	STATUS_FLAG__PLUGGED:                    "plugged",
	STATUS_FLAG__BATTERY_PRESENT:            "battpresent",
}

var STATUS_FLAG_FROM_NAME = map[string]StatusFlag{
	"NA":                                  STATUS_FLAG__NA,
	"Calibration":                         STATUS_FLAG__CALIBRATION,
	"Smart trim":                          STATUS_FLAG__TRIM,
	"Smart boost":                         STATUS_FLAG__BOOST,
	"Online":                              STATUS_FLAG__ONLINE,
	"On battery":                          STATUS_FLAG__ON_BATTERY,
	"Overloaded":                          STATUS_FLAG__OVERLOAD,
	"Battery low":                         STATUS_FLAG__BATTERY_LOW,
	"Replace battery":                     STATUS_FLAG__REPLACE_BATTERY,
	"Communication lost":                  STATUS_FLAG__COMMUNICATION_LOST,
	"Shutdown in progress":                STATUS_FLAG__SHUTDOWN,
	"Slave":                               STATUS_FLAG__SLAVE,
	"Slave down":                          STATUS_FLAG__SLAVE_DOWN,
	"On battery message sent":             STATUS_FLAG__ON_BATTERY_MESSAGE_SENT,
	"Fast poll":                           STATUS_FLAG__FAST_POLL,
	"Shutdown due to battery charge left": STATUS_FLAG__SHUTDOWN_BATTERY_LEVEL,
	"Shutdown due to on battery time":     STATUS_FLAG__SHUTDOWN_BATTERY_TIME,
	"Shutdown due to battery time left":   STATUS_FLAG__SHUTDOWN_BATTERY_TIME_LEFT,
	"Shutdown due to battery power has failed": STATUS_FLAG__SHUTDOWN_EMERGENCY,
	"Shutdown remote":                          STATUS_FLAG__SHUTDOWN_REMOTE,
	"Plugged":                                  STATUS_FLAG__PLUGGED,
	"Battery connected":                        STATUS_FLAG__BATTERY_PRESENT,
}

var STATUS_FLAG_FROM_SHORT_NAME = map[string]StatusFlag{
	"NA":          STATUS_FLAG__NA,
	"calibration": STATUS_FLAG__CALIBRATION,
	"trim":        STATUS_FLAG__TRIM,
	"boost":       STATUS_FLAG__BOOST,
	"online":      STATUS_FLAG__ONLINE,
	"onbatt":      STATUS_FLAG__ON_BATTERY,
	"overload":    STATUS_FLAG__OVERLOAD,
	"battlow":     STATUS_FLAG__BATTERY_LOW,
	"replacebatt": STATUS_FLAG__REPLACE_BATTERY,
	"commlost":    STATUS_FLAG__COMMUNICATION_LOST,
	"shutdown":    STATUS_FLAG__SHUTDOWN,
	"slave":       STATUS_FLAG__SLAVE,
	"slavedown":   STATUS_FLAG__SLAVE_DOWN,
	"onbatt_msg":  STATUS_FLAG__ON_BATTERY_MESSAGE_SENT,
	"fastpoll":    STATUS_FLAG__FAST_POLL,
	"shut_load":   STATUS_FLAG__SHUTDOWN_BATTERY_LEVEL,
	"shut_btime":  STATUS_FLAG__SHUTDOWN_BATTERY_TIME,
	"shut_ltime":  STATUS_FLAG__SHUTDOWN_BATTERY_TIME_LEFT,
	"shut_emerg":  STATUS_FLAG__SHUTDOWN_EMERGENCY,
	"shut_remote": STATUS_FLAG__SHUTDOWN_REMOTE,
	"plugged":     STATUS_FLAG__PLUGGED,
	"battpresent": STATUS_FLAG__BATTERY_PRESENT,
}
