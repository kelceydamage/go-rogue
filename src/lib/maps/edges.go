package maps

import (
	"go-rogue/src/lib/generics"
	"go-rogue/src/lib/utilities"
)

type EdgeType string

const (
	Tunnel       EdgeType = "Tunnel"
	Path         EdgeType = "Path"
	UnlockedDoor EdgeType = "UnlockedDoor"
	LockedDoor   EdgeType = "LockedDoor"
	HiddenDoor   EdgeType = "HiddenDoor"
	Crossing     EdgeType = "Crossing"
)

type EdgeStyle string

const (
	Solid     EdgeStyle = "solid"
	Dotted    EdgeStyle = "dotted"
	Dashed    EdgeStyle = "dashed"
	Bold      EdgeStyle = "bold"
	Invisible EdgeStyle = "invis"
)

type EdgeMetaData struct {
	Name         EdgeType
	Width        int
	Style        EdgeStyle
	Color        Colors
	IsResolvable bool
}

func NewTunnelEdgeMetaData() *EdgeMetaData {
	return &EdgeMetaData{
		Name:         EdgeType(Tunnel),
		Width:        1,
		Style:        EdgeStyle(Dashed),
		Color:        Colors(Brown),
		IsResolvable: false,
	}
}

func NewPathEdgeMetaData() *EdgeMetaData {
	return &EdgeMetaData{
		Name:         EdgeType(Path),
		Width:        1,
		Style:        EdgeStyle(Solid),
		Color:        Colors(Gray),
		IsResolvable: true,
	}
}

func NewUnlockedDoorEdgeMetaData() *EdgeMetaData {
	return &EdgeMetaData{
		Name:         EdgeType(UnlockedDoor),
		Width:        1,
		Style:        EdgeStyle(Solid),
		Color:        Colors(DarkGreen),
		IsResolvable: true,
	}
}

func NewLockedDoorEdgeMetaData() *EdgeMetaData {
	return &EdgeMetaData{
		Name:         EdgeType(LockedDoor),
		Width:        4,
		Style:        EdgeStyle(Bold),
		Color:        Colors(DarkRed),
		IsResolvable: true,
	}
}

func NewHiddenDoorEdgeMetaData() *EdgeMetaData {
	return &EdgeMetaData{
		Name:         EdgeType(HiddenDoor),
		Width:        1,
		Style:        EdgeStyle(Dotted),
		Color:        Colors(LightGray),
		IsResolvable: true,
	}
}

func NewCrossingEdgeMetaData() *EdgeMetaData {
	return &EdgeMetaData{
		Name:         EdgeType(Crossing),
		Width:        8,
		Style:        EdgeStyle(Dashed),
		Color:        Colors(LightBlue),
		IsResolvable: false,
	}
}

var EdgeTypes = map[EdgeType]*EdgeMetaData{
	Tunnel:       NewTunnelEdgeMetaData(),
	Path:         NewPathEdgeMetaData(),
	UnlockedDoor: NewUnlockedDoorEdgeMetaData(),
	LockedDoor:   NewLockedDoorEdgeMetaData(),
	HiddenDoor:   NewHiddenDoorEdgeMetaData(),
	Crossing:     NewCrossingEdgeMetaData(),
}

type Edge struct {
	metaData       *EdgeMetaData
	resolved       bool
	dificulty      int
	ids            []int
	previewText    string
	text           string
	transitionText string
}

func NewEdge(edgeType EdgeType, ids []int, difficulty int, textScenarios *utilities.EdgeTypeScenarios, scenario string) *Edge {
	return &Edge{
		metaData:       EdgeTypes[edgeType],
		resolved:       false,
		dificulty:      difficulty,
		ids:            ids,
		text:           textScenarios.Text[scenario],
		previewText:    textScenarios.Preview[scenario],
		transitionText: textScenarios.Transition[scenario],
	}
}

func (e *Edge) GetPreviewText() string {
	return e.previewText
}

func (e *Edge) GetText() string {
	return e.text
}

func (e *Edge) GetMetaData() *EdgeMetaData {
	return e.metaData
}

func (e *Edge) SetMetaData(metaData *EdgeMetaData) {
	e.metaData = metaData
}

func (n *Edge) GetDifficulty() int {
	return n.dificulty
}

func (e *Edge) GetId(currentNodeId int) int {
	for _, id := range e.ids {
		if id != currentNodeId {
			return id
		}
	}
	return currentNodeId
}

type Edges struct {
	edgeKeys generics.HashSet[int]
	edges    map[int]*Edge
	nodeId   int
}

func NewEdges(nodeId int) *Edges {
	return &Edges{
		edgeKeys: generics.NewHashSet[int](),
		edges:    make(map[int]*Edge),
		nodeId:   nodeId,
	}
}

func (n *Edges) AddEdge(edge *Edge) {
	n.edgeKeys.Add(edge.GetId(n.nodeId))
	n.edges[edge.GetId(n.nodeId)] = edge
}

func (e *Edges) GetAllEdges() generics.HashSet[int] {
	return e.edgeKeys
}

func (e *Edges) GetEdge(edgeId int) *Edge {
	return e.edges[edgeId]
}

func (e *Edges) GetEdgeCount() int {
	return e.edgeKeys.Size()
}

func (e *Edges) ClearEdges() {
	e.edgeKeys.Clear()
	e.edges = make(map[int]*Edge)
}
