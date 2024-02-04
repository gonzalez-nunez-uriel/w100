// MIT License
//
// Copyright (c) 2024 Uriel González Núñez
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the “Software”), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package generators

import (
	"testing"
)

// The function CreteStringExamples generates strings with properties that are used by
// property based tests.
func TestCreateStringExamples(t *testing.T) {

	const numOfProperties = 9
	var flags [numOfProperties]bool

	const numOfIterations = 10000

	for i := 0; i < numOfIterations; i++ {

		left, right := CreateStringExamples()

		// 1) It should output two strings that are equal
		if left == right {
			flags[0] = true
		}

		// 2) It should output strings with beggining whitespaces

		// 3) It should output strings with trailing whitespaces

		// 4) It should output strings that differ on the first char

		// 5) It should output strings that differ on the last char

		// 6) It should output strings that differ on one char

		// 7) It should output strings that differ on multiple consecutive chars

		// 8) It should output strings that differ on more than one nonconsecutive char

		// 9) It should output strings that differ on more than one window of consecutive chars

	}
}
