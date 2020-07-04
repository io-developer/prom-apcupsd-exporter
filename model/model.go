package model

import "time"

const defaultEventLimit = 30

// Model data
type Model struct {
	State         *State
	PrevState     *State
	events        []Event
	NewEvents     []Event
	ChangedFields map[string][]interface{}
	onChange      []func(*Model)
}

// NewModel ..
func NewModel() *Model {
	return &Model{
		State:         NewState(),
		PrevState:     NewState(),
		ChangedFields: map[string][]interface{}{},
		events:        []Event{},
		NewEvents:     []Event{},
		onChange:      []func(*Model){},
	}
}

// AddOnChange ..
func (m *Model) AddOnChange(handler func(*Model)) {
	m.onChange = append(m.onChange, handler)
}

// GetEvents ..
func (m *Model) GetEvents() []Event {
	return m.events
}

// AddEvent ..
func (m *Model) AddEvent(e Event) {
	m.NewEvents = append(m.NewEvents, e)
}

// Update method
func (m *Model) Update(newState *State) {
	prevState := m.State
	m.PrevState = prevState
	m.State = newState

	m.updateStatusCounts()
	m.updateTransferOnbatt()
	m.updateEvents()

	_, diff := prevState.Compare(newState)
	m.ChangedFields = diff

	if len(diff) > 0 || len(m.NewEvents) > 0 {
		for _, f := range m.onChange {
			f(m)
		}
	}

	m.trimEvents(defaultEventLimit)
}

// updateStatusCounts method
func (m *Model) updateStatusCounts() {
	old := m.PrevState
	curr := m.State
	curr.UpsStatus.FlagChangeCounts = old.UpsStatus.CloneFlagChangeCounts()
	flags := curr.UpsStatus.GetFlags()
	prevFlags := old.UpsStatus.GetFlags()
	for flagName := range StatusFlags {
		if flags[flagName] != prevFlags[flagName] {
			curr.UpsStatus.FlagChangeCounts[flagName]++
		}
	}
}

// updateTransferOnbatt method
func (m *Model) updateTransferOnbatt() {
	old := m.PrevState
	curr := m.State

	// Reveal hidden quick transfers on battery and back
	transDelta := int64(curr.UpsTransferOnBatteryCount) - int64(old.UpsTransferOnBatteryCount)
	if transDelta > 0 {
		minIncr := uint64(transDelta-1) * 2
		curr.UpsStatus.FlagChangeCounts["online"] += minIncr
		curr.UpsStatus.FlagChangeCounts["onbatt"] += minIncr

		flag := curr.UpsStatus.Flag
		if flag&StatusFlags["online"] != 0 {
			curr.UpsStatus.FlagChangeCounts["online"] += 2
		}
		if flag&StatusFlags["onbatt"] == 0 {
			curr.UpsStatus.FlagChangeCounts["onbatt"] += 2
		}
	}
}

// updateEvents method
func (m *Model) updateEvents() {
	prev := m.PrevState
	curr := m.State

	prevFlag := prev.UpsStatus.Flag
	currFlag := curr.UpsStatus.Flag

	if (prevFlag&StatusFlags["commlost"]) == 0 && (currFlag&StatusFlags["commlost"]) != 0 {
		m.AddEvent(Event{
			Ts:   time.Now(),
			Type: "commlost",
		})
		return
	}
	if (prevFlag&StatusFlags["commlost"]) != 0 && (currFlag&StatusFlags["commlost"]) == 0 {
		m.AddEvent(Event{
			Ts:   time.Now(),
			Type: "commlost_end",
		})
		return
	}

	if (prevFlag&StatusFlags["onbatt"]) == 0 && (currFlag&StatusFlags["onbatt"]) != 0 {
		m.AddEvent(Event{
			Ts:   time.Now(),
			Type: "onbatt",
			Data: map[string]interface{}{
				"ts_start":    curr.UpsTransferOnBatteryDate.Unix(),
				"reason_type": curr.UpsTransferOnBatteryReason.Type,
				"reason_text": curr.UpsTransferOnBatteryReason.Text,
			},
		})
	} else if (prevFlag&StatusFlags["onbatt"]) != 0 && (currFlag&StatusFlags["onbatt"]) == 0 {
		m.AddEvent(Event{
			Ts:   time.Now(),
			Type: "onbatt_end",
			Data: map[string]interface{}{
				"ts_start":    curr.UpsTransferOnBatteryDate.Unix(),
				"ts_end":      curr.UpsTransferOffBatteryDate.Unix(),
				"seconds":     prev.UpsOnBatterySeconds,
				"reason_type": curr.UpsTransferOnBatteryReason.Type,
				"reason_text": curr.UpsTransferOnBatteryReason.Text,
			},
		})
	} else if (prevFlag&StatusFlags["online"]) != 0 && (currFlag&StatusFlags["online"]) == 0 {
		m.AddEvent(Event{
			Ts:   time.Now(),
			Type: "offline",
		})
	} else if (prevFlag&StatusFlags["online"]) == 0 && (currFlag&StatusFlags["online"]) != 0 {
		m.AddEvent(Event{
			Ts:   time.Now(),
			Type: "online",
		})
	}

	onlineOkMask := StatusFlags["online"] | StatusFlags["trim"] | StatusFlags["boost"]
	onlineOkExpect := StatusFlags["online"]

	wasOnlineOk := (prevFlag & onlineOkMask) == onlineOkExpect
	isOnlineOk := (currFlag & onlineOkMask) == onlineOkExpect

	if !wasOnlineOk && isOnlineOk {
		m.AddEvent(Event{
			Ts:   time.Now(),
			Type: "line_ok",
		})
	}
	if (prevFlag&StatusFlags["trim"]) == 0 && (currFlag&StatusFlags["trim"]) != 0 {
		m.AddEvent(Event{
			Ts:   time.Now(),
			Type: "trim",
		})
	}
	if (prevFlag&StatusFlags["boost"]) == 0 && (currFlag&StatusFlags["boost"]) != 0 {
		m.AddEvent(Event{
			Ts:   time.Now(),
			Type: "boost",
		})
	}

	if (prevFlag&StatusFlags["overload"]) == 0 && (currFlag&StatusFlags["overload"]) != 0 {
		m.AddEvent(Event{
			Ts:   time.Now(),
			Type: "overload",
		})
	} else if (prevFlag&StatusFlags["overload"]) != 0 && (currFlag&StatusFlags["overload"]) == 0 {
		m.AddEvent(Event{
			Ts:   time.Now(),
			Type: "overload_end",
		})
	}

	if (prevFlag&StatusFlags["battpresent"]) != 0 && (currFlag&StatusFlags["battpresent"]) == 0 {
		m.AddEvent(Event{
			Ts:   time.Now(),
			Type: "nobatt",
		})
	} else if (prevFlag&StatusFlags["battpresent"]) == 0 && (currFlag&StatusFlags["battpresent"]) != 0 {
		m.AddEvent(Event{
			Ts:   time.Now(),
			Type: "nobatt_end",
		})
	}
}

// trimEvents method
func (m *Model) trimEvents(limit int) {
	if len(m.NewEvents) == 0 {
		return
	}
	m.events = append(m.events, m.NewEvents...)
	if len(m.events) > limit {
		copy(m.events, m.events[len(m.events)-limit:])
	}
	m.NewEvents = []Event{}
}

// Event ..
type Event struct {
	Ts   time.Time
	Type string
	Data map[string]interface{}
}
