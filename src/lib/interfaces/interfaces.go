package interfaces

type IEntity interface {
	GetAttributeMapOrderedKeys() []string
	GetAttributes() map[string]float32
	GetHealth() float32
	SetHealth(health float32)
	GetStrength() float32
	SetStrength(strength float32)
	GetAgility() float32
	SetAgility(agility float32)
	GetSpeed() float32
	SetSpeed(speed float32)
	GetIntelligence() float32
	SetIntelligence(intelligence float32)
	GetCharisma() float32
	SetCharisma(charisma float32)
}
