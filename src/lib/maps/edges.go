package maps

import (
	"fmt"
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

type Edge struct {
	Name  EdgeType
	Width int
	Style EdgeStyle
	Color Colors
}

func NewTunnelEdge() *Edge {
	return &Edge{
		Name:  EdgeType(Tunnel),
		Width: 1,
		Style: EdgeStyle(Dashed),
		Color: Colors(Brown),
	}
}

func NewPathEdge() *Edge {
	return &Edge{
		Name:  EdgeType(Path),
		Width: 1,
		Style: EdgeStyle(Solid),
		Color: Colors(Gray),
	}
}

func NewUnlockedDoorEdge() *Edge {
	return &Edge{
		Name:  EdgeType(UnlockedDoor),
		Width: 1,
		Style: EdgeStyle(Solid),
		Color: Colors(DarkGreen),
	}
}

func NewLockedDoorEdge() *Edge {
	return &Edge{
		Name:  EdgeType(LockedDoor),
		Width: 4,
		Style: EdgeStyle(Bold),
		Color: Colors(DarkRed),
	}
}

func NewHiddenDoorEdge() *Edge {
	return &Edge{
		Name:  EdgeType(HiddenDoor),
		Width: 1,
		Style: EdgeStyle(Dotted),
		Color: Colors(LightGray),
	}
}

func NewCrossingEdge() *Edge {
	return &Edge{
		Name:  EdgeType(Crossing),
		Width: 1,
		Style: EdgeStyle(Dashed),
		Color: Colors(LightBlue),
	}
}

var EdgeTypes = map[EdgeType]*Edge{
	Tunnel:       NewTunnelEdge(),
	Path:         NewPathEdge(),
	UnlockedDoor: NewUnlockedDoorEdge(),
	LockedDoor:   NewLockedDoorEdge(),
	HiddenDoor:   NewHiddenDoorEdge(),
	Crossing:     NewCrossingEdge(),
}

type EdgeStyle string

const (
	Solid     EdgeStyle = "solid"
	Dotted    EdgeStyle = "dotted"
	Dashed    EdgeStyle = "dashed"
	Bold      EdgeStyle = "bold"
	Invisible EdgeStyle = "invis"
)

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
	n.edges[edgeId] = EdgeTypes[EdgeType(Path)]
}

func (e *Edges) GetAllEdges() generics.HashSet[int] {
	return e.edgeKeys
}

func (e *Edges) GetEdge(edgeId int) *Edge {
	return e.edges[edgeId]
}

func (e *Edges) SetEdgeType(edgeId int, edgeType EdgeType) bool {
	if e.edgeKeys.Contains(edgeId) {
		e.edges[edgeId] = EdgeTypes[EdgeType(edgeType)]
		f := EdgeTypes[EdgeType(edgeType)]
		fmt.Println("--2-v Set Edge Type EdgeId", f.Name, "EdgeType", f.Style)
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
