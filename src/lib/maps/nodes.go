package maps

import (
	"go-rogue/src/lib/generics"
	"math/rand"
)

type NodeType string

const (
	StartNode     NodeType = "Start"
	DecisionNode  NodeType = "Decision"
	EncounterNode NodeType = "Encounter"
	SceneryNode   NodeType = "Scenery"
	EndingNode    NodeType = "Ending"
	DeadEndNode   NodeType = "Deadend"
)

type NodeSubType string

const (
	Initial          NodeSubType = "Initial"
	Fork             NodeSubType = "Fork"
	Obstacle         NodeSubType = "Obstacle"
	Choice           NodeSubType = "Choice"
	Combat           NodeSubType = "Combat"
	Passive          NodeSubType = "Passive"
	Puzzle           NodeSubType = "Puzzle"
	Landmarks        NodeSubType = "Landmarks"
	PointsOfInterest NodeSubType = "PointsOfInterest"
	Blocked          NodeSubType = "Blocked"
	HiddenCache      NodeSubType = "HiddenCache"
)

type NodeSubtypes struct {
	LUT map[string][]NodeSubType
}

func NewNodeSubtypes() *NodeSubtypes {
	return &NodeSubtypes{
		LUT: map[string][]NodeSubType{
			"Start":     {Initial},
			"Ending":    {Initial},
			"Decision":  {Fork, Obstacle, Choice},
			"Encounter": {Combat, Passive, Puzzle},
			"Scenery":   {Landmarks, PointsOfInterest},
			"Deadend":   {Blocked, HiddenCache},
		},
	}
}

func (ns *NodeSubtypes) GetRandomSubtype(nodeType string, theme *Theme) string {
	possibleSubtypes := ns.LUT[nodeType]
	var total float32
	for _, sub := range possibleSubtypes {
		total += theme.NodeSubtypeProbabilities[sub]
	}
	r := rand.Float32() * total
	for _, sub := range possibleSubtypes {
		p := theme.NodeSubtypeProbabilities[sub]
		if r < p {
			return string(sub)
		}
		r -= p
	}
	return string(possibleSubtypes[len(possibleSubtypes)-1])
}

type NodeMetaData struct {
	Name         NodeType
	Label        string
	IsDeadNode   bool
	Color        Colors
	IsResolvable bool
	dificulties  map[string]float32
}

func NewStartNodeMetaData() *NodeMetaData {
	return &NodeMetaData{
		Name:         NodeType(StartNode),
		Label:        "Start",
		IsDeadNode:   false,
		Color:        Colors(Cyan),
		IsResolvable: false,
		dificulties: map[string]float32{
			"Investigate": 0.5,
			"Light Torch": 0.5,
		},
	}
}

func NewDecisionNodeMetaData() *NodeMetaData {
	return &NodeMetaData{
		Name:         NodeType(DecisionNode),
		Label:        "D",
		IsDeadNode:   false,
		Color:        Colors(Blue),
		IsResolvable: true,
		dificulties: map[string]float32{
			"Investigate": 0.5,
			"Light Torch": 0.5,
		},
	}
}

func NewEncounterNodeMetaData() *NodeMetaData {
	return &NodeMetaData{
		Name:         NodeType(EncounterNode),
		Label:        "E",
		IsDeadNode:   false,
		Color:        Colors(Red),
		IsResolvable: true,
		dificulties: map[string]float32{
			"Investigate": 0.5,
			"Light Torch": 0.5,
		},
	}
}

func NewSceneryNodeMetaData() *NodeMetaData {
	return &NodeMetaData{
		Name:         NodeType(SceneryNode),
		Label:        "_",
		IsDeadNode:   false,
		Color:        Colors(Green),
		IsResolvable: true,
		dificulties: map[string]float32{
			"Investigate": 0.5,
			"Light Torch": 0.5,
		},
	}
}

func NewEndingNodeMetaData() *NodeMetaData {
	return &NodeMetaData{
		Name:         NodeType(EndingNode),
		Label:        "End",
		IsDeadNode:   false,
		Color:        Colors(Purple),
		IsResolvable: false,
		dificulties: map[string]float32{
			"Investigate": 0.5,
			"Light Torch": 0.5,
		},
	}
}

func NewDeadEndNodeMetaData() *NodeMetaData {
	return &NodeMetaData{
		Name:         NodeType(DeadEndNode),
		Label:        "Dead",
		IsDeadNode:   true,
		Color:        Colors(Black),
		IsResolvable: false,
		dificulties: map[string]float32{
			"Investigate": 0.5,
			"Light Torch": 0.5,
		},
	}
}

var NodeTypes = map[NodeType]*NodeMetaData{
	StartNode:     NewStartNodeMetaData(),
	DecisionNode:  NewDecisionNodeMetaData(),
	EncounterNode: NewEncounterNodeMetaData(),
	SceneryNode:   NewSceneryNodeMetaData(),
	EndingNode:    NewEndingNodeMetaData(),
	DeadEndNode:   NewDeadEndNodeMetaData(),
}

type Node struct {
	id             int
	metaData       *NodeMetaData
	isTerminusNode bool
	edges          *Edges
	resolved       bool
	previewText    string
	text           string
	subtypes       string
}

func NewNode(nodeId int, nodeType NodeType, subtype, previewText, text string) *Node {
	return &Node{
		id:             nodeId,
		metaData:       NodeTypes[nodeType],
		isTerminusNode: false,
		edges:          NewEdges(nodeId),
		resolved:       false,
		previewText:    previewText,
		text:           text,
		subtypes:       subtype,
	}
}

func (n *Node) GetSubtype() string {
	return n.subtypes
}

func (n *Node) GetPreviewText() string {
	return n.previewText
}

func (n *Node) GetText() string {
	return n.text
}

func (n *Node) GetNodeType() NodeType {
	return n.metaData.Name
}

func (n *Node) SetNodeType(nodeType NodeType) {
	n.metaData = NodeTypes[nodeType]
}

func (n *Node) SetDeadNode(isDeadNode bool) {
	if !n.metaData.IsDeadNode {
		n.metaData = NodeTypes[DeadEndNode]
	}
}

func (n *Node) SetTerminusNode(isTerminusNode bool) {
	n.isTerminusNode = isTerminusNode
}

func (n *Node) IsDeadEndNode() bool {
	return n.metaData.IsDeadNode
}

func (n *Node) IsTerminusNode() bool {
	return n.isTerminusNode
}

func (n *Node) GetId() int {
	return n.id
}

func (n *Node) AddEdge(edge *Edge) {
	n.edges.AddEdge(edge)
}

func (n *Node) GetAllEdges() generics.HashSet[int] {
	return n.edges.GetAllEdges()
}

func (n *Node) GetEdge(attachmentNodeId int) *Edge {
	return n.edges.GetEdge(attachmentNodeId)
}

func (n *Node) GetEdgeCount() int {
	return n.edges.GetEdgeCount()
}

func (n *Node) ClearEdges() {
	n.edges.ClearEdges()
}

func (n *Node) GetMetaData() *NodeMetaData {
	return n.metaData
}

func (n *Node) SetMetaData(metaData *NodeMetaData) {
	n.metaData = metaData
}

func (n *Node) GetDifficulties() map[string]float32 {
	return n.metaData.dificulties
}
