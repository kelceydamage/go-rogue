package entities

import (
	"go-rogue/src/lib/maps"
)

type Zone struct {
	id         int
	sceneGraph *maps.SceneGraph
	links      map[int]int
}

func NewZone(zoneId int, sceneGraph *maps.SceneGraph) *Zone {
	return &Zone{
		id:         zoneId,
		sceneGraph: sceneGraph,
		links:      make(map[int]int),
	}
}

func (z *Zone) AddLink(exitNode, zoneId int) {
	z.links[exitNode] = zoneId
}

func (z *Zone) GetLink(nodeId int) (int, bool) {
	zoneId, exists := z.links[nodeId]
	return zoneId, exists
}

func (z *Zone) GetLinks() map[int]int {
	return z.links
}

func (z *Zone) GetSceneGraph() *maps.SceneGraph {
	return z.sceneGraph
}
