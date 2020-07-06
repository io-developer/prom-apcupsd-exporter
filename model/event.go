package model

import "time"

// Event ..
type Event struct {
	Ts   time.Time
	Type EventType
	Data map[string]interface{}
}

// EventType ..
type EventType string

// Event types
const (
	EventTypeCommlost    = EventType("commlost")
	EventTypeCommlostEnd = EventType("commlost_end")
	EventTypeOnbatt      = EventType("onbatt")
	EventTypeOnbattEnd   = EventType("onbatt_end")
	EventTypeOffline     = EventType("offline")
	EventTypeOnline      = EventType("online")
	EventTypeLineOk      = EventType("line_ok")
	EventTypeTrim        = EventType("trim")
	EventTypeBoost       = EventType("boost")
	EventTypeOverload    = EventType("overload")
	EventTypeOverloadEnd = EventType("overload_end")
	EventTypeNobatt      = EventType("nobatt")
	EventTypeNobattEnd   = EventType("nobatt_end")
	EventTypeTurnedOn    = EventType("turned_on")
	EventTypeTurnedOff   = EventType("turned_off")
)

func eventsFromStateChanges(prev State, curr State) []Event {
	m := eventsCalcMap(prev, curr)
	m = eventsReduceMap(m)

	events := []Event{}
	for eventType, enabled := range m {
		if enabled {
			events = append(events, eventFromType(eventType, prev, curr))
		}
	}

	return events
}

func eventFromType(t EventType, prev State, curr State) Event {
	event := Event{
		Ts:   time.Now(),
		Type: t,
	}

	if t == EventTypeOnbatt {
		event.Data = map[string]interface{}{
			"ts_start":    curr.UpsTransferOnBatteryDate.Unix(),
			"reason_type": curr.UpsTransferOnBatteryReason.Type,
			"reason_text": curr.UpsTransferOnBatteryReason.Text,
		}
	} else if t == EventTypeOnbattEnd {
		event.Data = map[string]interface{}{
			"ts_start":    curr.UpsTransferOnBatteryDate.Unix(),
			"ts_end":      curr.UpsTransferOffBatteryDate.Unix(),
			"seconds":     prev.UpsOnBatterySeconds,
			"reason_type": curr.UpsTransferOnBatteryReason.Type,
			"reason_text": curr.UpsTransferOnBatteryReason.Text,
		}
	}

	return event
}

func eventsCalcMap(prev State, curr State) map[EventType]bool {
	prevFlag := prev.UpsStatus.Flag
	currFlag := curr.UpsStatus.Flag

	onlineOkMask := StatusFlags["online"] | StatusFlags["trim"] | StatusFlags["boost"]
	onlineOkExpect := StatusFlags["online"]

	wasOnlineOk := (prevFlag & onlineOkMask) == onlineOkExpect
	isOnlineOk := (currFlag & onlineOkMask) == onlineOkExpect

	wasTurnedOff := (prevFlag&StatusFlags["plugged"] != 0) &&
		(prevFlag&StatusFlags["online"]) == 0 &&
		(prevFlag&StatusFlags["onbatt"]) == 0

	isTurnedOff := (currFlag&StatusFlags["plugged"] != 0) &&
		(currFlag&StatusFlags["online"]) == 0 &&
		(currFlag&StatusFlags["onbatt"]) == 0

	m := map[EventType]bool{
		EventTypeLineOk:    !wasOnlineOk && isOnlineOk,
		EventTypeTurnedOn:  wasTurnedOff && !isTurnedOff,
		EventTypeTurnedOff: !wasTurnedOff && isTurnedOff,

		EventTypeCommlost: (prevFlag&StatusFlags["commlost"]) == 0 &&
			(currFlag&StatusFlags["commlost"]) != 0,

		EventTypeCommlostEnd: (prevFlag&StatusFlags["commlost"]) != 0 &&
			(currFlag&StatusFlags["commlost"]) == 0,

		EventTypeOnbatt: (prevFlag&StatusFlags["onbatt"]) == 0 &&
			(currFlag&StatusFlags["onbatt"]) != 0,

		EventTypeOnbattEnd: (prevFlag&StatusFlags["onbatt"]) != 0 &&
			(currFlag&StatusFlags["onbatt"]) == 0,

		EventTypeOffline: (prevFlag&StatusFlags["online"]) != 0 &&
			(currFlag&StatusFlags["online"]) == 0,

		EventTypeOnline: (prevFlag&StatusFlags["online"]) == 0 &&
			(currFlag&StatusFlags["online"]) != 0,

		EventTypeTrim: (prevFlag&StatusFlags["trim"]) == 0 &&
			(currFlag&StatusFlags["trim"]) != 0,

		EventTypeBoost: (prevFlag&StatusFlags["boost"]) == 0 &&
			(currFlag&StatusFlags["boost"]) != 0,

		EventTypeOverload: (prevFlag&StatusFlags["overload"]) == 0 &&
			(currFlag&StatusFlags["overload"]) != 0,

		EventTypeOverloadEnd: (prevFlag&StatusFlags["overload"]) != 0 &&
			(currFlag&StatusFlags["overload"]) == 0,

		EventTypeNobatt: (prevFlag&StatusFlags["battpresent"]) != 0 &&
			(currFlag&StatusFlags["battpresent"]) == 0,

		EventTypeNobattEnd: (prevFlag&StatusFlags["battpresent"]) == 0 &&
			(currFlag&StatusFlags["battpresent"]) != 0,
	}

	return m
}

// disable unecessary events
func eventsReduceMap(m map[EventType]bool) map[EventType]bool {

	if m[EventTypeTurnedOn] {
		m[EventTypeLineOk] = false
		m[EventTypeOnline] = false
		m[EventTypeOverloadEnd] = false
		m[EventTypeNobattEnd] = false
		m[EventTypeCommlostEnd] = false
	}

	if m[EventTypeTurnedOff] {
		m[EventTypeNobatt] = false
		m[EventTypeOffline] = false
		m[EventTypeCommlostEnd] = false
		m[EventTypeOverloadEnd] = false
	}

	if m[EventTypeLineOk] {
		m[EventTypeOnline] = false
		m[EventTypeCommlostEnd] = false
		m[EventTypeNobattEnd] = false
	}

	if m[EventTypeOnline] {
		m[EventTypeCommlostEnd] = false
		m[EventTypeNobattEnd] = false
	}

	if m[EventTypeOnbattEnd] {
		m[EventTypeOnline] = false
		m[EventTypeLineOk] = false
	}

	return m
}
