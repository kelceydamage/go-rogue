package actions

import "go-rogue/src/lib/maps"

type SearchForExistsAction struct {
	name       string
	difficulty float32
}

func NewSearchForExistsAction(node *maps.Node) *SearchForExistsAction {
	return &SearchForExistsAction{
		name:       "Search For Exists",
		difficulty: 0,
	}
}

func (a *SearchForExistsAction) GetName() string {
	return a.name
}

func (a *SearchForExistsAction) Execute() {
	// Implement the logic to perform a light torch action
}

func (a *SearchForExistsAction) GetText() string {
	return "Light torch action executed"
}
