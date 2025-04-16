package maps

import "math/rand"

type Theme struct {
	Name                     string
	EdgeTypeProbabilities    map[EdgeType]float32
	NodeSubtypeProbabilities map[NodeSubType]float32
}

var DungeonTheme = Theme{
	Name: "Dungeon",
	EdgeTypeProbabilities: map[EdgeType]float32{
		UnlockedDoor: 0.4,
		LockedDoor:   0.3,
		HiddenDoor:   0.2,
		Path:         0.1,
	},
	NodeSubtypeProbabilities: map[NodeSubType]float32{
		Initial:          1.0,
		Fork:             0.05,
		Obstacle:         0.45,
		Choice:           0.40,
		Combat:           0.5,
		Passive:          0.2,
		Puzzle:           0.3,
		Landmarks:        0.4,
		PointsOfInterest: 0.6,
		Blocked:          0.7,
		HiddenCache:      0.3,
	},
}

var ForestTheme = Theme{
	Name: "Forest",
	EdgeTypeProbabilities: map[EdgeType]float32{
		Path:     0.6,
		Crossing: 0.3,
		Tunnel:   0.1,
	},
	NodeSubtypeProbabilities: map[NodeSubType]float32{
		Initial:          1.0,
		Fork:             0.3,
		Obstacle:         0.5,
		Choice:           0.2,
		Combat:           0.6,
		Passive:          0.3,
		Puzzle:           0.1,
		Landmarks:        0.4,
		PointsOfInterest: 0.6,
		Blocked:          0.8,
		HiddenCache:      0.2,
	},
}

var CaveTheme = Theme{
	Name: "Cave",
	EdgeTypeProbabilities: map[EdgeType]float32{
		Tunnel:     0.7,
		Path:       0.2,
		HiddenDoor: 0.1,
	},
	NodeSubtypeProbabilities: map[NodeSubType]float32{
		Initial:          1.0,
		Fork:             0.5,
		Obstacle:         0.3,
		Choice:           0.2,
		Combat:           0.7,
		Passive:          0.1,
		Puzzle:           0.2,
		Landmarks:        0.2,
		PointsOfInterest: 0.8,
		Blocked:          0.8,
		HiddenCache:      0.2,
	},
}

func GetRandomEdgeType(theme *Theme) EdgeType {
	r := rand.Float32()
	accumulated := float32(0)

	for edgeType, probability := range theme.EdgeTypeProbabilities {
		accumulated += probability
		if r <= accumulated {
			return edgeType
		}
	}

	// Default to Path if no match (shouldn't happen if probabilities sum to 1)
	return Path
}

var ThemeLUT = map[int]*Theme{
	0: &DungeonTheme,
	1: &ForestTheme,
	2: &CaveTheme,
}
