package components

type Attributes struct {
	attributeMap           map[string]float32
	attibuteMapOrderedKeys []string
}

func NewAttributes() *Attributes {
	return &Attributes{
		attributeMap: map[string]float32{
			"Health":       100.0,
			"Strength":     10.0,
			"Agility":      10.0,
			"Speed":        10.0,
			"Intelligence": 10.0,
			"Charisma":     10.0,
		},
		attibuteMapOrderedKeys: []string{
			"Health",
			"Strength",
			"Agility",
			"Speed",
			"Intelligence",
			"Charisma",
		},
	}
}

func (a *Attributes) GetAttributeMapOrderedKeys() []string {
	return a.attibuteMapOrderedKeys
}

func (a *Attributes) GetAttributes() map[string]float32 {
	return a.attributeMap
}

func (a *Attributes) GetHealth() float32 {
	return a.GetAttributes()["Health"]
}

func (a *Attributes) SetHealth(health float32) {
	a.attributeMap["Health"] = health
}

func (a *Attributes) GetStrength() float32 {
	return a.GetAttributes()["Strength"]
}

func (a *Attributes) SetStrength(strength float32) {
	a.attributeMap["Strength"] = strength
}

func (a *Attributes) GetAgility() float32 {
	return a.GetAttributes()["Agility"]
}

func (a *Attributes) SetAgility(agility float32) {
	a.attributeMap["Agility"] = agility
}

func (a *Attributes) GetSpeed() float32 {
	return a.GetAttributes()["Speed"]
}

func (a *Attributes) SetSpeed(speed float32) {
	a.attributeMap["Speed"] = speed
}

func (a *Attributes) GetIntelligence() float32 {
	return a.GetAttributes()["Intelligence"]
}

func (a *Attributes) SetIntelligence(intelligence float32) {
	a.attributeMap["Intelligence"] = intelligence
}

func (a *Attributes) GetCharisma() float32 {
	return a.GetAttributes()["Charisma"]
}

func (a *Attributes) SetCharisma(charisma float32) {
	a.attributeMap["Charisma"] = charisma
}
