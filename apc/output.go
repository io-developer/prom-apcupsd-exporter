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
