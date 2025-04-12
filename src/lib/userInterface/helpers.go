package userInterface

import "fmt"

func Spaces(count int) string {
	if count > 0 {
		return fmt.Sprintf("%*s", count, "")
	}
	return ""
}

func CenterText(text string, width int) string {
	padding := (width - len(text)) / 2
	return fmt.Sprintf("%s%s", Spaces(padding), text)
}
