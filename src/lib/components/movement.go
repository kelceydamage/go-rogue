package components

import (
	"go-rogue/src/lib/generics"
	"go-rogue/src/lib/maps"
)

type Movement struct {
	// Current Node ID in graph
	currentPosition  int
	previousPosition int
}

func NewMovement(startPosition int) *Movement {
	return &Movement{
		currentPosition:  startPosition,
		previousPosition: 0,
	}
}

func (m *Movement) GetCurrentPosition() int {
	return m.currentPosition
}

func (m *Movement) GetPreviousPosition() int {
	return m.previousPosition
}

func (m *Movement) SetCurrentPosition(position int) {
	m.previousPosition = m.currentPosition
	m.currentPosition = position
}

func (m *Movement) GetMovementOptions(sceneGraph *maps.SceneGraph) generics.HashSet[int] {
	return sceneGraph.GetNeighbors(m.currentPosition)
}

func (m *Movement) Move(targetPosition int, sceneGraph map[int][]int) {
	m.SetCurrentPosition(targetPosition)
}
