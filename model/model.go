package model

// State - parsed 'apcaccess status' result
type State struct {
	// input
	InputSensivity           InputSensivity
	InputSensivityText       string
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
	UpsTimeleftLowBattery          int64
	UpsTransferOnBatteryCount      uint64
	UpsTransferOnBatteryReason     UpsTransferOnbatteryReason
	UpsTransferOnBatteryReasonText string
	UpsTransferOnBatteryTimestamp  int64
	UpsTransferOffBatteryTimestamp int64
	UpsOnBatterySeconds            uint64
	UpsOnBatterySecondsCumulative  uint64
	UpsTurnOffDelaySeconds         int64
	UpsTurnOnDelaySeconds          int64
	UpsTurnOnBatteryMin            float64
	UpsTempInternal                float64
	UpsTempAmbient                 float64
	UpsTempHumidity                float64
	UpsAlarmMode                   UpsAlarmMode
	UpsAlarmModeText               string
	UpsSelftestResult              UpsSelftestResult
	UpsSelftestResultText          string
	UpsSelftestIntervalSeconds     int64
	UpsCable                       UpsCable
	UpsCableText                   string
	UpsDriver                      UpsDriver
	UpsDriverText                  string
	UpsMode                        UpsMode
	UpsModeText                    string

	// shutdown
	ShutdownBatteryMin          float64
	ShutdownTimeleftSecondsMin  int64
	ShutdownOnBatterySecondsMax int64
}

// InputSensivity type
type InputSensivity uint8

// InputSensivities type enum
var InputSensivities = map[string]interface{}{
	"Low":         InputSensivity(1),
	"Medium":      InputSensivity(2),
	"High":        InputSensivity(3),
	"Auto Adjust": InputSensivity(4),
	"Unknown":     InputSensivity(5),
}

// UpsTransferOnbatteryReason type
type UpsTransferOnbatteryReason uint8

// UpsTransferOnbatteryReasons enum
var UpsTransferOnbatteryReasons = map[string]interface{}{
	"No transfers since turnon":         UpsTransferOnbatteryReason(1),
	"Automatic or explicit self test":   UpsTransferOnbatteryReason(2),
	"Forced by software":                UpsTransferOnbatteryReason(3),
	"Low line voltage":                  UpsTransferOnbatteryReason(4),
	"High line voltage":                 UpsTransferOnbatteryReason(5),
	"Unacceptable line voltage changes": UpsTransferOnbatteryReason(6),
	"Line voltage notch or spike":       UpsTransferOnbatteryReason(7),
	"Input frequency out of range":      UpsTransferOnbatteryReason(8),
	"UNKNOWN EVENT":                     UpsTransferOnbatteryReason(9),
}

// UpsAlarmMode type
type UpsAlarmMode uint8

// UpsAlarmModes type enum
var UpsAlarmModes = map[string]interface{}{
	"No alarm":    UpsAlarmMode(1),
	"Always":      UpsAlarmMode(2),
	"5 Seconds":   UpsAlarmMode(3),
	"5":           UpsAlarmMode(3),
	"30 Seconds":  UpsAlarmMode(4),
	"30":          UpsAlarmMode(4),
	"Low Battery": UpsAlarmMode(5),
}

// UpsSelftestResult type
type UpsSelftestResult uint8

// UpsSelftestResults type enum
var UpsSelftestResults = map[string]interface{}{
	"NO": UpsSelftestResult(1),
	"OK": UpsSelftestResult(2),
	"BT": UpsSelftestResult(3),
	"NG": UpsSelftestResult(4),
	"IP": UpsSelftestResult(5),
	"WN": UpsSelftestResult(6),
	"??": UpsSelftestResult(7),
}

// UpsCable type
type UpsCable uint8

// UpsCables type enum
var UpsCables = map[string]interface{}{
	"Custom Cable Simple":  UpsCable(1),
	"APC Cable 940-0119A":  UpsCable(2),
	"APC Cable 940-0127A":  UpsCable(3),
	"APC Cable 940-0128A":  UpsCable(4),
	"APC Cable 940-0020B":  UpsCable(5),
	"APC Cable 940-0020C":  UpsCable(6),
	"APC Cable 940-0023A":  UpsCable(7),
	"MAM Cable 04-02-2000": UpsCable(8),
	"APC Cable 940-0095A":  UpsCable(9),
	"APC Cable 940-0095B":  UpsCable(10),
	"APC Cable 940-0095C":  UpsCable(11),
	"Custom Cable Smart":   UpsCable(12),
	"APC Cable 940-0024B":  UpsCable(121),
	"APC Cable 940-0024C":  UpsCable(122),
	"APC Cable 940-1524C":  UpsCable(123),
	"APC Cable 940-0024G":  UpsCable(124),
	"APC Cable 940-0625A":  UpsCable(125),
	"Ethernet Link":        UpsCable(13),
	"USB Cable":            UpsCable(14),
}

// UpsDriver type
type UpsDriver uint8

// UpsDrivers type enum
var UpsDrivers = map[string]interface{}{
	"DUMB UPS Driver":     UpsDriver(1),
	"APC Smart UPS (any)": UpsDriver(2),
	"USB UPS Driver":      UpsDriver(3),
	"NETWORK UPS Driver":  UpsDriver(4),
	"TEST UPS Driver":     UpsDriver(5),
	"PCNET UPS Driver":    UpsDriver(6),
	"SNMP UPS Driver":     UpsDriver(7),
	"MODBUS UPS Driver":   UpsDriver(8),
}

// UpsMode type
type UpsMode uint8

// UpsModes type enum
var UpsModes = map[string]interface{}{
	"Stand Alone":     UpsMode(1),
	"ShareUPS Slave":  UpsMode(2),
	"ShareUPS Master": UpsMode(3),
}

// UpsStatus type
type UpsStatus uint64

// UpsStatusFlags enum
var UpsStatusFlags = map[string]interface{}{
	// bit values for APC UPS Status Byte (ups->Status)
	"calibration": UpsStatus(0x00000001),
	"trim":        UpsStatus(0x00000002),
	"boost":       UpsStatus(0x00000004),
	"online":      UpsStatus(0x00000008),
	"onbatt":      UpsStatus(0x00000010),
	"overload":    UpsStatus(0x00000020),
	"battlow":     UpsStatus(0x00000040),
	"replacebatt": UpsStatus(0x00000080),

	// Extended bit values added by apcupsd
	"commlost":    UpsStatus(0x00000100), // Communications with UPS lost
	"shutdown":    UpsStatus(0x00000200), // Shutdown in progress
	"slave":       UpsStatus(0x00000400), // Set if this is a slave
	"slavedown":   UpsStatus(0x00000800), // Slave not responding
	"onbatt_msg":  UpsStatus(0x00020000), // Set when UPS_ONBATT message is sent
	"fastpoll":    UpsStatus(0x00040000), // Set on power failure to poll faster
	"shut_load":   UpsStatus(0x00080000), // Set when BatLoad <= percent
	"shut_btime":  UpsStatus(0x00100000), // Set when time on batts > maxtime
	"shut_ltime":  UpsStatus(0x00200000), // Set when TimeLeft <= runtime
	"shut_emerg":  UpsStatus(0x00400000), // Set when battery power has failed
	"shut_remote": UpsStatus(0x00800000), // Set when remote shutdown
	"plugged":     UpsStatus(0x01000000), // Set if computer is plugged into UPS
	"battpresent": UpsStatus(0x04000000), // Indicates if battery is connected
}
