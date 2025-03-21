package search

import "testing"

func TestBinarySearch(t *testing.T) {
	type testCase struct {
		target   int
		haystack []int
		index    int
		found    bool
		name     string
	}

	cases := []testCase{
		{7, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6, true, "should find a number at middle of the haystack"},
		{10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, true, "should find a number at last place of the haystack"},
		{1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, true, "should find a number at first of the haystack"},
		{11, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, false, "should return zero index and false for non existing element within haystack"},
		{5, []int{}, 0, false, "should return zero index and false for searching over empty array"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			index, found := BinarySearch(tc.target, tc.haystack)

			if tc.index != index || tc.found != found {
				t.Errorf("expected (index: %d, found: %t) but got (index: %d, found: %t)", tc.index, tc.found, index, found)
			}
		})
	}
}
