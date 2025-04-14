package components

import (
	"fmt"
	"go-rogue/src/lib/interfaces"
	"time"
)

type Combat struct {
	TickRate float32
}

func NewCombat() *Combat {
	return &Combat{
		TickRate: 0.1, // Default tick rate
	}
}

func (c *Combat) Attack(player interfaces.IEntity, target interfaces.IEntity) {
	ticker := time.NewTicker(time.Duration(c.TickRate*1000) * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		// Perform attack logic
		damage := player.GetStrength() * 3
		if damage > 0 {
			target.SetHealth(target.GetHealth() - damage)
		}
	}
	fmt.Println()
}
