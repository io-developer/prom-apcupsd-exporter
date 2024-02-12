package apcupsd

type BatteryTransferReason uint8

const (
	BATTERY_TRANSFER_REASON__NA         = BatteryTransferReason(0)
	BATTERY_TRANSFER_REASON__NONE       = BatteryTransferReason(1)
	BATTERY_TRANSFER_REASON__SELFTEST   = BatteryTransferReason(2)
	BATTERY_TRANSFER_REASON__FORCED     = BatteryTransferReason(3)
	BATTERY_TRANSFER_REASON__UNDERVOLT  = BatteryTransferReason(4)
	BATTERY_TRANSFER_REASON__OVERVOLT   = BatteryTransferReason(5)
	BATTERY_TRANSFER_REASON__RIPPLE     = BatteryTransferReason(6)
	BATTERY_TRANSFER_REASON__NOTCHSPIKE = BatteryTransferReason(7)
	BATTERY_TRANSFER_REASON__FREQ       = BatteryTransferReason(8)
	BATTERY_TRANSFER_REASON__UNKNOWN    = BatteryTransferReason(9)
)
