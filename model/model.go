package model

// State - parsed 'apcaccess status' result
type State struct {
	// input
	InputSensivity           InputSensivity
	InputFrequency           float64
	InputVoltage             float64
	InputVoltageMin          float64
	InputVoltageMax          float64
	InputVoltageNominal      float64
	InputVoltageTransferLow  float64
	InputVoltageTransferHigh float64

	// output
	OutputLoad           float64
	OutputAmps           float64
	OutputPowerNominal   float64
	OutputVoltage        float64
	OutputVoltageNominal float64

	// battery
	BatteryCharge            float64
	BatteryVoltage           float64
	BatteryVoltageNominal    float64
	BatteryExternalCount     uint16
	BatteryBadCount          uint16
	BatteryReplacedTimestamp int64

	// ups
	UpsManafacturedTimestamp       int64
	UpsStatus                      map[string]UpsStatus
	UpsStatusActive                map[string]uint8 // 0 or 1
	UpsStatusInactive              map[string]uint8 // 0 or 1
	UpsStatusChangeCounts          map[string]uint64
	UpsStatusText                  string
	UpsDipSwitchFlag               uint64
	UpsReg1                        uint64
	UpsReg2                        uint64
	UpsReg3                        uint64
	UpsTimeleft                    uint64
	UpsTimeleftLowBattery          uint64
	UpsTransferOnBatteryCount      uint64
	UpsTransferOnBatteryReason     UpsTransferOnbatteryReason
	UpsTransferOnBatteryReasonText string
	UpsTransferOnBatteryTimestamp  int64
	UpsTransferOffBatteryTimestamp int64
	UpsOnBatterySeconds            uint64
	UpsOnBatterySecondsCumulative  uint64
	UpsTurnOffDelaySeconds         uint64
	UpsTurnOnDelaySeconds          uint64
	UpsTurnOnBatteryMin            float64
	UpsTempInternal                float64
	UpsTempAmbient                 float64
	UpsTempHumidity                float64
	UpsAlarmMode                   UpsAlarmMode
	UpsSelftestResult              UpsSelftestResult
	UpsSelftestIntervalSeconds     uint64
	UpsCable                       UpsCable
	UpsDriver                      UpsDriver
	UpsMode                        UpsMode

	// shutdown
	ShutdownBatteryMin          float64
	ShutdownTimeleftSecondsMin  uint64
	ShutdownOnBatterySecondsMax uint64
}

// InputSensivity type
type InputSensivity uint8

// InputSensivities type enum
var InputSensivities = map[string]InputSensivity{
	"Low":         1,
	"Medium":      2,
	"High":        3,
	"Auto Adjust": 4,
	"Unknown":     5,
}

// UpsTransferOnbatteryReason type
type UpsTransferOnbatteryReason uint8

// UpsTransferOnbatteryReasons enum
var UpsTransferOnbatteryReasons = map[string]UpsTransferOnbatteryReason{
	"No transfers since turnon":         1,
	"Automatic or explicit self test":   2,
	"Forced by software":                3,
	"Low line voltage":                  4,
	"High line voltage":                 5,
	"Unacceptable line voltage changes": 6,
	"Line voltage notch or spike":       7,
	"Input frequency out of range":      8,
	"UNKNOWN EVENT":                     9,
}

// UpsAlarmMode type
type UpsAlarmMode uint8

// UpsAlarmModes type enum
var UpsAlarmModes = map[string]UpsAlarmMode{
	"No alarm":    1,
	"Always":      2,
	"5 Seconds":   3,
	"5":           3,
	"30 Seconds":  4,
	"30":          4,
	"Low Battery": 5,
}

// UpsSelftestResult type
type UpsSelftestResult uint8

// UpsSelftestResults type enum
var UpsSelftestResults = map[string]UpsSelftestResult{
	"NO": 1,
	"OK": 2,
	"BT": 3,
	"NG": 4,
	"IP": 5,
	"WN": 6,
	"??": 7,
}

// UpsCable type
type UpsCable uint8

// UpsCables type enum
var UpsCables = map[string]UpsCable{
	"Custom Cable Simple":  1,
	"APC Cable 940-0119A":  2,
	"APC Cable 940-0127A":  3,
	"APC Cable 940-0128A":  4,
	"APC Cable 940-0020B":  5,
	"APC Cable 940-0020C":  6,
	"APC Cable 940-0023A":  7,
	"MAM Cable 04-02-2000": 8,
	"APC Cable 940-0095A":  9,
	"APC Cable 940-0095B":  10,
	"APC Cable 940-0095C":  11,
	"Custom Cable Smart":   12,
	"APC Cable 940-0024B":  121,
	"APC Cable 940-0024C":  122,
	"APC Cable 940-1524C":  123,
	"APC Cable 940-0024G":  124,
	"APC Cable 940-0625A":  125,
	"Ethernet Link":        13,
	"USB Cable":            14,
}

// UpsDriver type
type UpsDriver uint8

// UpsDrivers type enum
var UpsDrivers = map[string]UpsDriver{
	"DUMB UPS Driver":     1,
	"APC Smart UPS (any)": 2,
	"USB UPS Driver":      3,
	"NETWORK UPS Driver":  4,
	"TEST UPS Driver":     5,
	"PCNET UPS Driver":    6,
	"SNMP UPS Driver":     7,
	"MODBUS UPS Driver":   8,
}

// UpsMode type
type UpsMode uint8

// UpsModes type enum
var UpsModes = map[string]UpsMode{
	"Stand Alone":     1,
	"ShareUPS Slave":  2,
	"ShareUPS Master": 3,
}

// UpsStatus type
type UpsStatus uint64

// UpsStatusFlags enum
var UpsStatusFlags = map[string]UpsStatus{
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
