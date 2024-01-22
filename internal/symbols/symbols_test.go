package symbols

import "testing"

func TestSymbolsOnly(t *testing.T) {
  input := "this is a simple input. Which, contains!\n a lot of $sp chars"
  expectedOutput := "thisisasimpleinput.Which,contains!\nalotof$spchars"
  
  output := SymbolsOnly(input)
  
  if expectedOutput != output {
    t.Fail()
  }
}
