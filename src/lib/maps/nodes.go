package maps

import (
	"go-rogue/src/lib/generics"
)

type Node struct {
	id             int
	isDeadNode     bool
	isTerminusNode bool
	edges          generics.HashSet[int]
}

func NewNode(nodeId int) *Node {
	return &Node{
		id:             nodeId,
		isDeadNode:     false,
		isTerminusNode: false,
		edges:          generics.NewHashSet[int](),
	}
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

func (n *Node) AddEdge(edgeId int) {
	n.edges.Add(edgeId)
}

func (n *Node) GetEdges() generics.HashSet[int] {
	return n.edges
}

func (n *Node) GetEdgeCount() int {
	return n.edges.Size()
}

func (n *Node) ClearEdges() {
	n.edges.Clear()
}
