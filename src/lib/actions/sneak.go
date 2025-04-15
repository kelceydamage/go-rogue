package actions

import "go-rogue/src/lib/maps"

type SneakAction struct {
	name       string
	difficulty int
}

func NewSneakAction(edge *maps.Edge) *SneakAction {
	return &SneakAction{
		name:       "Sneak",
		difficulty: edge.GetDifficulty(),
	}
}

func (a *SneakAction) GetName() string {
	return a.name
}

func (a *SneakAction) Execute() {
	// Implement the logic to perform a sneak action
}

func (a *SneakAction) GetText() string {
	return "Sneak action executed"
}
