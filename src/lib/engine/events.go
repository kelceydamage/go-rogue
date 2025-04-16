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
	return e.DrawEvent(currentNode.GetText(), currentLine, config.General.Offset)
}

func (e *EventProcessor) DrawEvent(text string, currentLine int, offset int) int {
	wrappedText := utilities.WrapTextNoIndent(text, config.General.WordWrapWidth)
	for _, line := range wrappedText {
		fmt.Printf("\033[%d;%dH%s", currentLine, offset, line)
		currentLine++
	}
	currentLine++
	return currentLine
}
