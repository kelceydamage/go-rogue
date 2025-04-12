package userInterface

import (
	"fmt"
	"go-rogue/src/lib/config"
	"go-rogue/src/lib/interfaces"
)

func DrawCombatScreen(player interfaces.IEntity, target interfaces.IEntity) bool {
	var draw = config.CombatScreenSettingsInstance
	fmt.Printf("\033[%d;%dHCombat Started: Player vs Enemy\n", draw.StartLine, draw.Offset)
	fmt.Printf("\033[%d;%dHPlayer Health: %.2f", draw.StartLine+2, draw.Offset, player.GetHealth())
	fmt.Printf("\033[%d;%dHEnemy Health: %.2f", draw.StartLine+3, draw.Offset, target.GetHealth())
	if target.GetHealth() <= 0 {
		// Green background for "Enemy defeated!"
		fmt.Printf("\033[%d;%dH\033[42mEnemy defeated!%s\033[0m\n", draw.StartLine+5, draw.Offset, Spaces(30))
		return false
	}
	if player.GetHealth() <= 0 {
		// Red background for "You have been defeated!"
		fmt.Printf("\033[%d;%dH\033[41mYou have been defeated!%s\033[0m\n", draw.StartLine+5, draw.Offset, Spaces(30))
		return false
	}
	return true
}
