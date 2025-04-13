package maps

type GraphPathSearch struct {
	sceneGraph *SceneGraph
	visited    map[int]bool
}

// NewUnionFind initializes a UnionFind instance using the SceneGraph
func NewGraphPathSearch(sceneGraph *SceneGraph) *GraphPathSearch {
	return &GraphPathSearch{
		sceneGraph: sceneGraph,
		visited:    make(map[int]bool),
	}
}

func (uf *GraphPathSearch) IsPathToNodeZero(nodeId int) bool {
	uf.visited = make(map[int]bool)
	return uf.depthFirstSearch(nodeId, 0)
}

func (uf *GraphPathSearch) IsPathToNodeN(nodeId, targetNodeId int) bool {
	uf.visited = make(map[int]bool)
	return uf.depthFirstSearch(nodeId, targetNodeId)
}

// dfs performs a depth-first search to check if a node is connected to node 0
func (uf *GraphPathSearch) depthFirstSearch(current, target int) bool {
	if current == target {
		return true
	}
	uf.visited[current] = true
	for neighbor := range uf.sceneGraph.GetNode(current).GetAllEdges() {
		if !uf.visited[neighbor] {
			if uf.depthFirstSearch(neighbor, target) {
				return true
			}
		}
	}
	return false
}
