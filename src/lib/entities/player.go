package entities

import "go-rogue/src/lib/components"

type Player struct {
	*components.Attributes
	*components.Movement
}

func NewPlayer() *Player {
	return &Player{
		Attributes: components.NewAttributes(),
		Movement:   components.NewMovement(0),
	}
}
