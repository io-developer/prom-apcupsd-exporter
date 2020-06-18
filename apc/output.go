package apc

import (
	"math"
	"strconv"
	"strings"
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

// GetFloat64 value
func (o *Output) GetFloat64(name string) (val float64, exists bool) {
	str, exists := o.Parsed[name]
	if exists && str != "" {
		val, _ := strconv.ParseFloat(str, 64)
		return val, true
	}
	return math.NaN(), false
}
