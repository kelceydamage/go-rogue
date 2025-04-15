package actions

import "go-rogue/src/lib/maps"

type ClimbAction struct {
	name       string
	difficulty int
}

func NewClimbAction(edge *maps.Edge) *ClimbAction {
	return &ClimbAction{
		name:       "Climb",
		difficulty: edge.GetDifficulty(),
	}
}

func (a *ClimbAction) GetName() string {
	return a.name
}

func (a *ClimbAction) Execute() {
	// Implement the logic to perform a climb action
}

func (a *ClimbAction) GetText() string {
	return "Climb action executed"
}
