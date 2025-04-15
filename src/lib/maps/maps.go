package maps

import (
	"fmt"
	"go-rogue/src/lib/config"
	"go-rogue/src/lib/utilities"
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

type GraphGenerator struct {
	traversalTextLoader *utilities.TraversalTextLoader
	eventTextLoader     *utilities.EventTextLoader
	edgeMap             map[string]*Edge
}

func NewGraphGenerator() *GraphGenerator {
	return &GraphGenerator{
		traversalTextLoader: LoadTraversal(),
		eventTextLoader:     LoadEventText(),
		edgeMap:             make(map[string]*Edge),
	}
}

func (g *GraphGenerator) AddEdge(sceneGraph *SceneGraph, nodeA, nodeB int, edgeType EdgeType) *Edge {
	// Generate a consistent key for the edge
	edgeKey := generateEdgeKey(nodeA, nodeB)
	fmt.Println("edgeKey", edgeKey)

	// Check if the edge already exists
	if edge, exists := g.edgeMap[edgeKey]; exists {
		fmt.Println("Edge already exists:", edgeKey)
		return edge // Reuse the existing edge
	}

	// Create a new edge if it doesn't exist
	newEdge := NewEdge(
		edgeType,
		[]int{nodeA, nodeB},
		// TODO: Implement diffuculty calculation
		0,
		g.traversalTextLoader.GetPreview(sceneGraph.theme.Name, string(edgeType)),
		g.traversalTextLoader.GetText(sceneGraph.theme.Name, string(edgeType)),
	)

	// Add the edge to the map
	g.edgeMap[edgeKey] = newEdge

	// Add the edge to both nodes in the scene graph
	sceneGraph.GetNode(nodeA).AddEdge(newEdge)
	sceneGraph.GetNode(nodeB).AddEdge(newEdge)
	fmt.Println("Adding edge between nodes", nodeA, "and", nodeB)
	return newEdge
}

// Generates a random graph using a fixed seed, with random node count, edges, and cycles
func (g *GraphGenerator) GenerateRandomSceneGraph(seed int64, theme *Theme) *SceneGraph {
	if seed == 0 {
		seed = rand.Int63()
	}
	fmt.Println("Seed:", seed)
	rand.Seed(seed)
	sceneGraph := NewSceneGraph(theme)

	fmt.Println("PopulateSceneGraph")
	// Add a random number of unlinked nodes based on rule
	g.PopulateSceneGraph(sceneGraph)

	fmt.Println("GenerateDeadNodes")
	// Reserve a percentage of nodes to be dead-end nodes
	g.GenerateDeadNodes(sceneGraph)

	fmt.Println("GenerateEdges")
	// Randomly distribute edges betwen nodes based on rules
	g.GenerateEdges(sceneGraph, theme)

	fmt.Println("GenerateCycles")
	// Ensure a minimum amount of cycles exist in the graph
	g.GenerateCycles(sceneGraph, theme)

	fmt.Println("ConnectClusters")
	// Connect any clusters to the nearest node in the root cluster
	g.ConnectClusters(sceneGraph)

	g.ColorDeadEndNodes(sceneGraph)
	// TODO: Color nodes based on node type - path, dead-end, terminus, encounter, etc

	// TODO: Color edges based on passability - walkable, blocked, locked door, etc

	return sceneGraph
}

func (g *GraphGenerator) ColorDeadEndNodes(sceneGraph *SceneGraph) {
	for id, _ := range sceneGraph.GetAllNodes() {
		if id == 0 {
			continue
		}
		if sceneGraph.GetEdgeCount(id) == 1 && !sceneGraph.IsTerminusNode(id) {
			sceneGraph.GetNode(id).SetNodeType(DeadEndNode)
		}
	}
}

func (g *GraphGenerator) PopulateSceneGraph(sceneGraph *SceneGraph) {
	n := rand.Intn(config.SceneGraphSettingsInstance.MaxNodes) + config.SceneGraphSettingsInstance.MinNodes
	fmt.Println("NodeCount", n)
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
		sceneGraph.AddNode(
			i,
			nodeType,
			"",
			g.eventTextLoader.GetText(sceneGraph.theme.Name, string(nodeType)),
		)
	}

	sceneGraph.SetTerminusNode(n - 1)
	g.AddEdge(sceneGraph, n-2, n-1, LockedDoor)
}

