package apcupsd

type AlarmMode uint8

const (
	ALARM_MODE__NA = AlarmMode(iota)
	ALARM_MODE__DISABLED
	ALARM_MODE__ALWAYS
	ALARM_MODE__5_SEC
	ALARM_MODE__30_SEC
	ALARM_MODE__LOW_BATTERY
)

var ALARM_MODE_TO_NAME = map[AlarmMode]string{
	ALARM_MODE__NA:          "NA",
	ALARM_MODE__DISABLED:    "No alarm",
	ALARM_MODE__ALWAYS:      "Always",
	ALARM_MODE__5_SEC:       "5 Seconds",
	ALARM_MODE__30_SEC:      "30 Seconds",
	ALARM_MODE__LOW_BATTERY: "Low Battery",
}

var ALARM_MODE_TO_SHORT_NAME = map[AlarmMode]string{
	ALARM_MODE__NA:          "NA",
	ALARM_MODE__DISABLED:    "disabled",
	ALARM_MODE__ALWAYS:      "always",
	ALARM_MODE__5_SEC:       "5_sec",
	ALARM_MODE__30_SEC:      "30_sec",
	ALARM_MODE__LOW_BATTERY: "low_batt",
}

var ALARM_MODE_FROM_NAME = map[string]AlarmMode{
	"NA":          ALARM_MODE__NA,
	"No alarm":    ALARM_MODE__DISABLED,
	"Always":      ALARM_MODE__ALWAYS,
	"5 Seconds":   ALARM_MODE__5_SEC,
	"30 Seconds":  ALARM_MODE__30_SEC,
	"Low Battery": ALARM_MODE__LOW_BATTERY,
}

var ALARM_MODE_FROM_SHORT_NAME = map[string]AlarmMode{
	"NA":       ALARM_MODE__NA,
	"disabled": ALARM_MODE__DISABLED,
	"always":   ALARM_MODE__ALWAYS,
	"5_sec":    ALARM_MODE__5_SEC,
	"30_sec":   ALARM_MODE__30_SEC,
	"low_batt": ALARM_MODE__LOW_BATTERY,
}
