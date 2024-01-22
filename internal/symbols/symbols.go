package symbols

import "strings"

func SymbolsOnly(input string) string {
	withoutWhiteSpaces := strings.ReplaceAll(input, " ", "")
	withoutWhiteSpacesAndNewLines := strings.ReplaceAll(input, "\n", "")
	return withoutWhiteSpacesAndNewLines
}
