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

// GetParsed method
func (o *Output) GetParsed(key string, def string) string {
	return o.Get(key, def)
}

// Get ..
func (o *Output) Get(key string, def string) string {
	if val, exists := o.Parsed[key]; exists {
		return val
	}
	return def
}

// GetFloat ..
func (o *Output) GetFloat(key string, def float64) float64 {
	if raw, exists := o.Parsed[key]; exists {
		if val, err := ParseNumber(raw); err == nil {
			return val
		}
	}
	return def
}

// GetInt ..
func (o *Output) GetInt(key string, def int64) int64 {
	if raw, exists := o.Parsed[key]; exists {
		if val, err := strconv.ParseInt(raw, 0, 64); err == nil {
			return val
		}
	}
	return def
}

// GetUint ..
func (o *Output) GetUint(key string, def uint64) uint64 {
	if raw, exists := o.Parsed[key]; exists {
		if val, err := strconv.ParseUint(raw, 0, 64); err == nil {
			return val
		}
	}
	return def
}

// GetTime ..
func (o *Output) GetTime(key string, def time.Time) time.Time {
	if raw, exists := o.Parsed[key]; exists {
		if val, err := ParseTime(raw); err == nil {
			return val
		}
	}
	return def
}

// GetUnixtime ..
func (o *Output) GetUnixtime(key string, def int64) int64 {
	if raw, exists := o.Parsed[key]; exists {
		if val, err := ParseUnixtimeInt64(raw); err == nil {
			return val
		}
	}
	return def
}

// GetSeconds ..
func (o *Output) GetSeconds(key string, def int64) int64 {
	if raw, exists := o.Parsed[key]; exists {
		if val, err := ParseSecondsInt64(raw); err == nil {
			return val
		}
	}
	return def
}

// GetFlags ..
func (o *Output) GetFlags(key string, baseFlags map[string]uint64) map[string]uint64 {
	raw, exists := o.Parsed[key]
	if !exists {
		raw = ""
	}
	return ParseFlags(raw, baseFlags)
}

// GetMapped ..
func (o *Output) GetMapped(key string, kvMap map[string]interface{}, def uint64) interface{} {
	if raw, exists := o.Parsed[key]; exists {
		if mapped, mappedExists := kvMap[raw]; mappedExists {
			return mapped
		}
	}
	return def
}

// GetMappedUint64 ..
func (o *Output) GetMappedUint64(key string, kvMap map[string]uint64, def uint64) uint64 {
	if raw, exists := o.Parsed[key]; exists {
		if mapped, mappedExists := kvMap[raw]; mappedExists {
			return mapped
		}
	}
	return def
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

// ParseSecondsInt64 ..
func ParseSecondsInt64(raw string) (val int64, err error) {
	valf64, err := ParseSeconds(raw)
	return int64(valf64), err
}

// ParseUnixtime ..
func ParseUnixtime(raw string) (val float64, err error) {
	valInt, err := ParseUnixtimeInt64(raw)
	return float64(valInt), err
}

// ParseUnixtimeInt64 ..
func ParseUnixtimeInt64(raw string) (val int64, err error) {
	time, err := ParseTime(raw)
	return time.Unix(), nil
}

// ParseTime ..
func ParseTime(raw string) (val time.Time, err error) {
	t, err := time.Parse("2006-01-02 15:04:05 -0700", raw)
	if err != nil {
		t, err = time.Parse("2006-01-02 15:04:05", raw)
	}
	if err != nil {
		t, err = time.Parse("2006-01-02", raw)
	}
	if err != nil {
		return t, err
	}
	return t, nil
}

// ParseNumber ..
func ParseNumber(raw string) (val float64, err error) {
	numStr := regexp.MustCompile(`^[-+]?\d[\d.,]*`).FindString(raw)
	if numStr != "" {
		return strconv.ParseFloat(numStr, 64)
	}
	hexStr := regexp.MustCompile(`^0x[0-9a-fA-F]+`).FindString(raw)
	if hexStr != "" {
		intVal, err := strconv.ParseUint(raw, 0, 64)
		return float64(intVal), err
	}
	return 0, errors.New("None number parsed")
}

// ParseFlags ..
func ParseFlags(raw string, baseFlags map[string]uint64) map[string]uint64 {
	currentFlag := uint64(0)
	if parsed, err := strconv.ParseUint(raw, 0, 64); err == nil {
		currentFlag = parsed
	}
	result := map[string]uint64{}
	for flagName, flagVal := range baseFlags {
		result[flagName] = currentFlag & flagVal
	}
	return result
}
