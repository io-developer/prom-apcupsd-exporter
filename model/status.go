package model

// StatusFlags ..
var StatusFlags = map[string]uint64{
	// bit values for APC UPS Status Byte (ups->Status)
	"calibration": 0x00000001,
	"trim":        0x00000002,
	"boost":       0x00000004,
	"online":      0x00000008,
	"onbatt":      0x00000010,
	"overload":    0x00000020,
	"battlow":     0x00000040,
	"replacebatt": 0x00000080,

	// Extended bit values added by apcupsd
	"commlost":    0x00000100, // Communications with UPS lost
	"shutdown":    0x00000200, // Shutdown in progress
	"slave":       0x00000400, // Set if this is a slave
	"slavedown":   0x00000800, // Slave not responding
	"onbatt_msg":  0x00020000, // Set when UPS_ONBATT message is sent
	"fastpoll":    0x00040000, // Set on power failure to poll faster
	"shut_load":   0x00080000, // Set when BatLoad <= percent
	"shut_btime":  0x00100000, // Set when time on batts > maxtime
	"shut_ltime":  0x00200000, // Set when TimeLeft <= runtime
	"shut_emerg":  0x00400000, // Set when battery power has failed
	"shut_remote": 0x00800000, // Set when remote shutdown
	"plugged":     0x01000000, // Set if computer is plugged into UPS
	"battpresent": 0x04000000, // Indicates if battery is connected
}

// Status ..
type Status struct {
	Text  string
	Flags map[string]uint64
}

//func (from *Status) GetChangedFlags(to *Status) []string {
//
//}
