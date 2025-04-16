package engine

import (
	"fmt"
	"go-rogue/src/lib/config"
	"go-rogue/src/lib/entities"
	"go-rogue/src/lib/maps"
	"go-rogue/src/lib/utilities"
)

type EventProcessor struct{}

func NewEventProcessor() *EventProcessor {
	return &EventProcessor{}
}

func (e *EventProcessor) Execute(currentNode *maps.Node, player *entities.Player, currentLine int) int {
	transitionText := ""
	if currentNode.GetId() != 0 {
		edge := currentNode.GetEdge(player.GetPreviousPosition())
		transitionText = edge.GetTransitionText()
	}
	return e.DrawEvent(transitionText, currentNode.GetText(), currentLine, config.General.Offset)
}

func (e *EventProcessor) DrawEvent(transitionText, text string, currentLine, offset int) int {
	wrappedText := utilities.WrapTextNoIndent(transitionText, config.General.WordWrapWidth)
	for _, line := range wrappedText {
		fmt.Printf("\033[%d;%dH%s", currentLine, offset, line)
		currentLine++
	}
	currentLine++
	wrappedText = utilities.WrapTextNoIndent(text, config.General.WordWrapWidth)
	for _, line := range wrappedText {
		fmt.Printf("\033[%d;%dH%s", currentLine, offset, line)
		currentLine++
	}
	currentLine++
	return currentLine
}
