/*
The MIT License (MIT)
Copyright (c) 2013

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package bktree

import (
	"fmt"
	"testing"
)

var searchTests = []struct {
	dict           []string
	query          string
	expectedResult []string
}{
	{[]string{"cook", "book", "books", "what", "water"}, "a", []string{"what", "water"}},
}

func TestSearch(t *testing.T) {
	for index, searchTest := range searchTests {
		tree := Tree{}
		for _, v := range searchTest.dict {
			tree.Add(v)
		}

		result := tree.Search("wat", 2)

		if !areEqual(result, searchTest.expectedResult) {
			output := fmt.Sprintf("%v \t for query %v expected %v but obtained %v",
				index, searchTest.query, searchTest.expectedResult, result)
			t.Errorf(output)
		}
	}
}

func areEqual(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
