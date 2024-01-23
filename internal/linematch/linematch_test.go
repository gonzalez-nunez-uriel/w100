package linematch

import "testing"

func TestFindLineMismatch(t *testing.T) {
	testIdenticalStrings(t)
	testOne(t)
	testTwo(t)
	testThree(t)
	testFour(t)
	testFive(t)
}

// In case both strings are identical
func testIdenticalStrings(t *testing.T) {
  left := "Th3se TWO 5tings @re 1dentical$*!\n\nThe #^% #$() result should\nbe negative 0ne.\n"
  right := "Th3se TWO 5tings @re 1dentical$*!\n\nThe #^% #$() result should\nbe negative 0ne.\n"
  
  mismatchLineNumber := FindLineMismatch(left, right)

	if mismatchLineNumber != -1 {
		t.Logf("For two identical strings the output should be -1, got %d", mismatchLineNumber)
		t.Fail()
	}
}

// Simple test
func testOne(t *testing.T) {
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
func testTwo(t *testing.T) {
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
func testThree(t *testing.T) {
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
func testFour(t *testing.T) {
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
func testFive(t *testing.T) {
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

