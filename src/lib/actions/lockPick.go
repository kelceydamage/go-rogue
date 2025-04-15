package actions

import "go-rogue/src/lib/maps"

type LockPickAction struct {
	name       string
	difficulty int
}

func NewLockPickAction(edge *maps.Edge) *LockPickAction {
	return &LockPickAction{
		name:       "Pick Lock",
		difficulty: edge.GetDifficulty(),
	}
}

func (a *LockPickAction) GetName() string {
	return a.name
}

func (a *LockPickAction) Execute() {
	// Implement the logic to perform a lock pick action
}

func (a *LockPickAction) GetText() string {
	return "Lock pick action executed"
}
