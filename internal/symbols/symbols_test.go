package symbols

import (
	"strings"
	"testing"
)

func TestSymbolsOnly(t *testing.T) {
	input := "this is a simple input. Which, contains!\n a lot of $sp chars"
	expectedOutput := "thisisasimpleinput.Which,contains!alotof$spchars"

	output := SymbolsOnly(input)

	if strings.Contains(output, " ") || strings.Contains(output, "\n") {
		t.Log("Output contains whitespace or newline chars")
		t.Fail()
	}

	if expectedOutput != output {
		t.Logf("expected: %s\n", expectedOutput)
		t.Logf("actual: %s\n", output)
		t.Fail()
	}
}
