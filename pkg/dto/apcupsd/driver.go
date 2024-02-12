package apcupsd

type Driver uint8

const (
	DRIVER__NA       = Driver(0)
	DRIVER__DUMB     = Driver(1)
	DRIVER__APCSMART = Driver(2)
	DRIVER__USB      = Driver(3)
	DRIVER__NETWORK  = Driver(4)
	DRIVER__TEST     = Driver(5)
	DRIVER__PCNET    = Driver(6)
	DRIVER__SNMPLITE = Driver(7)
	DRIVER__MODBUS   = Driver(8)
)
