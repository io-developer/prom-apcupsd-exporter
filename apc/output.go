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

// IsEmpty method
func (o *Output) IsEmpty() bool {
	return len(o.Parsed) == 0
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
		if val, err := parseNumber(raw); err == nil {
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
		if val, err := parseTime(raw); err == nil {
			return val
		}
	}
	return def
}

// GetSeconds ..
func (o *Output) GetSeconds(key string, def int64) int64 {
	if raw, exists := o.Parsed[key]; exists {
		if val, err := parseSeconds(raw); err == nil {
			return val
		}
	}
	return def
}

// GetMapped ..
func (o *Output) GetMapped(key string, kvMap map[string]interface{}, def interface{}) interface{} {
	if raw, exists := o.Parsed[key]; exists {
		if mapped, mappedExists := kvMap[raw]; mappedExists {
			return mapped
		}
	}
	return def
}

var (
	reSec = regexp.MustCompile(`(?i)^\s*(?P<number>[-+]?\d[\d.,]+)\s*(Seconds|Second|Sec)$`)
	reMin = regexp.MustCompile(`(?i)^\s*(?P<number>[-+]?\d[\d.,]+)\s*(Minutes|Minute|Min)$`)
	reDay = regexp.MustCompile(`(?i)^\s*(?P<number>[-+]?\d[\d.,]+)\s*(Days|Day)$`)
)

func parseSeconds(raw string) (val int64, err error) {
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
	valf64, err := parseNumber(str)
	return int64(valf64 * mult), err
}

func parseTime(raw string) (val time.Time, err error) {
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

func parseNumber(raw string) (val float64, err error) {
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
