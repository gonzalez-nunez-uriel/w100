package linematch

import "strings"

func FindLineMismatch(left string, right string) int {
	if left == right {
		return -1
	}

	if left == "" || right == "" {
		return 1
	}

	return findLineMismatchOfTwoDifferentNonEmptyStrings()
}

func splitStringIntoLines(input string) []string {
	return strings.Split(input, "\n")
}

func bothSlicesCanBeAccessed(index int, maxLeftIndex int, maxRightIndex int) bool {
	return index <= maxLeftIndex && index <= maxRightIndex
}

func findLineMismatchOfTwoDifferentNonEmptyStrings(left string, right string) int {
	leftLines := splitStringIntoLines(left)
	rightLines := splitStringIntoLines(right)

	maxLeftIndex := len(leftLines) - 1
	maxRightIndex := len(rightLines) - 1

	index := 0
	lineNumber := 1
	var currentLeftLine string
	var currentRightLine string
	for {

		if bothSlicesCanBeAccessed(index, maxLeftIndex, maxRightIndex) {
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
