package main

import (
	"github.com/gonzalez-nunez-uriel/w100/internal/symbols"
	"os"
	"testing"
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
		t.Log(file.Name())
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

	t.Fail()
}

func validateTestCase(t *testing.T, testCase string, input string, output string, verification string) {
	if symbols.SymbolsOnly(input) != symbols.SymbolsOnly(output) {
		t.Logf("Error in %s: input and output files do not match symbol-wise", testCase)
		t.Fail()
	}
	
	if output != verification {
	  t.Logf("Error in %s: output and verification files do not match", testCase)
		t.Fail()
	}
}
