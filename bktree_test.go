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
