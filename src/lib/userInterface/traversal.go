package userInterface

import (
	"fmt"
	"go-rogue/src/lib/config"
	"go-rogue/src/lib/entities"
	"go-rogue/src/lib/maps"
	"go-rogue/src/lib/utilities"
)

type Directions struct {
	LUT map[int]string
}

func NewDirections() *Directions {
	return &Directions{
		LUT: map[int]string{
			0: "To the front",
			1: "To the left",
			2: "To the right",
			3: "Adjacent to the way you came from",
		},
	}
}

var directions = NewDirections()

func DrawTraversalOptionScreen(sceneGraph *maps.SceneGraph, player *entities.Player, traversalTextLoader *utilities.TraversalTextLoader, draw config.CombatScreenSettings, currentLine int) (int, int) {
	currentLine++
	currentNode := sceneGraph.GetNode(player.GetCurrentPosition())
	textwrap := 60

	// Prepare movement options
	var movementOptions []int
	var previousOption string
	direction := 0

	for i := range currentNode.GetAllEdges() {
		if i == player.GetPreviousPosition() {
			// Buffer the option for the previous direction
			edge := currentNode.GetEdge(i)
			previousOption = fmt.Sprintf("[%d] Behind you, there is a [%s] where you came from.", i, edge.GetMetaData().Name)
		} else {
			// Buffer all other movement options
			movementOptions = append(movementOptions, i)
		}
		direction++
	}

	// Print movement options
	currentLine++
	fmt.Printf("\033[%d;%dHMovement Options:\n", currentLine, draw.Offset)

	// Print the buffered movement options with preview text
	for idx, option := range movementOptions {
		edge := currentNode.GetEdge(option)
		previewText := traversalTextLoader.GetPreview(sceneGraph.GetTheme().Name, string(edge.GetMetaData().Name))
		optionText := fmt.Sprintf("[%d] %s you see %s", option, directions.LUT[idx], previewText)
		optionLines := wrapText(optionText, textwrap)
		for _, line := range optionLines {
			currentLine++
			fmt.Printf("\033[%d;%dH%s\n", currentLine, draw.Offset, line)
		}
	}

	// Print the previous direction option first (if it exists)
	if previousOption != "" {
		previousLines := wrapText(previousOption, textwrap)
		for _, line := range previousLines {
			currentLine++
			fmt.Printf("\033[%d;%dH%s\n", currentLine, draw.Offset, line)
		}
	}

	// Print a prompt for the player to select an option
	currentLine += 2
	input := utilities.GetValidInput(currentNode.GetAllEdges().Keys(), currentLine, draw.Offset)
	return currentLine, input
}
