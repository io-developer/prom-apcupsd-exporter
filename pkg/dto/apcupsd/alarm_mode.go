package apcupsd

type AlarmMode uint8

const (
	ALARM_MODE__NA          = AlarmMode(0)
	ALARM_MODE__NONE        = AlarmMode(1)
	ALARM_MODE__ALWAYS      = AlarmMode(2)
	ALARM_MODE__5_SEC       = AlarmMode(3)
	ALARM_MODE__30_SEC      = AlarmMode(4)
	ALARM_MODE__LOW_BATTERY = AlarmMode(5)
)
