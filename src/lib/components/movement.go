package components

import (
	"go-rogue/src/lib/generics"
	"go-rogue/src/lib/maps"
)

type Movement struct {
	// Current Node ID in graph
	currentPosition int
}

func NewMovement(startPosition int) *Movement {
	return &Movement{
		currentPosition: startPosition,
	}
}

func (m *Movement) GetCurrentPosition() int {
	return m.currentPosition
}

func (m *Movement) SetCurrentPosition(position int) {
	m.currentPosition = position
}

func (m *Movement) GetMovementOptions(sceneGraph *maps.SceneGraph) generics.HashSet[int] {
	return sceneGraph.GetNeighbors(m.currentPosition)
}

func (m *Movement) Move(targetPosition int, sceneGraph map[int][]int) {
	m.SetCurrentPosition(targetPosition)
}
