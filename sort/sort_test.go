package sort

import (
	"cmp"
	"reflect"
	"slices"
	"testing"
)

var UNSORTED_SLICE = []int{9, 4, 2, 7, 3}
var UNSORTED_SLICE_WITH_REPETITIONS = []int{9, 9, 9, 4, 3, 2, 2, 7, 7, 3, 4}
var EMPTY_SLICE = []int{}
var SAME_VALUE_SLICE = []int{0, 0, 0}

func copyAndSort[T cmp.Ordered, E ~[]T](elements E) E {
	sorted := make(E, len(elements))
	copy(sorted, elements)
	slices.Sort(sorted)

	return sorted
}

func TestSort(t *testing.T) {
	type testCase struct {
		elements []int
		name     string
	}
	sortingAlgorithms := []SortFn[int, []int]{BubbleSort[int, []int]}
	cases := []testCase{
		{UNSORTED_SLICE, "should sort unsorted slice"},
		{UNSORTED_SLICE_WITH_REPETITIONS, "should sort an slice with repetitions"},
		{[]int{}, "should return empty slice if there are not elements to sort"},
		{[]int{0, 0, 0}, "should return same array if all elements are same"},
	}

	for _, sortFn := range sortingAlgorithms {
		for _, tc := range cases {
			sorted := copyAndSort(tc.elements)
			got := sortFn(tc.elements)

			if !reflect.DeepEqual(sorted, got) {
				t.Errorf("expected %v but got %v", sorted, got)
			}
		}
	}
}
