package entities

import "go-rogue/src/lib/components"

type Enemy struct {
	*components.Attributes
}

func NewEnemy() *Enemy {
	return &Enemy{
		Attributes: components.NewAttributes(),
	}
}
