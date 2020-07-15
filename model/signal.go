package model

// Signal ..
type Signal string

// Signals
const (
	SignalCommfailure   = Signal("commfailure")
	SignalCommok        = Signal("commok")
	SignalStartselftest = Signal("startselftest")
	SignalEndselftest   = Signal("endselftest")
	SignalPowerout      = Signal("powerout")
	SignalMainsback     = Signal("mainsback")
	SignalOnbattery     = Signal("onbattery")
	SignalOffbattery    = Signal("offbattery")
	SignalBattattach    = Signal("battattach")
	SignalChangeme      = Signal("changeme")
	SignalFailing       = Signal("failing")
	SignalTimeout       = Signal("timeout")
	SignalLoadlimit     = Signal("loadlimit")
	SignalDoshutdown    = Signal("doshutdown")
	SignalAnnoyme       = Signal("annoyme")
	SignalRemotedown    = Signal("remotedown")
)
