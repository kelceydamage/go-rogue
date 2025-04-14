package actions

import (
	"go-rogue/src/lib/maps"
)

type InvestigateAction struct {
	name       string
	difficulty int
}

func NewInvestigateAction(edge *maps.Edge) *InvestigateAction {
	return &InvestigateAction{
		name:       "Investigate",
		difficulty: edge.GetId(),
	}
}

func (a *InvestigateAction) GetName() string {
	return a.name
}

func (a *InvestigateAction) Execute() {
	// Implement the logic to perform a search action
}

func (a *InvestigateAction) GetText() string {
	return "Search action executed"
}
