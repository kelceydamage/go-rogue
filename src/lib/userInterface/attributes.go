package userInterface

import (
	"fmt"
	"go-rogue/src/lib/config"
	"go-rogue/src/lib/interfaces"
)

func DrawPlayerAttributes(player interfaces.IEntity) {
	rules := config.AttributesScreenSettingsInstance
	fmt.Printf(getAttributeHeaderTextOutput(rules))
	for index, key := range player.GetAttributeMapOrderedKeys() {
		fmt.Printf(getAttributeTextOutput(key, player.GetAttributes()[key], index+1, rules))
	}
	fmt.Printf(getAttributeFooterTextOutput(len(player.GetAttributeMapOrderedKeys()), rules))
}

func getPaddedField(fieldName string, fieldPadding int) string {
	return fmt.Sprintf("%s:%s", fieldName, Spaces(fieldPadding-len(fieldName)))
}

func getPaddedValue(fieldName string, value float32, rules *config.AttributesScreenSettings) string {
	return Spaces(rules.Width - 1 - len(fmt.Sprintf("%s %.2f", fieldName, value)))
}

func getAttributeTextOutput(attributeName string, value float32, line int, rules *config.AttributesScreenSettings) string {
	fieldName := getPaddedField(attributeName, rules.NamePadding)
	return fmt.Sprintf(
		"\033[%d;0H%s %.2f%s|\n",
		rules.StartLine+line,
		fieldName,
		value,
		getPaddedValue(fieldName, value, rules),
	)
}

func getAttributeHeaderTextOutput(rules *config.AttributesScreenSettings) string {
	return fmt.Sprintf(
		"\033[%d;0H\033[47m\033[30mPlayer Attributes:%s\033[0m\n",
		rules.StartLine,
		Spaces(rules.Width-len("Player Attributes:")),
	)
}

func getAttributeFooterTextOutput(line int, rules *config.AttributesScreenSettings) string {
	return fmt.Sprintf(
		"\033[%d;0H\033[47m\033[30m%s\033[0m\n",
		rules.StartLine+line+1,
		Spaces(rules.Width-len("")),
	)
}
