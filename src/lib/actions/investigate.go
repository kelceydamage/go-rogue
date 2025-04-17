package actions

import (
	"go-rogue/src/lib/maps"
)

type InvestigateAction struct {
	name       string
	difficulty float32
}

func NewInvestigateAction(node *maps.Node) *InvestigateAction {
	return &InvestigateAction{
		name:       "Investigate",
		difficulty: node.GetDifficulties()["Investigate"],
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
