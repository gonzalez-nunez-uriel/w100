package symbols

import "testing"

func TestSymbolsOnly(t *testing.T) {
	input := "this is a simple input. Which, contains!\n a lot of $sp chars"
	expectedOutput := "thisisasimpleinput.Which,contains!\nalotof$spchars"

	output := SymbolsOnly(input)

	if expectedOutput != output {
		t.Fail()
	}

	if SymbolsOnly(inputDiscourseOnTheMethod()) != expectedOutputDiscourseOnTheMethod() {
		t.Fail()
	}
}

func inputDiscourseOnTheMethod() string {
	// excerpt from Discourse On the Method by Rene Descartes
	// taken from https://www.gutenberg.org/files/59/59-h/59-h.htm, PART I
	return "Good sense is, of all things among men, the most equally distributed; for every one thinks himself so abundantly provided with it, that those even who are the most difficult to satisfy in everything else, do not usually desire a larger measure of this quality than they already possess."
}

func expectedOutputDiscourseOnTheMethod() string {
	return "Goodsenseis,ofallthingsamongmen,themostequallydistributed;foreveryonethinkshimselfsoabundantlyprovidedwithit,thatthoseevenwhoarethemostdifficulttosatisfyineverythingelse,donotusuallydesirealargermeasureofthisquality thantheyalreadypossess."
}
