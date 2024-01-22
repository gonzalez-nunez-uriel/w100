package main

import (
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
    // For more information on naming conventions read TODO: EXPLAIN TESTING NAMING CONVENTIONS AND ARCHITECTURE
    inputFileName := file.Name() + "/input.md"
    outputFileName := file.Name() + "/output.md"
    verificationFileName := file.Name() + "/verification.md"
		
		inputFile, err := os.ReadFile(inputFileName)
		
		if err != nil {
		  t.Error("input.md file at %s could not be opened", file.Name())
		}
		
		outputFile, err := os.ReadFile(outputFileName)
	}
	

	t.Fail()
}



func generateFileNames(t *testing.T, dirsOfCases []DirEntry) ([]string, []string, []string) {
  
}
