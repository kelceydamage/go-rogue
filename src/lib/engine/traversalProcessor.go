package engine

import (
	"fmt"
	"go-rogue/src/lib/actions"
	"go-rogue/src/lib/components"
	"go-rogue/src/lib/config"
	"go-rogue/src/lib/entities"
	"go-rogue/src/lib/maps"
	"strconv"
)

type Directions struct {
	LUT map[string]string
}

func NewDirections() *Directions {
	return &Directions{
		LUT: map[string]string{
			"F": "To the front",
			"L": "To the left",
			"R": "To the right",
			"B": "Adjacent to the way you came from",
		},
	}
}

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

var directions = NewDirections()

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
	return tp.InputProcessor.GetValidInput(validOptions, currentLine, config.CombatScreenSettingsInstance.Offset)
}

func (tp *TraversalProcessor) Execute(currentNode *maps.Node, player *entities.Player, currentLine int) int {
	tp.SetTraversalOptions(currentNode, player)
	// TODO: There is a segfault going back to node 0
	// Execute the traversal logic here
	// This could involve moving the player, updating the game state, etc.
	currentLine = tp.DrawTraversalOptionScreen(currentNode, tp.GetTraversalOptions(), currentLine)
	currentLine = tp.DrawBackTrackingOptionScreen(currentNode, player, currentLine)

	input := tp.CaptureInput(tp.NavigationOptions.GetCardinalKeys(), currentLine)
	selectedEdge := currentNode.GetEdge(tp.GetNodeId(input))
	actions := tp.GetActions(selectedEdge)
	currentLine = DrawActionsScreen(actions, currentLine)
	input = tp.CaptureInput(tp.InputProcessor.StringRangeFromLength(len(actions)), currentLine)
	actionIndex, _ := strconv.Atoi(input)
	actionIndex--

	// Do Action

	currentLine = DrawActionResultsScreen(selectedEdge, &actions[actionIndex], currentLine)

	player.SetCurrentPosition(selectedEdge.GetId(currentNode.GetId()))
	return currentLine
}

func (tp *TraversalProcessor) GetActions(selectedEdge *maps.Edge) []actions.Action {
	return components.PresentEdgeActions(selectedEdge)
}

func (tp *TraversalProcessor) DrawTraversalOptionScreen(
	currentNode *maps.Node,
	movementOptions map[int]string,
	currentLine int,
) int {
	fmt.Println("342432424", movementOptions)
	for nodeId, direction := range movementOptions {
		edge := currentNode.GetEdge(nodeId)
		fmt.Printf(
			"\033[%d;%dH[%s][%d] %s: %s\n",
			currentLine,
			config.CombatScreenSettingsInstance.Offset,
			direction,
			nodeId,
			directions.LUT[direction],
			edge.GetPreviewText(),
		)
		currentLine++
	}
	return currentLine
}

func (tp *TraversalProcessor) DrawBackTrackingOptionScreen(currentNode *maps.Node, player *entities.Player, currentLine int) int {
	currentLine++
	if currentNode.GetId() != 0 {
		fmt.Printf(
			"\033[%d;%dH[%s][%d] %s: %s\n",
			currentLine,
			config.CombatScreenSettingsInstance.Offset,
			"U",
			player.GetPreviousPosition(),
			"Back the way you came",
			currentNode.GetEdge(player.GetPreviousPosition()).GetPreviewText(),
		)
	}
	currentLine++
	return currentLine
}

func DrawActionsScreen(
	actions []actions.Action,
	currentLine int,
) int {
	currentLine += 2
	fmt.Printf("\033[%d;%dHAvailable Actions: ", currentLine, config.CombatScreenSettingsInstance.Offset)
	currentLine++
	for i, action := range actions {
		currentLine++
		fmt.Printf("\033[%d;%dH[%d] %s\n", currentLine, config.CombatScreenSettingsInstance.Offset, i+1, action.GetName())
	}
	currentLine++
	return currentLine
}

func DrawActionResultsScreen(
	edge *maps.Edge,
	action *actions.Action,
	currentLine int,
) int {
	currentLine += 2
	fmt.Printf("\033[%d;%dH%s\n", currentLine, config.CombatScreenSettingsInstance.Offset, edge.GetText())
	return currentLine
}
