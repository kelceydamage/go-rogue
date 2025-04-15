package userInterface

import (
	"fmt"
	"go-rogue/src/lib/config"
	"go-rogue/src/lib/utilities"
)

func DrawTitleText(title string) {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	titelText := utilities.CenterText(title, config.Header.Width)
	fmt.Printf("\033[1;0H\033[1m\033[31m\033[40m%s%s\033[0m\n", titelText, utilities.Spaces(config.Header.Width-len(titelText)))
}