func (g *GraphGenerator) GenerateDeadNodes(sceneGraph *SceneGraph) {
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
			edgeType := GetRandomEdgeType(sceneGraph.theme)
			g.AddEdge(sceneGraph, nodeId, attachmentNodeId, edgeType)
		}
	}
}

func (g *GraphGenerator) GenerateEdges(sceneGraph *SceneGraph, theme *Theme) {
	edgeCount := rand.Intn(sceneGraph.GetNodeCount() * (sceneGraph.GetNodeCount() - 1) / 2)
	fmt.Println("edgeCount", edgeCount)
	for range edgeCount {
		randomNodeIds := rand.Perm(sceneGraph.GetNodeCount())[:2]
		nodeA, nodeB := randomNodeIds[0], randomNodeIds[1]
		if sceneGraph.IsReservedNode(nodeA) || sceneGraph.IsReservedNode(nodeB) || nodeA == nodeB {
			continue
		}
		if sceneGraph.GetEdgeCount(randomNodeIds[0]) < config.SceneGraphSettingsInstance.MaxEdgesPerNode &&
			sceneGraph.GetEdgeCount(randomNodeIds[1]) < config.SceneGraphSettingsInstance.MaxEdgesPerNode &&
			sceneGraph.GetNodeDistance(randomNodeIds[0], randomNodeIds[1]) <= config.SceneGraphSettingsInstance.MaxDistanceForEdgeToForm {
			edgeType := GetRandomEdgeType(theme)
			g.AddEdge(sceneGraph, nodeA, nodeB, edgeType)
		}
	}
}

func (g *GraphGenerator) GenerateCycles(sceneGraph *SceneGraph, theme *Theme) {
	for i := range sceneGraph.GetNodeCount() {
		for j := i + 1; j < sceneGraph.GetNodeCount(); j++ {
			if sceneGraph.IsReservedNode(i) || sceneGraph.IsReservedNode(j) || sceneGraph.ContainsEdge(i, j) {
				continue
			}
			sceneGraph.ContainsEdge(i, j)
			if sceneGraph.GetNodeDistance(i, j) <= config.SceneGraphSettingsInstance.MaxDistanceForEdgeToForm &&
				rand.Float32() < config.SceneGraphSettingsInstance.ProbabilityOfCycles {
				edgeType := GetRandomEdgeType(theme)
				g.AddEdge(sceneGraph, i, j, edgeType)
			}
		}
	}
}

func (g *GraphGenerator) ConnectClusters(sceneGraph *SceneGraph) {
	graphPathSearch := NewGraphPathSearch(sceneGraph)
	for i := 1; i < sceneGraph.GetNodeCount(); i++ {
		if !graphPathSearch.IsPathToNodeZero(i) {
			nearestNode := g.findNearestConnectedNodeConnectedToZero(sceneGraph, graphPathSearch, i)
			fmt.Println("nearestNode", nearestNode)
			if nearestNode != -1 {
				fmt.Println("Connecting node", i, "to nearest node", nearestNode)
				g.AddEdge(sceneGraph, i, nearestNode, LockedDoor)
			}
		}
	}
}

func (g *GraphGenerator) findNearestConnectedNodeConnectedToZero(sceneGraph *SceneGraph, graphPathSearch *GraphPathSearch, nodeId int) int {
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

func LoadTraversal() *utilities.TraversalTextLoader {
	traversalTextLoader := utilities.NewTraversalTextLoader()
	err := traversalTextLoader.LoadFromFile("src/lib/text/traversal.json")
	if err != nil {
		panic(fmt.Sprintf("Error loading traversal text: %s", err))
	}
	return traversalTextLoader
}

func LoadEventText() *utilities.EventTextLoader {
	eventTextLoader := utilities.NewEventTextLoader()
	// Load event text from JSON file
	err := eventTextLoader.LoadFromFile("src/lib/text/adventure.json")
	if err != nil {
		panic(fmt.Sprintf("Error loading event text: %s", err))
	}
	return eventTextLoader
}

func generateEdgeKey(nodeA, nodeB int) string {
	if nodeA < nodeB {
		return fmt.Sprintf("%d-%d", nodeA, nodeB)
	}
	return fmt.Sprintf("%d-%d", nodeB, nodeA)
}
