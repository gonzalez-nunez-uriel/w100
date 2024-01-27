package linematch

import (
	"math/rand"
	"strings"
	"testing"
)

func TestFindLineMismatch(t *testing.T) {

	type testCase struct {
		description          string
		left                 string
		right                string
		expectedBool         bool
		expectedLineNumber   int
		expectedColumnNumber int
	}

	var tests = []testCase{
		{"Two identical strings have no mismatched lines.",
			"Th3se TWO 5tings @re 1dentical$*!\n\nThe #^% #$() result should\nbe negative 0ne.\n",
			"Th3se TWO 5tings @re 1dentical$*!\n\nThe #^% #$() result should\nbe negative 0ne.\n",
			false,
			-1,
			-1},
		{"If one of the strings is empty the two strings differ at line 1.",
			"0ne of th3se %trings is empty.\n\n The result should BE one.",
			"",
			true,
			1,
			1},
		{"Single char mismatch neither at the first nor last line.",
			"These two strings\nhave a mismatch\nin the middle line.",
			"These two strings\nhave @ mismatch\nin the middle line.",
			true,
			2,
			6},
		{"Single char mismatch in the first line.",
			"Another two strings\nthat have a mismatch\nin the first line.",
			"Another tw* strings\nthat have a mismatch\nin the first line.",
			true,
			1,
			11},
		{"One word mismatch at the last line.",
			"These strings\nhave an empty line\n\nand a mismatch\nin the last line.",
			"These strings\nhave an empty line\n\nand a mismatch\nin the end.",
			true,
			5,
			8},
		{"Multiple single char mismatches.",
			"These strings\nhave more than\n one mismatch.\n\nThe first is in\nthe third line.",
			"These strings\nhave more than\n 0ne mismatch.\n\nThe first i% in\nthe third line.",
			true,
			3,
			2},
		{"Multiple word mismatches.",
			"These strings\nhave more than\n one mismatch.\n\nThe first is in\nthe second line.",
			"These strings\nhave two\nmismatches.\n\nThe first is in\nthe second line",
			true,
			2,
			6},
		{"When two stings match line by line but one is longer than the other the next line number of the longer string is the expected value.",
			"These two\nstrings match char\nby char.",
			"These two\nstrings match char\nby char.\nBut one has\nmore lines.",
			true,
			4,
			1},
		{"Single mismatching lines without newline chars.",
			"match",
			"mismatch",
			true,
			1,
			2},
		{"A single newline char difference at the end is a mismatch.",
			"first line\nsecond line\nfinal line",
			"first line\nsecond line\nfinal line\n",
			true,
			4,
			1},
		{"Multiple newline chars at the end of both strings.",
			"first line\nsecond line\nfinal line\n\n\n",
			"first line\nsecond line\nfinal line\n\n\n\n\n",
			true,
			7,
			1},
		{"Base edge case.",
			"\n",
			"\n\n",
			true,
			3,
			1},
	}

	failCount := 0
	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			thereIsAMismatch, mismatchLineNumber, mismatchColumnNumber := FindLineMismatch(testCase.left, testCase.right)
			if thereIsAMismatch != testCase.expectedBool || mismatchLineNumber != testCase.expectedLineNumber || mismatchColumnNumber != testCase.expectedColumnNumber {
				t.Logf("%sExpected (%t, %d, %d), got (%t, %d, %d)", testCase.description, testCase.expectedBool, testCase.expectedLineNumber, testCase.expectedColumnNumber, thereIsAMismatch, mismatchLineNumber, mismatchColumnNumber)
				failCount += 1
				t.Fail()
			}
		})
	}

	percentFailed := float64(failCount) / float64(len(tests)) * 100.0
	t.Logf("\n\nTOTAL: %d FAILED: %d         PERCENT PASSED: %.1f%%         PERCENT FAILED: %.1f%%\n\n", len(tests), failCount, 100.0-percentFailed, percentFailed)
}

