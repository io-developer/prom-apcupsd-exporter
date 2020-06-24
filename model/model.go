package model

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

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

	if !cmp.Equal(m.PrevState, m.State) {
		// TODO calc ChangedFields
		fmt.Println("NOT EQUALLLLLLLLS: ", cmp.Equal(m.PrevState, m.State))
	}
}

func updateStatusCounts(old *State, curr *State) {
	curr.UpsStatus.FlagChangeCounts = old.UpsStatus.FlagChangeCounts
	flags := curr.UpsStatus.GetFlags()
	prevFlags := old.UpsStatus.GetFlags()
	for flagName := range StatusFlags {
		if flags[flagName] != prevFlags[flagName] {
			fmt.Println(" flag NOT EQ: ", flagName)
			curr.UpsStatus.FlagChangeCounts[flagName]++
		}
	}
}
