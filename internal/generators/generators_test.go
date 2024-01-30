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
// property based tests. This function generates two strings with the following properties:
// 1) The strings might be equal
// 2) The strings might have beggining and trailing whitespaces
// 3) The strings might be equal up to a char index
// 4) The strings might differ in the first char
// 5) The strings might differ in more than one char
// 6) The strings might differ in multiple chars in a row
// 7) The strings might differ in  more than one non-consecutive char
func TestCreateStringExamples(t *testing.T) {
	left, right := CreateStringExamples()
}
