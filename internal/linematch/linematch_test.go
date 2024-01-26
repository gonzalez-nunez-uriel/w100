package linematch

import "testing"

func TestFindLineMismatch(t *testing.T) {
	testIdenticalStrings(t)
	testOneStringIsEmpty(t)
	simpleTest(t)
	testOneCharMismatchAtFirstLine(t)
	testOneWordMismatchAtLastLine(t)
	testMultipleSingleCharMismatches(t)
	testMultipleWordMismatches(t)
	testSingleLineWithoutNewlineMismatchingWords(t)
	testSameStringsButOneEndsWithNewlineChar(t)
	testSameStringsButOneEndsWithMoreNewlineChars(t)
	testEgdeCase(t)
}

// In case both strings are identical, the result should be -1
func testIdenticalStrings(t *testing.T) {
	left := "Th3se TWO 5tings @re 1dentical$*!\n\nThe #^% #$() result should\nbe negative 0ne.\n"
	right := "Th3se TWO 5tings @re 1dentical$*!\n\nThe #^% #$() result should\nbe negative 0ne.\n"

	thereIsAMismatch, mismatchLineNumber := FindLineMismatch(left, right)

	if !thereIsAMismatch && mismatchLineNumber != -1 {
		t.Logf("For two identical strings the output should be -1, got %d", mismatchLineNumber)
		t.Fail()
	}
}

// If one of the strings is empty, the result should be 1
func testOneStringIsEmpty(t *testing.T) {
	left := "0ne of th3se %trings is empty.\n\n The result should BE one."
	right := ""

	thereIsAMismatch, mismatchLineNumber := FindLineMismatch(left, right)

	if thereIsAMismatch && mismatchLineNumber != 1 {
		t.Logf("If one of the strings is empty the result should be 1, got %d", mismatchLineNumber)
		t.Fail()
	}
}

// Simple test
func simpleTest(t *testing.T) {
	left := "These two strings\nhave a mismatch\nin the middle line."
	right := "These two strings\nhave @ mismatch\nin the middle line."
	//                               ^
	// mismatch here ----------------|

	thereIsAMismatch, mismatchLineNumber := FindLineMismatch(left, right)

	if thereIsAMismatch && mismatchLineNumber != 2 {
		t.Logf("Failed simple test case. Expected 2, got %d", mismatchLineNumber)
		t.Fail()
	}
}

// Mismatch at the first line
func testOneCharMismatchAtFirstLine(t *testing.T) {
	left := "Another two strings\nthat have a mismatch\nin the first line."
	right := "Another tw* strings\nthat have a mismatch\nin the first line."
	//                  ^
	// mismatch here ---|

	thereIsAMismatch, mismatchLineNumber := FindLineMismatch(left, right)

	if thereIsAMismatch && mismatchLineNumber != 1 {
		t.Logf("Failed one char mismatch in the first line. Expected 1, got %d", mismatchLineNumber)
		t.Fail()
	}
}

// More than one char mismatch
// Mismatch at the last line
func testOneWordMismatchAtLastLine(t *testing.T) {
	left := "These strings\nhave an empty line\n\nand a mismatch\nin the last line."
	right := "These strings\nhave an empty line\n\nand a mismatch\nin the end."
	const expectedValue = 5
	//                                                                     ^
	// mismatch here ------------------------------------------------------|

	thereIsAMismatch, mismatchLineNumber := FindLineMismatch(left, right)

	if thereIsAMismatch && mismatchLineNumber != expectedValue {
		t.Logf("Failed one word mismatch at the last line. Expected %d, got %d", expectedValue, mismatchLineNumber)
		t.Fail()
	}
}

// Multiple single-char mismatches
// It should report the first mismatch
// The two mismatches are adjacent
func testMultipleSingleCharMismatches(t *testing.T) {
	left := "These strings\nhave more than\n one mismatch.\n\nThe first is in\nthe third line."
	right := "These strings\nhave more than\n 0ne mismatch.\n\nThe first i% in\nthe third line."
	//                                        ^                           ^
	// mismatches here -----------------------|---------------------------|

	thereIsAMismatch, mismatchLineNumber := FindLineMismatch(left, right)

	if thereIsAMismatch && mismatchLineNumber != 3 {
		t.Logf("Failed multiple single char mismatches. Expected 3, got %d", mismatchLineNumber)
		t.Fail()
	}
}

