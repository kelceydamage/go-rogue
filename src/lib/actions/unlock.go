package actions

import "go-rogue/src/lib/maps"

type UnlockAction struct {
	name       string
	difficulty int
}

func NewUnlockAction(edge *maps.Edge) *UnlockAction {
	return &UnlockAction{
		name:       "Unlock",
		difficulty: edge.GetId(),
	}
}

func (a *UnlockAction) GetName() string {
	return a.name
}

func (a *UnlockAction) Execute() {
	// Implement the logic to unlock an action
}

func (a *UnlockAction) GetText() string {
	return "Action unlocked"
}
