package actions

import "go-rogue/src/lib/maps"

type WalkAction struct {
	name       string
	difficulty int
}

func NewWalkAction(edge *maps.Edge) *WalkAction {
	return &WalkAction{
		name:       "Walk",
		difficulty: edge.GetDifficulty(),
	}
}

func (a *WalkAction) GetName() string {
	return a.name
}

func (a *WalkAction) Execute() {
	// Implement the logic to perform a bash action
}

func (a *WalkAction) GetText() string {
	return "Bash action executed"
}
