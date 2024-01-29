/*
MIT License

Copyright (c) 2024 Uriel González Núñez

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the “Software”), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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
