package actions

import "go-rogue/src/lib/maps"

type JumpAction struct {
	name       string
	difficulty int
}

func NewJumpAction(edge *maps.Edge) *JumpAction {
	return &JumpAction{
		name:       "Jump",
		difficulty: edge.GetDifficulty(),
	}
}

func (a *JumpAction) GetName() string {
	return a.name
}

func (a *JumpAction) Execute() {
	// Implement the logic to perform a jump action
}

func (a *JumpAction) GetText() string {
	return "Jump action executed"
}
