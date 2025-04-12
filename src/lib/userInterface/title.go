package userInterface

import (
	"fmt"
	"go-rogue/src/lib/config"
)

func DrawTitleText(title string) {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	titelText := CenterText(title, config.HeaderSettingsInstance.Width)
	fmt.Printf("\033[1;0H\033[1m\033[31m\033[40m%s%s\033[0m\n", titelText, Spaces(config.HeaderSettingsInstance.Width-len(titelText)))
}
