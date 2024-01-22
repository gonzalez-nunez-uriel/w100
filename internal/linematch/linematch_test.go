package linematch

import "testing"

func TestFindLineMismatch(t *testing.T) {
	testOne(t)
}

func testOne(t *testing.T) {
	left := "This two strings\nhave a mismatch\nin the middle line."
	right := "This two strings\nhave @ mismatch\nin the middle line."
	//                               ^
	// mismatch here ----------------|

	mismatchLineNumber := FindLineMismatch(left, right)

	if mismatchLineNumber != 2 {
		t.Logf("Failed test one. Expected 2, got %d", mismatchLineNumber)
		t.Fail()
	}
}

func testTwo(t *testing.T) {
	left := "Another two strings\nthat have a mismatch\nin the first line."
	right := "Another tw* strings\nthat have a mismatch\nin the first line."
	//                  ^
	// mismatch here ---|

	mismatchLineNumber := FindLineMismatch(left, right)

	if mismatchLineNumber != 1 {
		t.Logf("Failed test one. Expected 1, got %d", mismatchLineNumber)
		t.Fail()
	}
}
