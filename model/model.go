package model

// Model data
type Model struct {
	State         *State
	PrevState     *State
	ChangedFields map[string][]interface{}
	onChange      map[chan *Model]bool
}

// NewModel ..
func NewModel() *Model {
	return &Model{
		State:         NewState(),
		PrevState:     NewState(),
		ChangedFields: map[string][]interface{}{},
		onChange:      map[chan *Model]bool{},
	}
}

// AddOnChange ..
func (m *Model) AddOnChange(ch chan *Model) {
	m.onChange[ch] = true
}

// Update method
func (m *Model) Update(newState *State) {
	prevState := m.State
	m.PrevState = prevState
	m.State = newState

	updateStatusCounts(prevState, newState)

	_, diff := prevState.Compare(newState)
	m.ChangedFields = diff

	if len(diff) > 0 {
		for ch := range m.onChange {
			ch <- m
		}
	}
}

func updateStatusCounts(old *State, curr *State) {
	curr.UpsStatus.FlagChangeCounts = old.UpsStatus.CloneFlagChangeCounts()
	flags := curr.UpsStatus.GetFlags()
	prevFlags := old.UpsStatus.GetFlags()
	for flagName := range StatusFlags {
		if flags[flagName] != prevFlags[flagName] {
			curr.UpsStatus.FlagChangeCounts[flagName]++
		}
	}
}
