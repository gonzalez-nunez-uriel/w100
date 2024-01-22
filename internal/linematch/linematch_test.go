package linematch

import "testing"

func TestFindLineMismatch(t *testing.T) {
	testOne(t)
	testTwo(t)
	testThree(t)
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
