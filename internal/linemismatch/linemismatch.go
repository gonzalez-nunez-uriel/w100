// Utility functions to find differences between strings on a line-to-line basis

// This package contains functions to find differences between strings. It will be used by the tests of w100
// to guarantee that test data is valid.
package linematch

import "strings"

func FindLineMismatch(left string, right string) (bool, int, int) {
	if left == right {
		return false, -1, -1
	}

	if left == "" || right == "" {
		return true, 1, 1
	}

	lineNumber, columnNumber := findLineMismatchOfTwoDifferentNonEmptyStrings(left, right)
	return true, lineNumber, columnNumber
}

func findLineMismatchOfTwoDifferentNonEmptyStrings(left string, right string) (int, int) {

	leftLines := splitStringIntoLines(left)
	rightLines := splitStringIntoLines(right)

	maxLeftIndex := len(leftLines) - 1
	maxRightIndex := len(rightLines) - 1

	index := 0
	lineNumber := 1
	for {

		if bothSlicesCanBeAccessed(index, maxLeftIndex, maxRightIndex) {

			if linesDoNotMatch(index, leftLines, rightLines) {
				return lineNumber, -10
			} else {
				index += 1
				lineNumber += 1
			}
		} else {
			// One string is a subtring of the other
			// The current line number is where the strings differ
			return lineNumber, 1
		}
	}
}

func splitStringIntoLines(input string) []string {
	return strings.Split(input, "\n")
}

func bothSlicesCanBeAccessed(index int, maxLeftIndex int, maxRightIndex int) bool {
	return index <= maxLeftIndex && index <= maxRightIndex
}

func linesDoNotMatch(index int, leftLines []string, rightLines []string) bool {
	return leftLines[index] != rightLines[index]
}

func FindColumnMismatch(left string, right string) int {
	return -10
}