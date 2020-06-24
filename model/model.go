package model

// Model data
type Model struct {
	State         *State
	PrevState     *State
	ChangedFields []string
}

// NewModel ..
func NewModel() *Model {
	return &Model{
		State:         NewState(),
		PrevState:     NewState(),
		ChangedFields: []string{},
	}
}

// Update method
func (m *Model) Update(newState *State) {
	m.PrevState = m.State
	m.State = newState
	m.ChangedFields = []string{}

	updateStatusCounts(m.PrevState, m.State)

	_, diffFields := m.State.Compare(m.PrevState)
	m.ChangedFields = diffFields
}

func updateStatusCounts(old *State, curr *State) {
	curr.UpsStatus.FlagChangeCounts = old.UpsStatus.FlagChangeCounts
	flags := curr.UpsStatus.GetFlags()
	prevFlags := old.UpsStatus.GetFlags()
	for flagName := range StatusFlags {
		if flags[flagName] != prevFlags[flagName] {
			curr.UpsStatus.FlagChangeCounts[flagName]++
		}
	}
}
