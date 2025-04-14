package userInterface

import (
	"fmt"
	"go-rogue/src/lib/config"
)

func DrawEventScreen(text string, draw config.CombatScreenSettings, currentLine int) int {
	// Wrap and print the event text message
	eventLines := wrapText(text, 100)
	for _, line := range eventLines {
		fmt.Printf("\033[%d;%dH%s\n", currentLine, draw.Offset, line)
		currentLine++
	}
	return currentLine
}
