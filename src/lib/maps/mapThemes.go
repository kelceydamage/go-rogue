package maps

import "math/rand"

type Theme struct {
	Name                  string
	EdgeTypeProbabilities map[EdgeType]float32 // Probabilities for each edge type
}

var DungeonTheme = Theme{
	Name: "Dungeon",
	EdgeTypeProbabilities: map[EdgeType]float32{
		UnlockedDoor: 0.4,
		LockedDoor:   0.3,
		HiddenDoor:   0.2,
		Path:         0.1,
	},
}

var ForestTheme = Theme{
	Name: "Forest",
	EdgeTypeProbabilities: map[EdgeType]float32{
		Path:     0.6,
		Crossing: 0.3,
		Tunnel:   0.1,
	},
}

var CaveTheme = Theme{
	Name: "Cave",
	EdgeTypeProbabilities: map[EdgeType]float32{
		Tunnel:     0.7,
		Path:       0.2,
		HiddenDoor: 0.1,
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
