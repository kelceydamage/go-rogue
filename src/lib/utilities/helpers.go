package utilities

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

func WrapTextNoIndent(text string, maxWidth int) []string {
	var lines []string

	for len(text) > 0 {
		if len(text) <= maxWidth {
			lines = append(lines, text)
			break
		}

		split := maxWidth
		// Find the last space within the maxWidth limit
		for split > 0 && text[split] != ' ' {
			split--
		}
		if split == 0 { // No space found, force split
			split = maxWidth
		}

		lines = append(lines, text[:split])
		text = text[split:]
		// Remove leading space for the next line
		if len(text) > 0 && text[0] == ' ' {
			text = text[1:]
		}
	}

	return lines
}

func WrapText(text string, maxWidth int) []string {
	var lines []string
	indent := "    " // 4 spaces for indentation

	// Handle the first line without indentation
	if len(text) > maxWidth {
		// Find the last space within the maxWidth limit
		split := maxWidth
		for split > 0 && text[split] != ' ' {
			split--
		}
		if split == 0 { // No space found, force split
			split = maxWidth
		}

		// Add the first line without indentation
		lines = append(lines, text[:split])
		text = text[split:]
		if len(text) > 0 && text[0] == ' ' { // Remove leading space
			text = text[1:]
		}
	} else {
		// If the text fits within maxWidth, add it as a single line without indentation
		lines = append(lines, text)
		return lines
	}

	// Handle the remaining lines with indentation
	for len(text) > maxWidth {
		// Find the last space within the maxWidth limit
		split := maxWidth
		for split > 0 && text[split] != ' ' {
			split--
		}
		if split == 0 { // No space found, force split
			split = maxWidth
		}

		// Add the line with indentation and trim the text
		lines = append(lines, indent+text[:split])
		text = text[split:]
		if len(text) > 0 && text[0] == ' ' { // Remove leading space
			text = text[1:]
		}
	}

	// Add the remaining text with indentation (if any)
	if len(text) > 0 {
		lines = append(lines, indent+text)
	}

	return lines
}

func ClearScreenBelow(line int, offset int) {
	// Move the cursor to the specified line and offset
	fmt.Printf("\033[%d;%dH", line, offset)

	// Clear the screen from the cursor position downward
	fmt.Print("\033[J")
}
