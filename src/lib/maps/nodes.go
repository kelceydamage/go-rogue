package maps

import (
	"go-rogue/src/lib/generics"
)

type NodeType string

const (
	StartNode     NodeType = "start"
	DecisionNode  NodeType = "decision"
	EncounterNode NodeType = "encounter"
	SceneryNode   NodeType = "scenery"
	EndingNode    NodeType = "ending"
	DeadEndNode   NodeType = "deadend"
)

type NodeMetaData struct {
	Name         NodeType
	Label        string
	IsDeadNode   bool
	Color        Colors
	IsResolvable bool
}

func NewStartNodeMetaData() *NodeMetaData {
	return &NodeMetaData{
		Name:         NodeType(StartNode),
		Label:        "Start",
		IsDeadNode:   false,
		Color:        Colors(Cyan),
		IsResolvable: false,
	}
}

func NewDecisionNodeMetaData() *NodeMetaData {
	return &NodeMetaData{
		Name:         NodeType(DecisionNode),
		Label:        "D",
		IsDeadNode:   false,
		Color:        Colors(Blue),
		IsResolvable: true,
	}
}

func NewEncounterNodeMetaData() *NodeMetaData {
	return &NodeMetaData{
		Name:         NodeType(EncounterNode),
		Label:        "E",
		IsDeadNode:   false,
		Color:        Colors(Red),
		IsResolvable: true,
	}
}

func NewSceneryNodeMetaData() *NodeMetaData {
	return &NodeMetaData{
		Name:         NodeType(SceneryNode),
		Label:        "_",
		IsDeadNode:   false,
		Color:        Colors(Green),
		IsResolvable: true,
	}
}

func NewEndingNodeMetaData() *NodeMetaData {
	return &NodeMetaData{
		Name:         NodeType(EndingNode),
		Label:        "End",
		IsDeadNode:   false,
		Color:        Colors(Purple),
		IsResolvable: false,
	}
}

func NewDeadEndNodeMetaData() *NodeMetaData {
	return &NodeMetaData{
		Name:         NodeType(DeadEndNode),
		Label:        "Dead",
		IsDeadNode:   true,
		Color:        Colors(Black),
		IsResolvable: false,
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
}

func NewNode(nodeId int, nodeType NodeType) *Node {
	return &Node{
		id:             nodeId,
		metaData:       NodeTypes[nodeType],
		isTerminusNode: false,
		edges:          NewEdges(),
		resolved:       false,
	}
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

func (n *Node) AddEdge(edgeId int, edgeType EdgeType) {
	n.edges.AddEdge(edgeId)
	n.edges.SetEdgeType(edgeId, edgeType)
}

func (n *Node) GetAllEdges() generics.HashSet[int] {
	return n.edges.GetAllEdges()
}

func (n *Node) GetEdge(attachmentNodeId int) *Edge {
	return n.edges.GetEdge(attachmentNodeId)
}

func (n *Node) SetEdgeType(attachmentNodeId int, edgeType EdgeType) bool {
	return n.edges.SetEdgeType(attachmentNodeId, edgeType)
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
