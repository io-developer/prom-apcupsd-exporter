package apcupsd

type Driver uint8

const (
	DRIVER__NA = Driver(iota)
	DRIVER__DUMB
	DRIVER__APCSMART
	DRIVER__USB
	DRIVER__NETWORK
	DRIVER__TEST
	DRIVER__PCNET
	DRIVER__SNMPLITE
	DRIVER__MODBUS
)

var DRIVER_TO_NAME = map[Driver]string{
	DRIVER__NA:       "*invalid-ups-type*",
	DRIVER__DUMB:     "DUMB UPS Driver",
	DRIVER__APCSMART: "APC Smart UPS (any)",
	DRIVER__USB:      "USB UPS Driver",
	DRIVER__SNMPLITE: "SNMP UPS Driver",
	DRIVER__NETWORK:  "NETWORK UPS Driver",
	DRIVER__TEST:     "TEST UPS Driver",
	DRIVER__PCNET:    "PCNET UPS Driver",
	DRIVER__MODBUS:   "MODBUS UPS Driver",
}

var DRIVER_TO_SHORT_NAME = map[Driver]string{
	DRIVER__NA:       "NA",
	DRIVER__DUMB:     "dumb",
	DRIVER__APCSMART: "apcsmart",
	DRIVER__USB:      "usb",
	DRIVER__SNMPLITE: "snmp",
	DRIVER__NETWORK:  "net",
	DRIVER__TEST:     "test",
	DRIVER__PCNET:    "pcnet",
	DRIVER__MODBUS:   "modbus",
}

var DRIVER_FROM_NAME = map[string]Driver{
	"*invalid-ups-type*":  DRIVER__NA,
	"DUMB UPS Driver":     DRIVER__DUMB,
	"APC Smart UPS (any)": DRIVER__APCSMART,
	"USB UPS Driver":      DRIVER__USB,
	"SNMP UPS Driver":     DRIVER__SNMPLITE,
	"NETWORK UPS Driver":  DRIVER__NETWORK,
	"TEST UPS Driver":     DRIVER__TEST,
	"PCNET UPS Driver":    DRIVER__PCNET,
	"MODBUS UPS Driver":   DRIVER__MODBUS,
}

var DRIVER_FROM_SHORT_NAME = map[string]Driver{
	"NA":       DRIVER__NA,
	"dumb":     DRIVER__DUMB,
	"apcsmart": DRIVER__APCSMART,
	"usb":      DRIVER__USB,
	"snmp":     DRIVER__SNMPLITE,
	"net":      DRIVER__NETWORK,
	"test":     DRIVER__TEST,
	"pcnet":    DRIVER__PCNET,
	"modbus":   DRIVER__MODBUS,
}
