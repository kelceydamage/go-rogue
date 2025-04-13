package maps

import (
	"go-rogue/src/lib/generics"
)

type EdgeType string

const (
	Tunnel       EdgeType = "tunnel"
	Path         EdgeType = "path"
	UnlockedDoor EdgeType = "unlockedDoor"
	LockedDoor   EdgeType = "lockedDoor"
	HiddenDoor   EdgeType = "hiddenDoor"
	Crossing     EdgeType = "crossing"
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
		Width:        1,
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
	metaData *EdgeMetaData
	resolved bool
}

func (e *Edge) GetMetaData() *EdgeMetaData {
	return e.metaData
}

func (e *Edge) SetMetaData(metaData *EdgeMetaData) {
	e.metaData = metaData
}

func NewEdge(edgeType EdgeType) *Edge {
	return &Edge{
		metaData: EdgeTypes[edgeType],
		resolved: false,
	}
}

type Edges struct {
	edgeKeys generics.HashSet[int]
	edges    map[int]*Edge
}

func NewEdges() *Edges {
	return &Edges{
		edgeKeys: generics.NewHashSet[int](),
		edges:    make(map[int]*Edge),
	}
}

func (n *Edges) AddEdge(edgeId int) {
	n.edgeKeys.Add(edgeId)
	n.edges[edgeId] = NewEdge(Path)
}

func (e *Edges) GetAllEdges() generics.HashSet[int] {
	return e.edgeKeys
}

func (e *Edges) GetEdge(edgeId int) *Edge {
	return e.edges[edgeId]
}

func (e *Edges) SetEdgeType(edgeId int, edgeType EdgeType) bool {
	if e.edgeKeys.Contains(edgeId) {
		e.edges[edgeId].SetMetaData(EdgeTypes[edgeType])
		return true
	}
	return false
}

func (e *Edges) GetEdgeCount() int {
	return e.edgeKeys.Size()
}

func (e *Edges) ClearEdges() {
	e.edgeKeys.Clear()
	e.edges = make(map[int]*Edge)
}
