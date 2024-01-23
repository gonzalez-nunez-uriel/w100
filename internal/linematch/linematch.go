package linematch

import "strings"

func FindLineMismatch(left string, right string) int {
	if left == right {
		return -1
	}

	if left == "" || right == "" {
		return 1
	}

	// After this point both left and right must have at least one line

	leftLines := strings.Split(left, "\n")
	rightLines := strings.Split(right, "\n")

	maxLeftIndex := len(leftLines) - 1
	maxRightIndex := len(rightLines) - 1

	index := 0
	lineNumber := 1
	var currentLeftLine string
	var currentRightLine string
	for {
		canAccessLeftLines := index <= maxLeftIndex
		canAccessRightLines := index <= maxRightIndex

		if canAccessLeftLines && canAccessRightLines {
			currentLeftLine = leftLines[index]
			currentRightLine = rightLines[index]

			if currentLeftLine != currentRightLine {
				return lineNumber
			} else {
				index += 1
				lineNumber += 1
			}
		} else {
			// One string is longer than the other.
			// The current line number is where the strings differ
			return lineNumber
		}
	}
}
