package maps

import (
	"go-rogue/src/lib/generics"
	"math"
)

type SceneGraph struct {
	nodes          map[int]*Node
	terminusNodeId int
	deadEndNodes   generics.HashSet[int]
}

func NewSceneGraph() *SceneGraph {
	return &SceneGraph{
		nodes:          make(map[int]*Node),
		terminusNodeId: -1,
		deadEndNodes:   generics.NewHashSet[int](),
	}
}

func (sg *SceneGraph) AddNode(nodeId int) {
	sg.nodes[nodeId] = NewNode(nodeId)
}

func (sg *SceneGraph) AddEdge(nodeIdA, nodeIdB int) {
	sg.nodes[nodeIdA].AddEdge(nodeIdB)
	sg.nodes[nodeIdB].AddEdge(nodeIdA)
}

func (sg *SceneGraph) SetTerminusNode(nodeId int) {
	sg.terminusNodeId = nodeId
	sg.nodes[nodeId].SetTerminusNode(true)
	sg.AddEdge(nodeId, nodeId-1)
}

func (sg *SceneGraph) IsTerminusNode(nodeId int) bool {
	return sg.nodes[nodeId].IsTerminusNode()
}

func (sg *SceneGraph) GetTerminusNodeId() int {
	return sg.terminusNodeId
}

func (sg *SceneGraph) SetDeadEndNodes(nodeIds []int) {
	for _, id := range nodeIds {
		sg.deadEndNodes.Add(id)
		sg.nodes[id].SetDeadNode(true)
		sg.nodes[id].ClearEdges()
	}
}

func (sg *SceneGraph) IsDeadEndNode(nodeId int) bool {
	return sg.nodes[nodeId].IsDeadEndNode()
}

func (sg *SceneGraph) GetDeadEndNodes() generics.HashSet[int] {
	return sg.deadEndNodes
}

func (sg *SceneGraph) GetNeighbors(nodeId int) generics.HashSet[int] {
	return sg.nodes[nodeId].GetEdges()
}

func (sg *SceneGraph) GetNodeCount() int {
	return len(sg.nodes)
}

func (sg *SceneGraph) GetEdgeCount(nodeId int) int {
	return sg.nodes[nodeId].GetEdgeCount()
}

func (sg *SceneGraph) GetNodeDistance(nodeIdA, nodeIdB int) int {
	return int(math.Abs(float64(nodeIdA - nodeIdB)))
}

func (sg *SceneGraph) GetAllNodes() map[int]*Node {
	return sg.nodes
}

func (sg *SceneGraph) GetNode(nodeId int) *Node {
	return sg.nodes[nodeId]
}

func (sg *SceneGraph) ContainsEdge(nodeId, edgeId int) bool {
	return sg.nodes[nodeId].GetEdges().Contains(edgeId)
}

func (sg *SceneGraph) IsReservedNode(nodeId int) bool {
	return sg.nodes[nodeId].IsDeadEndNode() || sg.nodes[nodeId].IsTerminusNode()
}

func (sg *SceneGraph) GetOrignId() int {
	return 0
}
