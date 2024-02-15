package parsing

import (
	"time"

	"github.com/io-developer/prom-apcupsd-exporter/pkg/dto/apcupsd"
)

type ApcaccessResponseReader struct {
	Response *apcupsd.ApcaccessResponse
	Parser   *ApcaccessParser
}

func NewApcaccessResponseReader(response *apcupsd.ApcaccessResponse) *ApcaccessResponseReader {
	return &ApcaccessResponseReader{
		Response: response,
		Parser:   NewApcaccessParser(),
	}
}

func (r *ApcaccessResponseReader) GetValueAt(key string, def string) string {
	if val, exists := r.Response.KeyValues[key]; exists {
		return val
	}
	return def
}

func (r *ApcaccessResponseReader) GetFloatAt(key string, def float64) float64 {
	if raw, exists := r.Response.KeyValues[key]; exists {
		if val, err := r.Parser.ParseTextAsNumber(raw); err == nil {
			return val
		}
	}
	return def
}

func (r *ApcaccessResponseReader) GetUintAt(key string, def uint64) uint64 {
	if raw, exists := r.Response.KeyValues[key]; exists {
		if val, err := r.Parser.ParseTextAsUint(raw); err == nil {
			return val
		}
	}
	return def
}

func (r *ApcaccessResponseReader) GetTimeAt(key string, def time.Time) time.Time {
	if raw, exists := r.Response.KeyValues[key]; exists {
		if val, err := r.Parser.ParseTextAsTime(raw); err == nil {
			return val
		}
	}
	return def
}

func (r *ApcaccessResponseReader) GetDurationAt(key string, def time.Duration) time.Duration {
	if raw, exists := r.Response.KeyValues[key]; exists {
		if val, err := r.Parser.ParseTextAsDurationSeconds(raw); err == nil {
			return time.Duration(val) * time.Second
		}
	}
	return def
}

func (r *ApcaccessResponseReader) GetDurationSecondsAt(key string, def int64) int64 {
	if raw, exists := r.Response.KeyValues[key]; exists {
		if val, err := r.Parser.ParseTextAsDurationSeconds(raw); err == nil {
			return val
		}
	}
	return def
}

func (r *ApcaccessResponseReader) GetMappedAt(key string, kvMap map[string]interface{}, def interface{}) interface{} {
	if raw, exists := r.Response.KeyValues[key]; exists {
		if mapped, mappedExists := kvMap[raw]; mappedExists {
			return mapped
		}
	}
	return def
}
