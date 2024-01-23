package linematch

import "strings"

func FindLineMismatch(left string, right string) int {
	if left == right {
		return -1
	}

	if left == "" || right == "" {
		return 1
	}

	return findLineMismatchOfTwoDifferentNonEmptyStrings(left, right)
}

func findLineMismatchOfTwoDifferentNonEmptyStrings(left string, right string) int {

	leftLines := splitStringIntoLines(left)
	rightLines := splitStringIntoLines(right)

	maxLeftIndex := len(leftLines) - 1
	maxRightIndex := len(rightLines) - 1

	index := 0
	lineNumber := 1
	for {

		if bothSlicesCanBeAccessed(index, maxLeftIndex, maxRightIndex) {

			if linesAreDifferent(index, leftLines, rightLines) {
				return lineNumber
			} else {
				index += 1
				lineNumber += 1
			}
		} else {
			// One string is a subtring of the other
			// The current line number is where the strings differ
			return lineNumber
		}
	}
}

func splitStringIntoLines(input string) []string {
	return strings.Split(input, "\n")
}

func bothSlicesCanBeAccessed(index int, maxLeftIndex int, maxRightIndex int) bool {
	return index <= maxLeftIndex && index <= maxRightIndex
}

func linesAreDifferent(index int, leftLines []string, rightLines []string) bool {
	return leftLines[index] != rightLines[index]
}
