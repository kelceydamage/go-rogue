package config

type AttributesScreenSettings struct {
	StartLine   int
	Width       int
	NamePadding int
}

type HeaderSettings struct {
	StartLine int
	Width     int
}

type CombatScreenSettings struct {
	StartLine int
	Width     int
	Offset    int
}
type sceneGraphSettings struct {
	MinNodes                     int
	MaxNodes                     int
	MaxEdgesPerNode              int
	MaxDistanceForEdgeToForm     int
	ProbabilityOfMoreThanOneEdge float32
	ProbabilityOfDeadEndNode     float32
	ProbabilityOfCycles          float32
}
