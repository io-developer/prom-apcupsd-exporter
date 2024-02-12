package apcupsd

type Mode uint8

const (
	MODE__INVALID      = Mode(0)
	MODE__STAND_ALONE  = Mode(1)
	MODE__SHARE_SLAVE  = Mode(2)
	MODE__SHARE_MASTER = Mode(3)
)
