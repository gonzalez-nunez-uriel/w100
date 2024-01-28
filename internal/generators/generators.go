// Functions that generate test examples

package generators

import (
	"math/rand"
	"strings"
)

func CreateStringExamples(runeSet []rune) (string, string) {
	mismatchAtThisWordCount := rand.Intn(10) + 1
	wordCountAfterMismatch := rand.Intn(10)

	var leftWords []string
	var rightWords []string
	for i := 0; i < mismatchAtThisWordCount-1; i++ {
		newWord := createWord(runeSet)
		leftWords = append(leftWords, newWord)
		rightWords = append(rightWords, newWord)
	}

	// chances are that these words are going to differ
	// in the first few runes. Is that an issue?
	// It must be guaranteed that both strings have a difference
	var mismatchLeftWord string
	var mismatchRightWord string
	for {

		mismatchLeftWord = createWord(runeSet)
		mismatchRightWord = createWord(runeSet)

		if mismatchLeftWord != mismatchRightWord {
			break
		}
	}

	leftWords = append(leftWords, mismatchLeftWord)
	rightWords = append(rightWords, mismatchRightWord)
	// only a single mismatch is need
	// it might be a good idea to guarantee that there is only one difference per example
	// for debugging purposes
	for i := 0; i < wordCountAfterMismatch; i++ {
		newWord := createWord(runeSet)
		leftWords = append(leftWords, newWord)
		rightWords = append(rightWords, newWord)
	}

	// would the absence of whitespace in the beggining and the end be an issue?
	left := strings.Join(leftWords, " ")
	right := strings.Join(rightWords, " ")
	return left, right
}

func createWord(runeSet []rune) string {

	wordSize := rand.Intn(10) + 1
	var runes []rune
	for i := 0; i < wordSize; i++ {
		newRune := runeSet[rand.Intn(len(runeSet))]
		runes = append(runes, newRune)
	}

	// thanks https://yourbasic.org/golang/convert-string-to-rune-slice/
	// there is a chance for optimization
	return string(runes)
}
