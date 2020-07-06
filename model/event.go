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
)

func eventsFromStateChanges(prev State, curr State) []Event {
	prevFlag := prev.UpsStatus.Flag
	currFlag := curr.UpsStatus.Flag

	if (prevFlag&StatusFlags["commlost"]) == 0 && (currFlag&StatusFlags["commlost"]) != 0 {
		return []Event{{
			Ts:   time.Now(),
			Type: EventTypeCommlost,
		}}
	}
	if (prevFlag&StatusFlags["commlost"]) != 0 && (currFlag&StatusFlags["commlost"]) == 0 {
		return []Event{{
			Ts:   time.Now(),
			Type: EventTypeCommlostEnd,
		}}
	}

	events := []Event{}

	if (prevFlag&StatusFlags["onbatt"]) == 0 && (currFlag&StatusFlags["onbatt"]) != 0 {
		events = append(events, Event{
			Ts:   time.Now(),
			Type: EventTypeOnbatt,
			Data: map[string]interface{}{
				"ts_start":    curr.UpsTransferOnBatteryDate.Unix(),
				"reason_type": curr.UpsTransferOnBatteryReason.Type,
				"reason_text": curr.UpsTransferOnBatteryReason.Text,
			},
		})
	} else if (prevFlag&StatusFlags["onbatt"]) != 0 && (currFlag&StatusFlags["onbatt"]) == 0 {
		events = append(events, Event{
			Ts:   time.Now(),
			Type: EventTypeOnbattEnd,
			Data: map[string]interface{}{
				"ts_start":    curr.UpsTransferOnBatteryDate.Unix(),
				"ts_end":      curr.UpsTransferOffBatteryDate.Unix(),
				"seconds":     prev.UpsOnBatterySeconds,
				"reason_type": curr.UpsTransferOnBatteryReason.Type,
				"reason_text": curr.UpsTransferOnBatteryReason.Text,
			},
		})
	} else if (prevFlag&StatusFlags["online"]) != 0 && (currFlag&StatusFlags["online"]) == 0 {
		events = append(events, Event{
			Ts:   time.Now(),
			Type: EventTypeOffline,
		})
	} else if (prevFlag&StatusFlags["online"]) == 0 && (currFlag&StatusFlags["online"]) != 0 {
		events = append(events, Event{
			Ts:   time.Now(),
			Type: EventTypeOnline,
		})
	}

	onlineOkMask := StatusFlags["online"] | StatusFlags["trim"] | StatusFlags["boost"]
	onlineOkExpect := StatusFlags["online"]

	wasOnlineOk := (prevFlag & onlineOkMask) == onlineOkExpect
	isOnlineOk := (currFlag & onlineOkMask) == onlineOkExpect

	if !wasOnlineOk && isOnlineOk {
		events = append(events, Event{
			Ts:   time.Now(),
			Type: EventTypeLineOk,
		})
	}
	if (prevFlag&StatusFlags["trim"]) == 0 && (currFlag&StatusFlags["trim"]) != 0 {
		events = append(events, Event{
			Ts:   time.Now(),
			Type: EventTypeTrim,
		})
	}
	if (prevFlag&StatusFlags["boost"]) == 0 && (currFlag&StatusFlags["boost"]) != 0 {
		events = append(events, Event{
			Ts:   time.Now(),
			Type: EventTypeBoost,
		})
	}

	if (prevFlag&StatusFlags["overload"]) == 0 && (currFlag&StatusFlags["overload"]) != 0 {
		events = append(events, Event{
			Ts:   time.Now(),
			Type: EventTypeOverload,
		})
	} else if (prevFlag&StatusFlags["overload"]) != 0 && (currFlag&StatusFlags["overload"]) == 0 {
		events = append(events, Event{
			Ts:   time.Now(),
			Type: EventTypeOverloadEnd,
		})
	}

	if (prevFlag&StatusFlags["battpresent"]) != 0 && (currFlag&StatusFlags["battpresent"]) == 0 {
		events = append(events, Event{
			Ts:   time.Now(),
			Type: EventTypeNobatt,
		})
	} else if (prevFlag&StatusFlags["battpresent"]) == 0 && (currFlag&StatusFlags["battpresent"]) != 0 {
		events = append(events, Event{
			Ts:   time.Now(),
			Type: EventTypeNobattEnd,
		})
	}

	return events
}
