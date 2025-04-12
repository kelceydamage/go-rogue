package entities

type World struct {
	currenZone int
	zones      map[int]*Zone
}

func NewWorld() *World {
	return &World{
		currenZone: 0,
		zones:      make(map[int]*Zone),
	}
}

func (w *World) AddZone(zoneId int, seed int64, exitNode, exitZoneId int, forwardTraversal bool) {
	w.zones[zoneId] = NewZone(zoneId, seed)
	if forwardTraversal {
		w.zones[zoneId].AddLink(0, exitZoneId)
		w.zones[exitZoneId].AddLink(exitNode, zoneId)
	} else {
		w.zones[zoneId].AddLink(w.zones[zoneId].sceneGraph.GetTerminusNodeId(), exitZoneId)
		w.zones[exitZoneId].AddLink(exitNode, zoneId)
	}
}

func (w *World) GetZoneCount() int {
	return len(w.zones)
}

func (w *World) GetCurrentZone() *Zone {
	return w.zones[w.currenZone]
}

func (w *World) GetCurrentZoneId() int {
	return w.currenZone
}

func (w *World) SetCurrentZone(zoneId int) {
	w.currenZone = zoneId
}
