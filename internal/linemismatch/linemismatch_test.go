// MIT License
//
// Copyright (c) 2024 Uriel González Núñez
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the “Software”), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package linematch

import (
	"testing"

	"github.com/gonzalez-nunez-uriel/w100/internal/generators"
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
		// Should the case of both columns being equal be considered?
		// The function calling this one already makes that check
		// This function is called once it is known that two lines are different
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
		{"If one of the strings is empty, they differ at the first column",
			"mismatch",
			"",
			1},
		{"If one of the strings is empty, they differ at the first column",
			"",
			"mismatch",
			1},
		{"The strings can have many whitespaces in the beggining",
			"   mismatch",
			"   match",
			5},
		{"The strings can have many whitespaces in the beggining and differ by a whitespace",
			"   mismatch",
			"    match",
			4},
		{"The strings can have many whitespaces at the end and differ by a whitespace",
			"match  ",
			"match   ",
			8},
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

	numberOfProperty_basedTests := 10
	for i := 0; i < numberOfProperty_basedTests; i++ {
		t.Run("Property-based test", findColumnMismatchPropertyCheck)
	}
}

func findColumnMismatchPropertyCheck(t *testing.T) {
	// Idea from https://programming.guide/go/generate-random-character.html

	left, right := generators.CreateStringExamples()

	mismatchColumnNumber := FindColumnMismatch(left, right)

	if mismatchColumnNumber <= 0 {
		// this is enough information to replicate the error
		t.Log("The mismatch column must be a number greater than zero.")
		t.Logf("left:\n%s\n\n", left)
		t.Logf("right:\n%s\n\n", right)
		t.Logf("Reported mismatch: %d\n\n", mismatchColumnNumber)
		t.FailNow()
	}

	leftBeforeMismatch := left[0:mismatchColumnNumber]
	rightBeforeMismatch := right[0:mismatchColumnNumber]

	if stringsBeforeMismatchAreNotEqual(leftBeforeMismatch, rightBeforeMismatch) {
		// this is enough information to replicate the error
		t.Log("Strings before reported mismatch should be equal.")
		t.Logf("left:\n%s\n\n", left)
		t.Logf("right:\n%s\n\n", right)
		t.Logf("Reported mismatch: %d\n\n", mismatchColumnNumber)
		t.FailNow()
	}

}

func stringsBeforeMismatchAreNotEqual(leftBeforeMismatch string, rightBeforeMismatch string) bool {
	return leftBeforeMismatch != rightBeforeMismatch
}
