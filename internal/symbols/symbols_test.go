package symbols

import (
	"strings"
	"testing"
)

func TestSymbolsOnly(t *testing.T) {

  input := "\n(~)this is - a 10 simple 124input.\n\n Which, +contains!\n a lot of $sp% &chars"
  output := SymbolsOnly(input)
  
	if strings.Contains(output, " ") {
		t.Log("Output contains at least one whitespace character")
		t.Fail()
	}
	
	if strings.Contains(output, "\n") {
		t.Log("Output contains at least one newline character")
		t.Fail()
	}
}
