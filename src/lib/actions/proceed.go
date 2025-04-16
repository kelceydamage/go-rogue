package actions

import "go-rogue/src/lib/maps"

type ProceedAction struct {
	name       string
	difficulty int
}

func NewProceedAction(edge *maps.Edge) *ProceedAction {
	return &ProceedAction{
		name:       "Proceed",
		difficulty: edge.GetDifficulty(),
	}
}

func (a *ProceedAction) GetName() string {
	return a.name
}

func (a *ProceedAction) Execute() {
	// Implement the logic to perform a bash action
}

func (a *ProceedAction) GetText() string {
	return "Bash action executed"
}
