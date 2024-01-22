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
}

func validateTests(t *testing.T) {

	testCases, err := os.ReadDir("testdata")

	if err != nil {
		t.Log("testdata directory could not be read")
	}

	// For the time being, just print the dir names
	for _, file := range testCases {
		t.Log(file.Name())
	}

	//for each folder in the list append the appropriate file names to open the right data

	//perform verifications on the contents of the files
}
