package engine

import (
	"fmt"
	"slices"
)

type InputProcessor struct{}

func (i *InputProcessor) GetValidInput(validInputs []string, currentLine, offset int) string {
	var input string
	currentLine++
	for {
		fmt.Printf("\033[47m\033[30m\033[%d;%dHSelect an option:\033[0m ", currentLine, offset)
		_, err := fmt.Scanln(&input)
		if err == nil {
			if slices.Contains(validInputs, input) {
				return input
			}
		}
	}
}

func (i *InputProcessor) IntRangeFromLength(length int) []int {
	result := make([]int, length)

	for i := range length {
		result[i] = i + 1
	}

	return result
}

func (i *InputProcessor) StringRangeFromLength(length int) []string {
	result := make([]string, length)

	for i := range length {
		result[i] = fmt.Sprintf("%d", i+1)
	}

	return result
}

func IntSliceToStringSlice(intSlice []int) []string {
	stringSlice := make([]string, len(intSlice))

	for i, v := range intSlice {
		stringSlice[i] = fmt.Sprintf("%d", v) // Convert int to string
	}

	return stringSlice
}
