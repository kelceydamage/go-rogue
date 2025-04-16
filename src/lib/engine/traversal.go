package engine

import (
	"fmt"
	"go-rogue/src/lib/actions"
	"go-rogue/src/lib/components"
	"go-rogue/src/lib/config"
	"go-rogue/src/lib/entities"
	"go-rogue/src/lib/maps"
	"go-rogue/src/lib/utilities"
	"strconv"
)

type DirectionText struct {
	LUT map[string]string
}

func NewDirections() *DirectionText {
	return &DirectionText{
		LUT: map[string]string{
			"F": "To the front",
			"L": "To the left",
			"R": "To the right",
			"B": "Adjacent to the way you came from",
		},
	}
}

var directionText = NewDirections()

type DirectionOptions struct {
	CardinalLUT        map[string]int
	EdgeLUT            map[int]string
	BackTrackingEdgeId int
}

func NewNavigationOptions(options []int, BackTrackingEdgeId int) *DirectionOptions {
	directions := &DirectionOptions{
		CardinalLUT:        make(map[string]int),
		EdgeLUT:            make(map[int]string),
		BackTrackingEdgeId: BackTrackingEdgeId,
	}
	cardinals := []string{"F", "L", "R", "B"}

	index := 0
	for _, edgeId := range options {
		if edgeId == directions.BackTrackingEdgeId {
			continue
		}
		directions.EdgeLUT[edgeId] = cardinals[index]
		directions.CardinalLUT[cardinals[index]] = edgeId
		index++
	}
	return directions
}

func (d *DirectionOptions) GetCardinalKeys() []string {
	keys := make([]string, 0, len(d.CardinalLUT))
	for key := range d.CardinalLUT {
		keys = append(keys, key)
	}
	return append(keys, "U")
}

func (d *DirectionOptions) SetBackTrackingEdgeId(edgeId int) {
	d.BackTrackingEdgeId = edgeId
}

func (d *DirectionOptions) GetBackTrackingEdgeId() int {
	return d.BackTrackingEdgeId
}

func (d *DirectionOptions) GetCardinal(node int) string {
	return d.EdgeLUT[node]
}

func (d *DirectionOptions) GetNodeId(cardinal string) int {
	if cardinal == "U" {
		return d.BackTrackingEdgeId
	}
	return d.CardinalLUT[cardinal]
}

type TraversalProcessor struct {
	NavigationOptions *DirectionOptions
	InputProcessor    *InputProcessor
	PreviousNodeId    int
}

func NewTraversalProcessor(
	inputProcessor *InputProcessor,
) *TraversalProcessor {
	TraversalProcessor := &TraversalProcessor{
		InputProcessor: inputProcessor,
	}
	return TraversalProcessor
}

func (tp *TraversalProcessor) Execute(currentNode *maps.Node, player *entities.Player, currentLine int) int {
	tp.SetTraversalOptions(currentNode, player)
	// TODO: There is a segfault going back to node 0
	// Execute the traversal logic here
	// This could involve moving the player, updating the game state, etc.
	//fmt.Println()
	//fmt.Println("Current Line:", currentLine)
	currentLine = tp.DrawTraversalOptionScreen(currentNode, tp.GetTraversalOptions(), currentLine)
	//fmt.Println("Current Line:", currentLine)
	currentLine = tp.DrawBackTrackingOptionScreen(currentNode, player, currentLine)
	//fmt.Println("Current Line:", currentLine)
	input := tp.CaptureInput(tp.NavigationOptions.GetCardinalKeys(), currentLine)
	selectedEdge := currentNode.GetEdge(tp.GetNodeId(input))
	actions := tp.GetActions(selectedEdge)
	currentLine = tp.DrawActionsScreen(actions, currentLine)
	//fmt.Println("Current Line:", currentLine)
	input = tp.CaptureInput(tp.InputProcessor.StringRangeFromLength(len(actions)), currentLine)
	actionIndex, _ := strconv.Atoi(input)
	actionIndex--

	// Do Action

	currentLine = tp.DrawActionResultsScreen(selectedEdge, &actions[actionIndex], currentLine)
	//fmt.Println("Current Line:", currentLine)

	player.SetCurrentPosition(selectedEdge.GetId(currentNode.GetId()))
	return currentLine
}

func (tp *TraversalProcessor) GetCardinal(node int) string {
	return tp.NavigationOptions.GetCardinal(node)
}

func (tp *TraversalProcessor) GetNodeId(cardinal string) int {
	return tp.NavigationOptions.GetNodeId(cardinal)
}

func (tp *TraversalProcessor) SetTraversalOptions(node *maps.Node, player *entities.Player) {
	tp.NavigationOptions = NewNavigationOptions(node.GetAllEdges().Keys(), player.GetPreviousPosition())
}

func (tp *TraversalProcessor) GetTraversalOptions() map[int]string {
	return tp.NavigationOptions.EdgeLUT
}

func (tp *TraversalProcessor) GetPreviousPosition(player *entities.Player) int {
	return player.GetPreviousPosition()
}

func (tp *TraversalProcessor) CaptureInput(validOptions []string, currentLine int) string {
	return tp.InputProcessor.GetValidInput(validOptions, currentLine, config.General.Offset)
}

func (tp *TraversalProcessor) GetActions(selectedEdge *maps.Edge) []actions.Action {
	return components.PresentEdgeActions(selectedEdge)
}

func (tp *TraversalProcessor) DrawTraversalOptionScreen(
	currentNode *maps.Node,
	movementOptions map[int]string,
	currentLine int,
) int {
	for nodeId, direction := range movementOptions {
		edge := currentNode.GetEdge(nodeId)
		wrappedText := utilities.WrapText(
			fmt.Sprintf(
				"[%s][%d] %s: %s",
				direction,
				nodeId,
				directionText.LUT[direction],
				edge.GetPreviewText(),
			),
			config.General.WordWrapWidth,
		)
		for _, line := range wrappedText {
			fmt.Printf("\033[%d;%dH%s\n", currentLine, config.General.Offset, line)
			currentLine++
		}
	}
	return currentLine
}

func (tp *TraversalProcessor) DrawBackTrackingOptionScreen(currentNode *maps.Node, player *entities.Player, currentLine int) int {
	currentLine++
	if currentNode.GetId() != 0 {
		wrappedText := utilities.WrapText(
			fmt.Sprintf(
				"[%s][%d] Back the way you came: %s",
				"U",
				player.GetPreviousPosition(),
				currentNode.GetEdge(player.GetPreviousPosition()).GetPreviewText(),
			),
			config.General.WordWrapWidth,
		)
		for _, line := range wrappedText {
			fmt.Printf("\033[%d;%dH%s\n", currentLine, config.General.Offset, line)
			currentLine++
		}
	}
	currentLine++
	return currentLine
}

func (tp *TraversalProcessor) DrawActionsScreen(
	actions []actions.Action,
	currentLine int,
) int {
	currentLine += 3
	fmt.Printf("\033[%d;%dHAvailable Actions: ", currentLine, config.General.Offset)
	currentLine++
	for i, action := range actions {
		currentLine++
		fmt.Printf("\033[%d;%dH[%d] %s\n", currentLine, config.General.Offset, i+1, action.GetName())
	}
	currentLine++
	return currentLine
}

func (tp *TraversalProcessor) DrawActionResultsScreen(
	edge *maps.Edge,
	action *actions.Action,
	currentLine int,
) int {
	currentLine += 3
	fmt.Printf("\033[%d;%dH%s\n", currentLine, config.General.Offset, edge.GetText())
	return currentLine
}
