package actions

import "go-rogue/src/lib/maps"

type OpenAction struct {
	name       string
	difficulty int
}

func NewOpenAction(edge *maps.Edge) *OpenAction {
	return &OpenAction{
		name:       "Open Door",
		difficulty: edge.GetDifficulty(),
	}
}

func (a *OpenAction) GetName() string {
	return a.name
}

func (a *OpenAction) Execute() {
	// Implement the logic to perform a light torch action
}

func (a *OpenAction) GetText() string {
	return "Light torch action executed"
}
