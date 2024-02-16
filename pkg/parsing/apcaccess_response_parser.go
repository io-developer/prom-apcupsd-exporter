package parsing

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/io-developer/prom-apcupsd-exporter/pkg/dto"
)

type ApcaccessParser struct {
	ReDurationSeconds *regexp.Regexp
	ReDurationMinutes *regexp.Regexp
	ReDurationDays    *regexp.Regexp
}

func NewApcaccessParser() *ApcaccessParser {
	return &ApcaccessParser{
		ReDurationSeconds: regexp.MustCompile(`(?i)^\s*(?P<number>[-+]?\d[\d.,]+)\s*(Seconds|Second|Sec)$`),
		ReDurationMinutes: regexp.MustCompile(`(?i)^\s*(?P<number>[-+]?\d[\d.,]+)\s*(Minutes|Minute|Min)$`),
		ReDurationDays:    regexp.MustCompile(`(?i)^\s*(?P<number>[-+]?\d[\d.,]+)\s*(Days|Day)$`),
	}
}

func (p *ApcaccessParser) ParseOutput(output string) (*dto.ApcaccessResponse, error) {
	keyVals := make(map[string]string)
	for _, line := range strings.Split(output, "\n") {
		slice := strings.SplitN(line, ":", 2)
		if len(slice) == 2 {
			key := strings.Trim(slice[0], " \t")
			val := strings.Trim(slice[1], " \t")
			keyVals[key] = val
		}
	}
	return &dto.ApcaccessResponse{
		Output:    output,
		KeyValues: keyVals,
	}, nil
}

func (r *ApcaccessParser) ParseTextAsUint(text string) (uint64, error) {
	return strconv.ParseUint(text, 0, 64)
}

func (p *ApcaccessParser) ParseTextAsNumber(text string) (val float64, err error) {
	numStr := regexp.MustCompile(`^[-+]?\d[\d.,]*`).FindString(text)
	if numStr != "" {
		return strconv.ParseFloat(numStr, 64)
	}
	hexStr := regexp.MustCompile(`^0x[0-9a-fA-F]+`).FindString(text)
	if hexStr != "" {
		intVal, err := strconv.ParseUint(text, 0, 64)
		return float64(intVal), err
	}
	return 0, errors.New("None number parsed")
}

func (p *ApcaccessParser) ParseTextAsTime(text string) (val time.Time, err error) {
	t, err := time.Parse("2006-01-02 15:04:05 -0700", text)
	if err != nil {
		t, err = time.Parse("2006-01-02 15:04:05", text)
	}
	if err != nil {
		t, err = time.Parse("2006-01-02", text)
	}
	if err != nil {
		t, err = time.Parse("01/02/06", text)
	}
	if err != nil {
		return t, err
	}
	return t, nil
}

func (p *ApcaccessParser) ParseTextAsDurationSeconds(text string) (val int64, err error) {
	str, mult := "", 1.0
	if m := p.ReDurationSeconds.FindStringSubmatch(text); m != nil {
		str, mult = m[1], 1.0
	}
	if m := p.ReDurationMinutes.FindStringSubmatch(text); m != nil {
		str, mult = m[1], 60
	}
	if m := p.ReDurationDays.FindStringSubmatch(text); m != nil {
		str, mult = m[1], 24*3600
	}
	valf64, err := p.ParseTextAsNumber(str)
	return int64(valf64 * mult), err
}
