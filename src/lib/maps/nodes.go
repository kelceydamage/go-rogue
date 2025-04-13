package maps

import (
	"fmt"
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

type Node struct {
	id             int
	nodeType       NodeType
	isDeadNode     bool
	isTerminusNode bool
	edges          *Edges
	colors         Colors
}

func NewNode(nodeId int, nodeType NodeType) *Node {
	return &Node{
		id:             nodeId,
		nodeType:       nodeType,
		isDeadNode:     false,
		isTerminusNode: false,
		edges:          NewEdges(),
		colors:         Colors(Gray),
	}
}

func (n *Node) GetNodeType() NodeType {
	return n.nodeType
}

func (n *Node) SetNodeType(nodeType NodeType) {
	n.nodeType = nodeType
}

func (n *Node) SetDeadNode(isDeadNode bool) {
	n.isDeadNode = isDeadNode
}

func (n *Node) SetTerminusNode(isTerminusNode bool) {
	n.isTerminusNode = isTerminusNode
}

func (n *Node) IsDeadEndNode() bool {
	return n.isDeadNode
}

func (n *Node) IsTerminusNode() bool {
	return n.isTerminusNode
}

func (n *Node) GetId() int {
	return n.id
}

func (n *Node) AddEdge(edgeId int, edgeType EdgeType) {
	n.edges.AddEdge(edgeId)
	fmt.Println("Set Edge Type EdgeId", edgeId, "EdgeType", edgeType)
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
