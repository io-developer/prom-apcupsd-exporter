package model

const defaultEventLimit = 30

// Model data
type Model struct {
	State         State
	PrevState     State
	events        []Event
	NewEvents     []Event
	EventLimit    int
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
		EventLimit:    defaultEventLimit,
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
func (m *Model) Update(newState State) {
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

	m.trimEvents()
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

			m.AddEvent(eventFromType(EventTypeOnbatt, old, curr))
			m.AddEvent(eventFromType(EventTypeOnbattEnd, old, curr))
		}
	}
}

// updateEvents method
func (m *Model) updateEvents() {
	events := eventsFromStateChanges(m.PrevState, m.State)
	m.NewEvents = append(m.NewEvents, events...)
}

// trimEvents method
func (m *Model) trimEvents() {
	if len(m.NewEvents) == 0 {
		return
	}
	m.events = append(m.events, m.NewEvents...)
	if len(m.events) > m.EventLimit {
		copy(m.events, m.events[len(m.events)-m.EventLimit:])
	}
	m.NewEvents = []Event{}
}
