package actions

import "go-rogue/src/lib/maps"

type LightTorchAction struct {
	name       string
	difficulty float32
}

func NewLightTorchAction(node *maps.Node) *LightTorchAction {
	return &LightTorchAction{
		name:       "Light Torch",
		difficulty: node.GetDifficulties()["Light Torch"],
	}
}

func (a *LightTorchAction) GetName() string {
	return a.name
}

func (a *LightTorchAction) Execute() {
	// Implement the logic to perform a light torch action
}

func (a *LightTorchAction) GetText() string {
	return "Light torch action executed"
}
