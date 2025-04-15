package actions

import (
	"go-rogue/src/lib/maps"
)

type BashAction struct {
	name       string
	difficulty int
}

func NewBashAction(edge *maps.Edge) *BashAction {
	return &BashAction{
		name:       "Bash",
		difficulty: edge.GetDifficulty(),
	}
}

func (a *BashAction) GetName() string {
	return a.name
}

func (a *BashAction) Execute() {
	// Implement the logic to perform a bash action
}

func (a *BashAction) GetText() string {
	return "Bash action executed"
}
