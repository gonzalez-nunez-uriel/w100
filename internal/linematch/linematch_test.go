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
}

// In case both strings are identical, the result should be -1
func testIdenticalStrings(t *testing.T) {
	left := "Th3se TWO 5tings @re 1dentical$*!\n\nThe #^% #$() result should\nbe negative 0ne.\n"
	right := "Th3se TWO 5tings @re 1dentical$*!\n\nThe #^% #$() result should\nbe negative 0ne.\n"

	mismatchLineNumber := FindLineMismatch(left, right)

	if mismatchLineNumber != -1 {
		t.Logf("For two identical strings the output should be -1, got %d", mismatchLineNumber)
		t.Fail()
	}
}

// If one of the strings is empty, the result should be 1
func testOneStringIsEmpty(t *testing.T) {
	left := "0ne of th3se %trings is empty.\n\n The result should BE one."
	right := ""

	mismatchLineNumber := FindLineMismatch(left, right)

	if mismatchLineNumber != 1 {
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

	mismatchLineNumber := FindLineMismatch(left, right)

	if mismatchLineNumber != 2 {
		t.Logf("Failed test one. Expected 2, got %d", mismatchLineNumber)
		t.Fail()
	}
}

// Mismatch at the first line
func testOneCharMismatchAtFirstLine(t *testing.T) {
	left := "Another two strings\nthat have a mismatch\nin the first line."
	right := "Another tw* strings\nthat have a mismatch\nin the first line."
	//                  ^
	// mismatch here ---|

	mismatchLineNumber := FindLineMismatch(left, right)

	if mismatchLineNumber != 1 {
		t.Logf("Failed test two. Expected 1, got %d", mismatchLineNumber)
		t.Fail()
	}
}

// More than one char mismatch
// Mismatch at the last line
func testOneWordMismatchAtLastLine(t *testing.T) {
	left := "These strings\nhave an empty line\n\nand a mistmatch\nin the last line."
	right := "These strings\nhave an empty line\n\nand a mistmatch\nin the end."
	//                                                                     ^
	// mismatch here ------------------------------------------------------|

	mismatchLineNumber := FindLineMismatch(left, right)

	if mismatchLineNumber != 4 {
		t.Logf("Failed test three. Expected 4, got %d", mismatchLineNumber)
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

	mismatchLineNumber := FindLineMismatch(left, right)

	if mismatchLineNumber != 3 {
		t.Logf("Failed test four. Expected 3, got %d", mismatchLineNumber)
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

	mismatchLineNumber := FindLineMismatch(left, right)

	if mismatchLineNumber != 2 {
		t.Logf("Failed test five. Expected 2, got %d", mismatchLineNumber)
		t.Fail()
	}
}
