package linematch

import "testing"

type testCase struct {
	description        string
	left               string
	right              string
	expectedBool       bool
	expectedLineNumber int
}

var tests = []testCase{
	{"Two identical strings have no mismatched lines.",
		"Th3se TWO 5tings @re 1dentical$*!\n\nThe #^% #$() result should\nbe negative 0ne.\n",
		"Th3se TWO 5tings @re 1dentical$*!\n\nThe #^% #$() result should\nbe negative 0ne.\n",
		false,
		-1},
	{"If one of the strings is empty the two strings differ at line 1.",
		"0ne of th3se %trings is empty.\n\n The result should BE one.",
		"",
		true,
		1},
	{"Failed simple test case.",
		"These two strings\nhave a mismatch\nin the middle line.",
		"These two strings\nhave @ mismatch\nin the middle line.",
		true,
		2},
	{"Failed one char mismatch in the first line.",
		"Another two strings\nthat have a mismatch\nin the first line.",
		"Another tw* strings\nthat have a mismatch\nin the first line.",
		true,
		1},
	{"Failed one word mismatch at the last line.",
		"These strings\nhave an empty line\n\nand a mismatch\nin the last line.",
		"These strings\nhave an empty line\n\nand a mismatch\nin the end.",
		true,
		5},
	{"Failed multiple single char mismatches.",
		"These strings\nhave more than\n one mismatch.\n\nThe first is in\nthe third line.",
		"These strings\nhave more than\n 0ne mismatch.\n\nThe first i% in\nthe third line.",
		true,
		3},
	{"Failed multiple word mismatches.",
		"These strings\nhave more than\n one mismatch.\n\nThe first is in\nthe second line.",
		"These strings\nhave two\nmismatches.\n\nThe first is in\nthe second line",
		true,
		2},
	{"When two stings match line by line but one is longer than the other the next line number of the longer string is the expected value.",
		"These two\nstrings match char\nby char.",
		"These two\nstrings match char\nby char.\nBut one has\nmore lines.",
		true,
		4},
	{"Failed single lines without newlines with mismatching words.",
		"match",
		"mismatch",
		true,
		1},
	{"A single newline char difference at the end is a mismatch.",
		"first line\nsecond line\nfinal line",
		"first line\nsecond line\nfinal line\n",
		true,
		4},
	{"Failed multiple newline chars at the end of both strings.",
		"first line\nsecond line\nfinal line\n\n\n",
		"first line\nsecond line\nfinal line\n\n\n\n\n",
		true,
		7},
	{"Failed base edge case.",
		"\n",
		"\n\n",
		true,
		3},
}

func TestFindLineMismatch(t *testing.T) {

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			thereIsAMismatch, mismatchLineNumber := FindLineMismatch(testCase.left, testCase.right)
			if thereIsAMismatch != testCase.expectedBool || mismatchLineNumber != testCase.expectedLineNumber {
				t.Logf("%sExpected (%t, %d), got (%t, %d)", testCase.description, testCase.expectedBool, testCase.expectedLineNumber, thereIsAMismatch, mismatchLineNumber)
				t.Fail()
			}
		})
	}
}
