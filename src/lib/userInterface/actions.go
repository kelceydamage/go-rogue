package userInterface

import (
	"fmt"
	"go-rogue/src/lib/components"
	"go-rogue/src/lib/config"
	"go-rogue/src/lib/maps"
	"go-rogue/src/lib/utilities"
)

func DrawActionsScreen(sceneGraph *maps.SceneGraph, edge *maps.Edge, traversalTextLoader *utilities.TraversalTextLoader, draw config.CombatScreenSettings, currentLine int) int {
	currentLine += 2
	actions := components.PresentEdgeActions(edge)
	fmt.Printf("\033[%d;%dHAvailable Actions: ", currentLine, draw.Offset)
	currentLine++
	for i, action := range actions {
		currentLine++
		fmt.Printf("\033[%d;%dH[%d] %s\n", currentLine, draw.Offset, i+1, action.GetName())
	}
	currentLine += 2
	utilities.GetValidInput(utilities.RangeFromLength(len(actions)), currentLine, draw.Offset)
	text := traversalTextLoader.GetText(
		sceneGraph.GetTheme().Name,
		string(edge.GetMetaData().Name),
	)
	currentLine += 2
	fmt.Printf("\033[%d;%dH%s\n", currentLine, draw.Offset, text)
	return currentLine
}
