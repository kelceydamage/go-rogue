package actions

import "go-rogue/src/lib/maps"

type DetectTrapAction struct {
	name       string
	difficulty int
}

func NewDetectTrapAction(edge *maps.Edge) *DetectTrapAction {
	return &DetectTrapAction{
		name:       "Detect Trap",
		difficulty: edge.GetId(),
	}
}

func (a *DetectTrapAction) GetName() string {
	return a.name
}

func (a *DetectTrapAction) Execute() {
	// Implement the logic to perform a detect trap action
}

func (a *DetectTrapAction) GetText() string {
	return "Detect trap action executed"
}
