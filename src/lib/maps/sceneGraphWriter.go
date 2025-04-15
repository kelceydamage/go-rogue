package maps

import (
	"fmt"
	"os"
)

func WriteDotFile(filename string, sceneGraph *SceneGraph) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Start the DOT graph
	_, err = file.WriteString(fmt.Sprintf("graph G {\n  label=\"%s\";\n  labelloc=\"t\";\n  fontsize=\"20\";\n", sceneGraph.GetTheme().Name))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Write nodes with labels and colors
	for _, node := range sceneGraph.GetAllNodes() {
		nodeMetaData := node.GetMetaData()
		_, err = file.WriteString(fmt.Sprintf("  %d [label=\"%s\", color=\"%s\"];\n", node.GetId(), nodeMetaData.Label, nodeMetaData.Color))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	// Write edges
	for _, node := range sceneGraph.GetAllNodes() {
		for neighbor := range node.GetAllEdges() {
			if node.GetId() < neighbor { // Avoid duplicate edges
				edgeMetaData := node.GetEdge(neighbor).GetMetaData()
				_, err = file.WriteString(fmt.Sprintf("  %d -- %d [label=\"%s\", color=\"%s\", style=\"%s\", penwidth=\"%d\"];\n", node.GetId(), neighbor, edgeMetaData.Name, edgeMetaData.Color, edgeMetaData.Style, edgeMetaData.Width))
				if err != nil {
					fmt.Println("Error writing to file:", err)
					return
				}
			}
		}
	}

	// End the DOT graph
	_, err = file.WriteString("}\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Graph written to", filename)
}
