package metric

import (
	"local/apcupsd_exporter/apc"
)

// Mapper ..
type Mapper interface {
	Map(raw string, def float64) float64
}

// DefaultMapper ..
type DefaultMapper struct {
	Mapper
}

// Map ..
func (m DefaultMapper) Map(raw string, def float64) float64 {
	if val, err := apc.ParseCommon(raw); err == nil {
		return val
	}
	return def
}

// DictMapper ..
type DictMapper struct {
	Mapper
	Dict map[string]float64
}

// Map ..
func (m DictMapper) Map(raw string, def float64) float64 {
	if val, exists := m.Dict[raw]; exists {
		return val
	}
	return def
}
