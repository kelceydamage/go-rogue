package maps

import (
	"fmt"
	"go-rogue/src/lib/config"
	"math/rand"
)

type Colors string

const (
	Red       Colors = "red"
	Green     Colors = "green"
	Blue      Colors = "blue"
	Black     Colors = "black"
	White     Colors = "white"
	Gray      Colors = "gray"
	Yellow    Colors = "yellow"
	Purple    Colors = "purple"
	Orange    Colors = "orange"
	Cyan      Colors = "cyan"
	Pink      Colors = "pink"
	Magenta   Colors = "magenta"
	Brown     Colors = "brown"
	LightBlue Colors = "lightblue"
	DarkGreen Colors = "darkgreen"
	LightGray Colors = "lightgray"
	DarkRed   Colors = "darkred"
)

// Generates a random graph using a fixed seed, with random node count, edges, and cycles
func GenerateRandomSceneGraph(seed int64, theme *Theme) *SceneGraph {
	if seed == 0 {
		seed = rand.Int63()
	}
	fmt.Println("Seed:", seed)
	rand.Seed(seed)
	sceneGraph := NewSceneGraph(theme)

	fmt.Println("PopulateSceneGraph")
	// Add a random number of unlinked nodes based on rule
	PopulateSceneGraph(sceneGraph)

	fmt.Println("GenerateDeadNodes")
	// Reserve a percentage of nodes to be dead-end nodes
	GenerateDeadNodes(sceneGraph)

	fmt.Println("GenerateEdges")
	// Randomly distribute edges betwen nodes based on rules
	GenerateEdges(sceneGraph, theme)

	fmt.Println("GenerateCycles")
	// Ensure a minimum amount of cycles exist in the graph
	GenerateCycles(sceneGraph, theme)

	fmt.Println("ConnectClusters")
	// Connect any clusters to the nearest node in the root cluster
	ConnectClusters(sceneGraph)

	ColorDeadEndNodes(sceneGraph)
	// TODO: Color nodes based on node type - path, dead-end, terminus, encounter, etc

	// TODO: Color edges based on passability - walkable, blocked, locked door, etc

	return sceneGraph
}

func ColorDeadEndNodes(sceneGraph *SceneGraph) {
	for id, _ := range sceneGraph.GetAllNodes() {
		if id == 0 {
			continue
		}
		if sceneGraph.GetEdgeCount(id) == 1 && !sceneGraph.IsTerminusNode(id) {
			sceneGraph.GetNode(id).SetNodeType(DeadEndNode)
		}
	}
}

func PopulateSceneGraph(sceneGraph *SceneGraph) {
	n := rand.Intn(config.SceneGraphSettingsInstance.MaxNodes) + config.SceneGraphSettingsInstance.MinNodes

	for i := range n {
		var nodeType NodeType

		switch {
		case i == 0:
			nodeType = StartNode // First node is the start
		case i == n-1:
			nodeType = EndingNode // Last node is the ending
		case rand.Float32() < 0.3:
			nodeType = DecisionNode // 30% chance of being a decision node
		case rand.Float32() < 0.5:
			nodeType = EncounterNode // 50% chance of being an encounter node
		default:
			nodeType = SceneryNode // Default to scenery
		}
		sceneGraph.AddNode(i, nodeType)
	}

	sceneGraph.SetTerminusNode(n - 1)
}

func GenerateDeadNodes(sceneGraph *SceneGraph) {
	deadNodeCount := int(float32(sceneGraph.GetNodeCount()) * config.SceneGraphSettingsInstance.ProbabilityOfDeadEndNode)
	sceneGraph.SetDeadEndNodes(rand.Perm(sceneGraph.GetNodeCount() - 3)[:deadNodeCount])
	fmt.Println("DeadEndNods", sceneGraph.GetDeadEndNodes())
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
			sceneGraph.AddEdge(nodeId, attachmentNodeId, EdgeType(Path))
		}
	}
}

func GenerateEdges(sceneGraph *SceneGraph, theme *Theme) {
	edgeCount := rand.Intn(sceneGraph.GetNodeCount() * (sceneGraph.GetNodeCount() - 1) / 2)
	for range edgeCount {
		randomNodeIds := rand.Perm(sceneGraph.GetNodeCount())[:2]
		if sceneGraph.IsReservedNode(randomNodeIds[0]) || sceneGraph.IsReservedNode(randomNodeIds[1]) {
			continue
		}
		if sceneGraph.GetEdgeCount(randomNodeIds[0]) < config.SceneGraphSettingsInstance.MaxEdgesPerNode &&
			sceneGraph.GetEdgeCount(randomNodeIds[1]) < config.SceneGraphSettingsInstance.MaxEdgesPerNode &&
			sceneGraph.GetNodeDistance(randomNodeIds[0], randomNodeIds[1]) <= config.SceneGraphSettingsInstance.MaxDistanceForEdgeToForm {
			edgeType := GetRandomEdgeType(theme)
			sceneGraph.AddEdge(randomNodeIds[0], randomNodeIds[1], edgeType)
		}
	}
}

func GenerateCycles(sceneGraph *SceneGraph, theme *Theme) {
	for i := range sceneGraph.GetNodeCount() {
		for j := i + 1; j < sceneGraph.GetNodeCount(); j++ {
			if sceneGraph.IsReservedNode(i) || sceneGraph.IsReservedNode(j) || sceneGraph.ContainsEdge(i, j) {
				continue
			}
			sceneGraph.ContainsEdge(i, j)
			if sceneGraph.GetNodeDistance(i, j) <= config.SceneGraphSettingsInstance.MaxDistanceForEdgeToForm &&
				rand.Float32() < config.SceneGraphSettingsInstance.ProbabilityOfCycles {
				edgeType := GetRandomEdgeType(theme)
				sceneGraph.AddEdge(i, j, edgeType)
				fmt.Println("Adding cycle between nodes", i, "and", j)
			}
		}
	}
}

func ConnectClusters(sceneGraph *SceneGraph) {
	graphPathSearch := NewGraphPathSearch(sceneGraph)
	for i := 1; i < sceneGraph.GetNodeCount(); i++ {
		if !graphPathSearch.IsPathToNodeZero(i) {
			nearestNode := findNearestConnectedNodeConnectedToZero(sceneGraph, graphPathSearch, i)
			fmt.Println("nearestNode", nearestNode)
			if nearestNode != -1 {
				fmt.Println("Connecting node", i, "to nearest node", nearestNode)
				sceneGraph.AddEdge(i, nearestNode, EdgeType(LockedDoor))
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

func GetNodeColor(nodeType NodeType) string {
	switch nodeType {
	case StartNode:
		return "cyan"
	case DecisionNode:
		return "blue"
	case EncounterNode:
		return "red"
	case SceneryNode:
		return "green"
	case EndingNode:
		return "purple"
	case DeadEndNode:
		return "black"
	default:
		return "gray"
	}
}
