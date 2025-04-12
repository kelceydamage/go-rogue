package maps

import (
	"fmt"
	"go-rogue/src/lib/config"
	"math/rand"
)

// Generates a random graph using a fixed seed, with random node count, edges, and cycles
func GenerateRandomSceneGraph(seed int64) *SceneGraph {
	if seed == 0 {
		seed = rand.Int63()
	}
	fmt.Println("Seed:", seed)
	rand.Seed(seed)
	sceneGraph := NewSceneGraph()

	fmt.Println("PopulateSceneGraph")
	// Add a random number of unlinked nodes based on rule
	PopulateSceneGraph(sceneGraph)

	fmt.Println("GenerateDeadNodes")
	// Reserve a percentage of nodes to be dead-end nodes
	GenerateDeadNodes(sceneGraph)

	fmt.Println("GenerateEdges")
	// Randomly distribute edges betwen nodes based on rules
	GenerateEdges(sceneGraph)

	fmt.Println("GenerateCycles")
	// Ensure a minimum amount of cycles exist in the graph
	GenerateCycles(sceneGraph)

	fmt.Println("ConnectClusters")
	// Connect any clusters to the nearest node in the root cluster
	ConnectClusters(sceneGraph)

	// TODO: Color nodes based on node type - path, dead-end, terminus, encounter, etc

	// TODO: Color edges based on passability - walkable, blocked, locked door, etc

	return sceneGraph
}

func PopulateSceneGraph(sceneGraph *SceneGraph) {
	n := rand.Intn(config.SceneGraphSettingsInstance.MaxNodes) + config.SceneGraphSettingsInstance.MinNodes
	for i := range n {
		sceneGraph.AddNode(i)
	}
	sceneGraph.SetTerminusNode(n - 1)
}

func GenerateDeadNodes(sceneGraph *SceneGraph) {
	deadNodeCount := int(float32(sceneGraph.GetNodeCount()) * config.SceneGraphSettingsInstance.ProbabilityOfDeadEndNode)
	sceneGraph.SetDeadEndNodes(rand.Perm(sceneGraph.GetNodeCount() - 3)[:deadNodeCount])
	for nodeId := range sceneGraph.GetDeadEndNodes() {
		if nodeId == sceneGraph.GetTerminusNodeId() {
			continue
		}
		attachmentNodeId := nodeId
		for {
			attachmentNodeId = rand.Intn(sceneGraph.GetNodeCount())
			if sceneGraph.IsReservedNode(attachmentNodeId) {
				continue
			}
			break
		}
		if attachmentNodeId != nodeId {
			sceneGraph.AddEdge(nodeId, attachmentNodeId)
		}
	}
}

func GenerateEdges(sceneGraph *SceneGraph) {
	edgeCount := rand.Intn(sceneGraph.GetNodeCount() * (sceneGraph.GetNodeCount() - 1) / 2)
	for range edgeCount {
		randomNodeIds := rand.Perm(sceneGraph.GetNodeCount())[:2]
		if sceneGraph.IsReservedNode(randomNodeIds[0]) || sceneGraph.IsReservedNode(randomNodeIds[1]) {
			continue
		}
		if sceneGraph.GetEdgeCount(randomNodeIds[0]) < config.SceneGraphSettingsInstance.MaxEdgesPerNode &&
			sceneGraph.GetEdgeCount(randomNodeIds[1]) < config.SceneGraphSettingsInstance.MaxEdgesPerNode &&
			sceneGraph.GetNodeDistance(randomNodeIds[0], randomNodeIds[1]) <= config.SceneGraphSettingsInstance.MaxDistanceForEdgeToForm {
			sceneGraph.AddEdge(randomNodeIds[0], randomNodeIds[1])
		}
	}
}

func GenerateCycles(sceneGraph *SceneGraph) {
	for i := range sceneGraph.GetNodeCount() {
		for j := i + 1; j < sceneGraph.GetNodeCount(); j++ {
			if sceneGraph.IsReservedNode(i) || sceneGraph.IsReservedNode(j) || sceneGraph.ContainsEdge(i, j) {
				continue
			}
			sceneGraph.ContainsEdge(i, j)
			if sceneGraph.GetNodeDistance(i, j) <= config.SceneGraphSettingsInstance.MaxDistanceForEdgeToForm &&
				rand.Float32() < config.SceneGraphSettingsInstance.ProbabilityOfCycles {
				sceneGraph.AddEdge(i, j)
			}
		}
	}
}

func ConnectClusters(sceneGraph *SceneGraph) {
	graphPathSearch := NewGraphPathSearch(sceneGraph)
	for i := 1; i < sceneGraph.GetNodeCount(); i++ {
		if !graphPathSearch.IsPathToNodeZero(i) {
			nearestNode := findNearestConnectedNodeConnectedToZero(sceneGraph, graphPathSearch, i)
			if nearestNode != -1 {
				sceneGraph.AddEdge(i, nearestNode)
			}
		}
	}
}

func findNearestConnectedNodeConnectedToZero(sceneGraph *SceneGraph, graphPathSearch *GraphPathSearch, nodeId int) int {
	for i := nodeId - 1; i >= 0; i-- {
		if sceneGraph.IsReservedNode(i) {
			continue
		}
		if graphPathSearch.IsPathToNodeZero(i) {
			return i
		}
	}
	for i := nodeId + 1; i < sceneGraph.GetNodeCount(); i++ {
		if sceneGraph.IsReservedNode(i) {
			continue
		}
		if graphPathSearch.IsPathToNodeZero(i) {
			return i
		}
	}
	return -1
}
