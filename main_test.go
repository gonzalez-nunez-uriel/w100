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

package main

import (
	"os"
	"testing"

	linematch "github.com/gonzalez-nunez-uriel/w100/internal/linemismatch"
	"github.com/gonzalez-nunez-uriel/w100/internal/symbols"
)

/*
For this kind of application, it is possible to have (input, output) pairs that are independent of
implementation.
*/

func TestFormatWidth100(t *testing.T) {
	t.Log("Start of test")
	validateTests(t)

	r := FormatWidth100("dummy string")
	t.Logf("End of test%s\n", r)
	t.Fail()
}

func validateTests(t *testing.T) {

	testCases, err := os.ReadDir("testdata")

	if err != nil {
		t.Log("testdata directory could not be read")
		t.Fail()
	}

	for _, file := range testCases {
		// For more information on naming conventions read TODO: EXPLAIN TESTING NAMING CONVENTIONS AND ARCHITECTURE
		inputFileName := "testdata/" + file.Name() + "/input.md"
		outputFileName := "testdata/" + file.Name() + "/output.md"
		verificationFileName := "testdata/" + file.Name() + "/verification.md"

		input, err := os.ReadFile(inputFileName)

		if err != nil {
			t.Logf("input.md file at %s could not be opened\n", inputFileName)
			t.Fail()
		}

		output, err := os.ReadFile(outputFileName)

		if err != nil {
			t.Logf("output.md file at %s could not be opened\n", outputFileName)
			t.Fail()
		}

		verification, err := os.ReadFile(verificationFileName)

		if err != nil {
			t.Logf("verification.md file at %s could not be opened\n", verificationFileName)
			t.Fail()
		}

		validateTestCase(t, file.Name(), string(input), string(output), string(verification))
	}
}

func validateTestCase(t *testing.T, testCase string, input string, output string, verification string) {

	if symbols.SymbolsOnly(input) != symbols.SymbolsOnly(output) {
		t.Logf("Error in %s: input and output files do not match symbol-wise", testCase)
		//findStowawayChar(t, input, output)
		t.Fail()
	}

	thereIsMismatch, lineNumber := linematch.FindLineMismatch(output, verification)
	if thereIsMismatch {
		t.Logf("Error in %s: output and verification files do not match. Mismatch at line %d", testCase, lineNumber)
		t.Fail()
	}
}
