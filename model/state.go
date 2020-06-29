package model

import (
	"reflect"
	"time"
)

// State - parsed 'apcaccess status' result
type State struct {

	// input
	InputSensivity           Sensivity
	InputFrequency           float64
	InputVoltage             float64
	InputVoltageMin          float64
	InputVoltageMax          float64
	InputVoltageNominal      float64
	InputVoltageTransferLow  float64
	InputVoltageTransferHigh float64

	// output
	OutputLoad                 float64
	OutputAmps                 float64
	OutputPowerNominal         float64
	OutputPowerApparentNominal float64
	OutputVoltage              float64
	OutputVoltageNominal       float64

	// battery
	BatteryCharge         float64
	BatteryVoltage        float64
	BatteryVoltageNominal float64
	BatteryExternalCount  uint16
	BatteryBadCount       uint16
	BatteryReplacedDate   time.Time

	// ups
	UpsManafacturedDate           time.Time
	UpsModel                      string
	UpsSerial                     string
	UpsFirmware                   string
	UpsName                       string
	UpsStatus                     Status
	UpsDipSwitchFlag              uint64
	UpsReg1                       uint64
	UpsReg2                       uint64
	UpsReg3                       uint64
	UpsTimeleftSeconds            int64
	UpsTimeleftSecondsLowBattery  int64
	UpsTransferOnBatteryCount     uint64
	UpsTransferOnBatteryReason    TransferOnbatteryReason
	UpsTransferOnBatteryDate      time.Time
	UpsTransferOffBatteryDate     time.Time
	UpsOnBatterySeconds           int64
	UpsOnBatterySecondsCumulative int64
	UpsTurnOffDelaySeconds        int64
	UpsTurnOnDelaySeconds         int64
	UpsTurnOnBatteryMin           float64
	UpsTempInternal               float64
	UpsTempAmbient                float64
	UpsHumidity                   float64
	UpsAlarmMode                  AlarmMode
	UpsSelftestResult             SelftestResult
	UpsSelftestIntervalSeconds    int64
	UpsCable                      Cable
	UpsDriver                     Driver
	UpsMode                       Mode

	// shutdown
	ShutdownBatteryMin          float64
	ShutdownTimeleftSecondsMin  int64
	ShutdownOnBatterySecondsMax int64

	// apcupsd
	ApcupsdHost    string
	ApcupsdVersion string
}

// NewState cotructor
func NewState() *State {
	return &State{
		InputSensivity:             Sensivity{},
		BatteryReplacedDate:        time.Time{},
		UpsManafacturedDate:        time.Time{},
		UpsStatus:                  NewStatus(0, ""),
		UpsTransferOnBatteryReason: TransferOnbatteryReason{},
		UpsTransferOnBatteryDate:   time.Time{},
		UpsTransferOffBatteryDate:  time.Time{},
		UpsAlarmMode:               AlarmMode{},
		UpsSelftestResult:          SelftestResult{},
		UpsCable:                   Cable{},
		UpsDriver:                  Driver{},
		UpsMode:                    Mode{},
	}
}

// Compare method
func (s *State) Compare(b *State) (equal bool, diff map[string][]interface{}) {
	diff = map[string][]interface{}{}
	typeElem := reflect.TypeOf(s).Elem()
	elemA := reflect.ValueOf(s).Elem()
	elemB := reflect.ValueOf(b).Elem()
	for i := 0; i < typeElem.NumField(); i++ {
		equal := false
		field := typeElem.Field(i)
		fieldA := elemA.Field(i)
		fieldB := elemB.Field(i)
		if method, exists := field.Type.MethodByName("Equal"); exists {
			equal = method.Func.Call([]reflect.Value{fieldA, fieldB})[0].Bool()
		} else {
			equal = reflect.DeepEqual(fieldA.Interface(), fieldB.Interface())
		}
		if !equal {
			diff[field.Name] = []interface{}{fieldA.Interface(), fieldB.Interface()}
		}
	}
	return len(diff) == 0, diff
}

// Sensivity ..
type Sensivity struct {
	Type SensivityType
	Text string
}

// SensivityType type
type SensivityType uint8

// SensivityTypes type enum
var SensivityTypes = map[string]interface{}{
	"Low":         SensivityType(1),
	"Medium":      SensivityType(2),
	"High":        SensivityType(3),
	"Auto Adjust": SensivityType(4),
	"Unknown":     SensivityType(5),
}

// TransferOnbatteryReason ..
type TransferOnbatteryReason struct {
	Type TransferOnbatteryReasonType
	Text string
}

// TransferOnbatteryReasonType type
type TransferOnbatteryReasonType uint8

// TransferOnbatteryReasonTypes enum
var TransferOnbatteryReasonTypes = map[string]interface{}{
	"No transfers since turnon":         TransferOnbatteryReasonType(1),
	"Automatic or explicit self test":   TransferOnbatteryReasonType(2),
	"Forced by software":                TransferOnbatteryReasonType(3),
	"Low line voltage":                  TransferOnbatteryReasonType(4),
	"High line voltage":                 TransferOnbatteryReasonType(5),
	"Unacceptable line voltage changes": TransferOnbatteryReasonType(6),
	"Line voltage notch or spike":       TransferOnbatteryReasonType(7),
	"Input frequency out of range":      TransferOnbatteryReasonType(8),
	"UNKNOWN EVENT":                     TransferOnbatteryReasonType(9),
}

