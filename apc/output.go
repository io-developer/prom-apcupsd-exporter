package apc

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Output struct
type Output struct {
	Raw    string
	Parsed map[string]string
}

// NewOutput constructor
func NewOutput(raw string) *Output {
	return &Output{Raw: raw}
}

// Parse method
func (o *Output) Parse() {
	o.Parsed = make(map[string]string)
	for _, line := range strings.Split(o.Raw, "\n") {
		slice := strings.SplitN(line, ":", 2)
		if len(slice) == 2 {
			key := strings.Trim(slice[0], " \t")
			val := strings.Trim(slice[1], " \t")
			o.Parsed[key] = val
		}
	}
}

// ParseCommon ..
func ParseCommon(raw string) (val float64, err error) {
	val, err = ParseSeconds(raw)
	if err == nil {
		return
	}
	val, err = ParseUnixtime(raw)
	if err == nil {
		return
	}
	return ParseNumber(raw)
}

var (
	reSec = regexp.MustCompile(`(?i)^\s*(?P<number>[-+]?\d[\d.,]+)\s*(Seconds|Second|Sec)$`)
	reMin = regexp.MustCompile(`(?i)^\s*(?P<number>[-+]?\d[\d.,]+)\s*(Minutes|Minute|Min)$`)
	reDay = regexp.MustCompile(`(?i)^\s*(?P<number>[-+]?\d[\d.,]+)\s*(Days|Day)$`)
)

// ParseSeconds ..
func ParseSeconds(raw string) (val float64, err error) {
	str, mult := "", 1.0
	if m := reSec.FindStringSubmatch(raw); m != nil {
		str, mult = m[1], 1.0
	}
	if m := reMin.FindStringSubmatch(raw); m != nil {
		str, mult = m[1], 60
	}
	if m := reDay.FindStringSubmatch(raw); m != nil {
		str, mult = m[1], 24*3600
	}
	val, err = ParseNumber(str)
	return val * mult, err
}

// ParseUnixtime ..
func ParseUnixtime(raw string) (val float64, err error) {
	t, err := time.Parse("2006-01-02 15:04:05 -0700", raw)
	if err != nil {
		t, err = time.Parse("2006-01-02 15:04:05", raw)
	}
	if err != nil {
		t, err = time.Parse("2006-01-02", raw)
	}
	if err != nil {
		return 0, err
	}
	return float64(t.Unix()), nil
}

// ParseNumber ..
func ParseNumber(raw string) (val float64, err error) {
	re := regexp.MustCompile(`^[-+]?\s*\d[\d.,]+`)
	numStr := re.FindString(raw)
	if numStr == "" {
		return 0, errors.New("Empty text")
	}
	return strconv.ParseFloat(numStr, 64)
}

// StatFlags map
var StatFlags = map[string]uint64{
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

// ParseStatFlags ..
func ParseStatFlags(raw string) map[string]uint64 {
	currentFlag := uint64(0)
	if parsed, err := strconv.ParseUint(raw, 0, 64); err == nil {
		currentFlag = parsed
	}
	result := map[string]uint64{}
	for flagName, flagVal := range StatFlags {
		result[flagName] = currentFlag & flagVal
	}
	return result
}