// Multiple word mismatches
// It should report the first mismatch
// The two mismatches are adjacent
func testMultipleWordMismatches(t *testing.T) {
	left := "These strings\nhave more than\n one mismatch.\n\nThe first is in\nthe second line."
	right := "These strings\nhave two\nmismatches.\n\nThe first is in\nthe second line"
	//                            ^    ^
	// mismatches here -----------|----|

	thereIsAMismatch, mismatchLineNumber := FindLineMismatch(left, right)

	if thereIsAMismatch && mismatchLineNumber != 2 {
		t.Logf("Failed multiple word mismatches. Expected 2, got %d", mismatchLineNumber)
		t.Fail()
	}
}

func testMatchingStringsButOneHasMoreLines(t *testing.T) {
	left := "These two\nstrings match char\nby char."
	right := "These two\nstrings match char\nby char.\nBut one has\nmore lines."
	// line       1             2               3      ^     4
	// mismatches here --------------------------------|

	thereIsAMismatch, mismatchLineNumber := FindLineMismatch(left, right)

	if thereIsAMismatch && mismatchLineNumber != 4 {
		t.Logf("When two stings match line by line but one is longer than the other the next line number of the longer string is the expected value. Expected 4, got %d", mismatchLineNumber)
		t.Fail()
	}
}

func testSingleLineWithoutNewlineMismatchingWords(t *testing.T) {
	left := "match"
	right := "mismatch"
	const expectedValue = 1

	thereIsAMismatch, mismatchLineNumber := FindLineMismatch(left, right)

	if thereIsAMismatch && mismatchLineNumber != expectedValue {
		t.Logf("Failed single lines without newlines with mismatching words. Expected %d, got %d", expectedValue, mismatchLineNumber)
		t.Fail()
	}
}

func testSameStringsButOneEndsWithNewlineChar(t *testing.T) {
	// lines 1         2            3            4
	left := "first line\nsecond line\nfinal line"
	right := "first line\nsecond line\nfinal line\n"
	//                                           ^
	// mismatch here ----------------------------|

	const expectedValue = 4

	thereIsAMismatch, mismatchLineNumber := FindLineMismatch(left, right)

	if thereIsAMismatch && mismatchLineNumber != expectedValue {
		t.Logf("A single newline char difference at the end of the strings is not a mismatch. Expected %d, got %d", expectedValue, mismatchLineNumber)
		t.Fail()
	}
}

func testSameStringsButOneEndsWithMoreNewlineChars(t *testing.T) {
	// lines       1           2           3    4 5 6 7 8
	left := "first line\nsecond line\nfinal line\n\n\n"
	right := "first line\nsecond line\nfinal line\n\n\n\n\n"
	//                                                 ^
	// mismatch here ----------------------------------|

	const expectedValue = 7

	thereIsAMismatch, mismatchLineNumber := FindLineMismatch(left, right)

	if thereIsAMismatch && mismatchLineNumber != expectedValue {
		t.Logf("Newline chars at the end of the file are not a mismatch. Expected %d, got %d", expectedValue, mismatchLineNumber)
		t.Fail()
	}
}

func testEgdeCase(t *testing.T) {
	//lines 1  2
	left := "\n"
	//lines  1  2 3
	right := "\n\n"
	// If you hesitate, you can confirm it with a text editor.
	// Press enter once and there are two lines
	// Press entre two times and there are three lines

	const expectedValue = 3

	thereIsAMismatch, mismatchLineNumber := FindLineMismatch(left, right)

	if thereIsAMismatch && mismatchLineNumber != expectedValue {
		t.Logf("Failed edge case. Expected %d, got %d", expectedValue, mismatchLineNumber)
		t.Fail()
	}
}
