package config

func NewAttributesScreenSettings() *AttributesScreenSettings {
	return &AttributesScreenSettings{
		StartLine:   3,
		Width:       28,
		NamePadding: 15,
	}
}

func NewHeaderSettings() *HeaderSettings {
	return &HeaderSettings{
		StartLine: 1,
		Width:     120,
	}
}

func NewCombatScreenSettings() *CombatScreenSettings {
	return &CombatScreenSettings{
		StartLine: 11,
		Width:     70,
		Offset:    50,
	}
}

func NewSceneGraphSettings() *sceneGraphSettings {
	return &sceneGraphSettings{
		MinNodes:                     10,
		MaxNodes:                     20,
		MaxEdgesPerNode:              4,
		MaxDistanceForEdgeToForm:     3,
		ProbabilityOfMoreThanOneEdge: 0.7,
		ProbabilityOfDeadEndNode:     0.15,
		ProbabilityOfCycles:          0.2,
	}
}

var CombatScreenSettingsInstance = NewCombatScreenSettings()
var HeaderSettingsInstance = NewHeaderSettings()
var AttributesScreenSettingsInstance = NewAttributesScreenSettings()
var SceneGraphSettingsInstance = NewSceneGraphSettings()