// AlarmMode ..
type AlarmMode struct {
	Type AlarmModeType
	Text string
}

// AlarmModeType type
type AlarmModeType uint8

// AlarmModeTypes type enum
var AlarmModeTypes = map[string]interface{}{
	"No alarm":    AlarmModeType(1),
	"Always":      AlarmModeType(2),
	"5 Seconds":   AlarmModeType(3),
	"30 Seconds":  AlarmModeType(4),
	"Low Battery": AlarmModeType(5),
}

// SelftestResult ..
type SelftestResult struct {
	Type SelftestResultType
	Text string
}

// SelftestResultType type
type SelftestResultType uint8

// SelftestResultTypes type enum
var SelftestResultTypes = map[string]interface{}{
	"NO": SelftestResultType(1),
	"OK": SelftestResultType(2),
	"BT": SelftestResultType(3),
	"NG": SelftestResultType(4),
	"IP": SelftestResultType(5),
	"WN": SelftestResultType(6),
	"??": SelftestResultType(7),
}

// Cable ..
type Cable struct {
	Type CableType
	Text string
}

// CableType type
type CableType uint8

// CableTypes type enum
var CableTypes = map[string]interface{}{
	"Custom Cable Simple":  CableType(1),
	"APC Cable 940-0119A":  CableType(2),
	"APC Cable 940-0127A":  CableType(3),
	"APC Cable 940-0128A":  CableType(4),
	"APC Cable 940-0020B":  CableType(5),
	"APC Cable 940-0020C":  CableType(6),
	"APC Cable 940-0023A":  CableType(7),
	"MAM Cable 04-02-2000": CableType(8),
	"APC Cable 940-0095A":  CableType(9),
	"APC Cable 940-0095B":  CableType(10),
	"APC Cable 940-0095C":  CableType(11),
	"Custom Cable Smart":   CableType(12),
	"APC Cable 940-0024B":  CableType(121),
	"APC Cable 940-0024C":  CableType(122),
	"APC Cable 940-1524C":  CableType(123),
	"APC Cable 940-0024G":  CableType(124),
	"APC Cable 940-0625A":  CableType(125),
	"Ethernet Link":        CableType(13),
	"USB Cable":            CableType(14),
}

// Driver ..
type Driver struct {
	Type DriverType
	Text string
}

// DriverType type
type DriverType uint8

// DriverTypes type enum
var DriverTypes = map[string]interface{}{
	"DUMB UPS Driver":     DriverType(1),
	"APC Smart UPS (any)": DriverType(2),
	"USB UPS Driver":      DriverType(3),
	"NETWORK UPS Driver":  DriverType(4),
	"TEST UPS Driver":     DriverType(5),
	"PCNET UPS Driver":    DriverType(6),
	"SNMP UPS Driver":     DriverType(7),
	"MODBUS UPS Driver":   DriverType(8),
}

// Mode ..
type Mode struct {
	Type ModeType
	Text string
}

// ModeType type
type ModeType uint8

// ModeTypes type enum
var ModeTypes = map[string]interface{}{
	"Stand Alone":     ModeType(1),
	"ShareUPS Slave":  ModeType(2),
	"ShareUPS Master": ModeType(3),
}

// Status ..
type Status struct {
	Text             string
	Flag             uint64
	FlagChangeCounts map[string]uint64
}

// NewStatus func
func NewStatus(flag uint64, text string) Status {
	counts := make(map[string]uint64, len(StatusFlags))
	for flagName := range StatusFlags {
		counts[flagName] = 0
	}
	return Status{
		Text:             text,
		Flag:             flag,
		FlagChangeCounts: counts,
	}
}

// Equal method
func (s Status) Equal(b Status) bool {
	return s.Flag == b.Flag && s.Text == b.Text
}

// GetFlags method
func (s Status) GetFlags() map[string]uint64 {
	flags := make(map[string]uint64, len(StatusFlags))
	for flagName, flagVal := range StatusFlags {
		flags[flagName] = s.Flag & flagVal
	}
	return flags
}

// GetNormedFlags method
func (s Status) GetNormedFlags(invert bool) map[string]uint8 {
	match, notMatch := uint8(1), uint8(0)
	if invert {
		match, notMatch = 0, 1
	}
	flags := make(map[string]uint8, len(StatusFlags))
	for flagName, flagVal := range StatusFlags {
		if s.Flag&flagVal > 0 {
			flags[flagName] = match
		} else {
			flags[flagName] = notMatch
		}
	}
	return flags
}

// CloneFlagChangeCounts method
func (s Status) CloneFlagChangeCounts() map[string]uint64 {
	clone := make(map[string]uint64, len(s.FlagChangeCounts))
	for k, v := range s.FlagChangeCounts {
		clone[k] = v
	}
	return clone
}

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
