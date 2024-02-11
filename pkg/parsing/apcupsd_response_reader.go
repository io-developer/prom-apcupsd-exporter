package parsing

import (
	"time"

	"github.com/io-developer/prom-apcupsd-exporter/pkg/dto"
)

type ApcupsdResponseReader struct {
	Response *dto.ApcupsdResponse
	Parser   *ApcupsdParser
}

func NewApcupsdResponseReader(response *dto.ApcupsdResponse) *ApcupsdResponseReader {
	return &ApcupsdResponseReader{
		Response: response,
		Parser:   NewApcupsdParser(),
	}
}

func (r *ApcupsdResponseReader) GetValueAt(key string, def string) string {
	if val, exists := r.Response.KeyValues[key]; exists {
		return val
	}
	return def
}

func (r *ApcupsdResponseReader) GetFloatAt(key string, def float64) float64 {
	if raw, exists := r.Response.KeyValues[key]; exists {
		if val, err := r.Parser.ParseTextAsNumber(raw); err == nil {
			return val
		}
	}
	return def
}

func (r *ApcupsdResponseReader) GetUintAt(key string, def uint64) uint64 {
	if raw, exists := r.Response.KeyValues[key]; exists {
		if val, err := r.Parser.ParseTextAsUint(raw); err == nil {
			return val
		}
	}
	return def
}

func (r *ApcupsdResponseReader) GetTimeAt(key string, def time.Time) time.Time {
	if raw, exists := r.Response.KeyValues[key]; exists {
		if val, err := r.Parser.ParseTextAsTime(raw); err == nil {
			return val
		}
	}
	return def
}

func (r *ApcupsdResponseReader) GetDurationSecondsAt(key string, def int64) int64 {
	if raw, exists := r.Response.KeyValues[key]; exists {
		if val, err := r.Parser.ParseTextAsDurationSeconds(raw); err == nil {
			return val
		}
	}
	return def
}

func (r *ApcupsdResponseReader) GetMappedAt(key string, kvMap map[string]interface{}, def interface{}) interface{} {
	if raw, exists := r.Response.KeyValues[key]; exists {
		if mapped, mappedExists := kvMap[raw]; mappedExists {
			return mapped
		}
	}
	return def
}
