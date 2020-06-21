package apc

import (
	"errors"
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
	val, err = strconv.ParseFloat(raw, 64)
	return
}

// ParseSeconds ..
func ParseSeconds(raw string) (val float64, err error) {
	val, err = strconv.ParseFloat(raw, 64)
	err = errors.New("Not implemented")
	return
}

// ParseUnixtime ..
func ParseUnixtime(raw string) (val float64, err error) {
	return -1, errors.New("Not implemented")
}
