package actions

import "go-rogue/src/lib/maps"

type DisarmTrapAction struct {
	name       string
	difficulty int
}

func NewDisarmTrapAction(edge *maps.Edge) *DisarmTrapAction {
	return &DisarmTrapAction{
		name:       "Disarm Trap",
		difficulty: edge.GetDifficulty(),
	}
}

func (a *DisarmTrapAction) GetName() string {
	return a.name
}

func (a *DisarmTrapAction) Execute() {
	// Implement the logic to perform a detect trap action
}

func (a *DisarmTrapAction) GetText() string {
	return "Detect trap action executed"
}
