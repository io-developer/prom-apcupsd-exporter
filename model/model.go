package model

// Model data
type Model struct {
	State         *State
	PrevState     *State
	ChangedFields map[string][]interface{}
	onChange      []func(*Model)
}

// NewModel ..
func NewModel() *Model {
	return &Model{
		State:         NewState(),
		PrevState:     NewState(),
		ChangedFields: map[string][]interface{}{},
		onChange:      []func(*Model){},
	}
}

// AddOnChange ..
func (m *Model) AddOnChange(handler func(*Model)) {
	m.onChange = append(m.onChange, handler)
}

// Update method
func (m *Model) Update(newState *State) {
	prevState := m.State
	m.PrevState = prevState
	m.State = newState

	m.updateStatusCounts()

	_, diff := prevState.Compare(newState)
	m.ChangedFields = diff
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
