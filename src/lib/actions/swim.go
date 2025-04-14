package actions

import "go-rogue/src/lib/maps"

type SwimAction struct {
	name       string
	difficulty int
}

func NewSwimAction(edge *maps.Edge) *SwimAction {
	return &SwimAction{
		name:       "Swim",
		difficulty: edge.GetId(),
	}
}

func (a *SwimAction) GetName() string {
	return a.name
}

func (a *SwimAction) Execute() {
	// Implement the logic to perform a swim action
}

func (a *SwimAction) GetText() string {
	return "Swim action executed"
}