func TestFindColumnMismatch(t *testing.T) {
	type testCase struct {
		description          string
		left                 string
		right                string
		expectedColumnNumber int
	}

	var tests = []testCase{
		{"One word difference at the beggining",
			"match",
			"mismatch",
			2},
		{"Lowercase and uppercase are different",
			"mismatch",
			"Mismatch",
			1},
		{"Second word is differnt",
			"same match",
			"same mismatch",
			7},
		{"One word mismatch. Both have alphanumberics",
			"one # w0rd difference",
			"one # w0rd deviation",
			13},
		{"One numeric difference",
			"one num3ric difference",
			"one numeric difference",
			8},
		{"One symbol difference",
			"one $ymbol difference",
			"one symbol difference",
			5},
		{"One is a substring of the other",
			"same sentence",
			"same sentence mismatch",
			14},
		{"One is a substring of the other and ends with a number",
			"same sentenc3",
			"same sentenc3 mismatch",
			14},
		{"One is a substring of the other and ends with a symbol",
			"same sentenc&",
			"same sentenc& mismatch",
			14},
	}

	failCount := 0
	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			mismatchColumnNumber := FindColumnMismatch(testCase.left, testCase.right)
			if mismatchColumnNumber != testCase.expectedColumnNumber {
				t.Logf("%sExpected %d, got %d", testCase.description, testCase.expectedColumnNumber, mismatchColumnNumber)
				failCount += 1
				t.Fail()
			}
		})
	}

	percentFailed := float64(failCount) / float64(len(tests)) * 100.0
	t.Logf("\n\nTOTAL: %d FAILED: %d         PERCENT PASSED: %.1f%%         PERCENT FAILED: %.1f%%\n\n", len(tests), failCount, 100.0-percentFailed, percentFailed)

	findColumnMismatchPropertyCheck(t)
}

func findColumnMismatchPropertyCheck(t *testing.T) {
	// Idea from https://programming.guide/go/generate-random-character.html
	runeSet := []rune("abcdefghijklmnopqrstuvwsyz1234567890!@@#$%^&*().,")

	left, right := createStringExamples(runeSet)

	mismatchColumnNumber := FindColumnMismatch(left, right)

	leftBeforeMismatch := left[0:mismatchColumnNumber]
	rightBeforeMismatch := right[0:mismatchColumnNumber]

	if stringsBeforeMismatchAreNotEqual(leftBeforeMismatch, rightBeforeMismatch) {
		// this is enough information to replicate the error
		t.Log("Strings before reported mismatch should be equal.")
		t.Logf("left:\n%s\n\n", left)
		t.Logf("right:\n%s\n\n", right)
		t.Logf("Reported mismatch: %d\n\n", mismatchColumnNumber)
	}

}

func stringsBeforeMismatchAreNotEqual(leftBeforeMismatch string, rightBeforeMismatch string) bool {
	return leftBeforeMismatch != rightBeforeMismatch
}

func createStringExamples(runeSet []rune) (string, string) {
	mismatchAtThisWordCount := rand.Intn(10) + 1
	wordCountAfterMismatch := rand.Intn(10)

	var leftWords []string
	var rightWords []string
	for i := 0; i < mismatchAtThisWordCount-1; i++ {
		newWord := createWord(runeSet)
		leftWords = append(leftWords, newWord)
		rightWords = append(rightWords, newWord)
	}

	// chances are that these words are going to differ
	// in the first few runes. Is that an issue?
	leftWords = append(leftWords, createWord(runeSet))
	rightWords = append(rightWords, createWord(runeSet))

	// only a single mismatch is need
	// it might be a good idea to guarantee that there is only one difference per example
	// for debugging purposes
	for i := 0; i < wordCountAfterMismatch; i++ {
		newWord := createWord(runeSet)
		leftWords = append(leftWords, newWord)
		rightWords = append(rightWords, newWord)
	}

	// would the absence of whitespace in the beggining and the end be an issue?
	left := strings.Join(leftWords, " ")
	right := strings.Join(rightWords, " ")
	return left, right
}

func createWord(runeSet []rune) string {

	wordSize := rand.Intn(10) + 1
	var runes []rune
	for i := 0; i < wordSize; i++ {
		newRune := runeSet[rand.Intn(len(runeSet))]
		runes = append(runes, newRune)
	}

	// thanks https://yourbasic.org/golang/convert-string-to-rune-slice/
	// there is a chance for optimization
	return string(runes)
}
