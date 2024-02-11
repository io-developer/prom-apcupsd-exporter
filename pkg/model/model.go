package model

type Model struct {
	State         State
	PrevState     State
	ChangedFields map[string][]interface{}
}

func NewModel() *Model {
	return &Model{
		State:         NewState(),
		PrevState:     NewState(),
		ChangedFields: map[string][]interface{}{},
	}
}

func (m *Model) Update(newState State) {
	m.PrevState = m.State
	m.State = newState

	m.updateStatusCounts()
	m.updateTransferOnbatt()

	_, diff := m.PrevState.Compare(newState)
	m.ChangedFields = diff
}

func (m *Model) updateStatusCounts() {
	m.State.UpsStatus.FlagChangeCounts = m.PrevState.UpsStatus.CloneFlagChangeCounts()

	if m.State.UpsStatus.Flag != m.PrevState.UpsStatus.Flag {
		flags := m.State.UpsStatus.GetFlags()
		prevFlags := m.PrevState.UpsStatus.GetFlags()
		for flagName := range StatusFlags {
			if flags[flagName] != prevFlags[flagName] {
				m.State.UpsStatus.FlagChangeCounts[flagName]++
			}
		}
	}
}

func (m *Model) updateTransferOnbatt() {
	// Reveal hidden quick transfers on battery and back
	transDelta := int64(m.State.UpsTransferOnBatteryCount) -
		int64(m.PrevState.UpsTransferOnBatteryCount)

	if transDelta > 0 {
		minIncr := uint64(transDelta-1) * 2
		m.State.UpsStatus.FlagChangeCounts["online"] += minIncr
		m.State.UpsStatus.FlagChangeCounts["onbatt"] += minIncr

		flag := m.State.UpsStatus.Flag
		if flag&StatusFlags["online"] != 0 {
			m.State.UpsStatus.FlagChangeCounts["online"] += 2
		}
		if flag&StatusFlags["onbatt"] == 0 {
			m.State.UpsStatus.FlagChangeCounts["onbatt"] += 2
		}
	}
}
