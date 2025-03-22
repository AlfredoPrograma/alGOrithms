package sort

import (
	"cmp"
	"reflect"
	"runtime"
	"slices"
	"testing"
)

var UNSORTED_SLICE = []int{9, 4, 2, 7, 3}
var UNSORTED_SLICE_WITH_REPETITIONS = []int{9, 9, 9, 4, 3, 2, 2, 7, 7, 3, 4}
var EMPTY_SLICE = []int{}
var SAME_VALUE_SLICE = []int{0, 0, 0}

func copySlice[T cmp.Ordered, E ~[]T](elements E) E {
	slice := make(E, len(elements))
	copy(slice, elements)

	return slice
}

func getFunctionName(f any) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func TestSort(t *testing.T) {
	type testCase struct {
		elements []int
		name     string
	}
	sortingAlgorithms := []SortFn[int, []int]{SelectionSort[int, []int], BubbleSort[int, []int], Quicksort[int, []int]}
	cases := []testCase{
		{UNSORTED_SLICE, "should sort unsorted slice"},
		{UNSORTED_SLICE_WITH_REPETITIONS, "should sort an slice with repetitions"},
		{[]int{}, "should return empty slice if there are not elements to sort"},
		{[]int{0, 0, 0}, "should return same array if all elements are same"},
	}

	for _, sortFn := range sortingAlgorithms {
		for _, tc := range cases {
			sorted := copySlice(tc.elements)
			got := copySlice(tc.elements)

			slices.Sort(sorted)
			got = sortFn(got)

			if !reflect.DeepEqual(sorted, got) {
				t.Errorf("[%s] expected %v but got %v", getFunctionName(sortFn), sorted, got)
			}
		}
	}
}
