package engine

import (
	"fmt"
	"go-rogue/src/lib/components"
	"go-rogue/src/lib/config"
	"go-rogue/src/lib/entities"
	"go-rogue/src/lib/maps"
	"go-rogue/src/lib/utilities"
)

type EventProcessor struct {
	TransitionLoader *utilities.TransitionLoader
	InputProcessor   *InputProcessor
}

func NewEventProcessor(inputProcessor *InputProcessor) *EventProcessor {
	return &EventProcessor{
		TransitionLoader: utilities.LoadTransitionText(),
		InputProcessor:   inputProcessor,
	}
}

func (e *EventProcessor) Execute(sceneGraph *maps.SceneGraph, player *entities.Player, currentLine int) int {
	currentNode := sceneGraph.GetNode(player.GetCurrentPosition())
	transitionText := ""
	if currentNode.GetId() != 0 {
		edge := currentNode.GetEdge(player.GetPreviousPosition())
		transitionText = e.TransitionLoader.GetTransition(
			sceneGraph.GetTheme().Name,
			string(edge.GetMetaData().Name),
			edge.GetLastUsedTraversalAction(),
			edge.GetScenarioId(),
		)
	}
	currentLine = e.DrawEvent(transitionText, currentNode.GetText(), currentLine, config.General.Offset)
	currentLine = e.DrawActions(currentNode, currentLine)
	input := e.CaptureInput(
		e.InputProcessor.StringRangeFromLength(
			len(components.PresentEventActions(currentNode)),
		),
		currentLine,
	)
	if input == "Search For Exists" {
		return currentLine
	}
	currentLine += 3
	return currentLine
}

func (e *EventProcessor) CaptureInput(validOptions []string, currentLine int) string {
	return e.InputProcessor.GetValidInput(validOptions, currentLine, config.General.Offset)
}

func (e *EventProcessor) DrawActions(currentNode *maps.Node, currentLine int) int {
	actions := components.PresentEventActions(currentNode)
	fmt.Printf("\033[%d;%dHAvailable Actions: ", currentLine, config.General.Offset)
	currentLine += 2
	for i, action := range actions {
		fmt.Printf("\033[%d;%dH[%d] %s\n", currentLine, config.General.Offset+5, i+1, action.GetName())
		currentLine++
	}
	return currentLine
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
