package utilities

import (
	"fmt"
	"slices"
)

func GetValidInput(validInputs []int, currentLine, offset int) int {
	var input int

	for {
		fmt.Printf("\033[%d;%dHSelect an option: ", currentLine, offset)
		_, err := fmt.Scanln(&input)

		// Check if the input is valid
		if err == nil {
			if slices.Contains(validInputs, input) {
				return input // Return the valid input
			}
		}
	}
}

func RangeFromLength(length int) []int {
	result := make([]int, length)

	for i := 0; i < length; i++ {
		result[i] = i + 1
	}

	return result
}
