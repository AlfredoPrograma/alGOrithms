// Package sort implements a set of utility functions for sorting collections of elements.
package sort

import (
	"cmp"
	"slices"
)

type SortFn[T cmp.Ordered, E ~[]T] func(elements E) E

// BubbleSort performs a bubble sort on a slice haystack in place.
// It returns the sorted slice in ascending order.
//
// The sort operates in O(n^2) time complexity and O(1) space complexity.
func BubbleSort[T cmp.Ordered, E ~[]T](elements E) E {
	for i := 0; i < len(elements); i++ {
		for j := i + 1; j < len(elements); j++ {
			v1 := elements[i]
			v2 := elements[j]

			if v1 <= v2 {
				continue
			}

			elements[i] = v2
			elements[j] = v1
		}
	}

	return elements
}

// SelectionSort performs a selection sort algorithm on a slice haystack in place.
// Basically it checks the slice taking the lowest element and setting it at first place continuously.
//
// It returns the sorted slice in ascending order.
//
// The sort operates in O(n^2) time complexity and O(1) space complexity.
func SelectionSort[T cmp.Ordered, E ~[]T](elements E) E {
	for i := 0; i < len(elements); i++ {
		minIndex := i
		min := elements[i]

		for j := i; j < len(elements); j++ {
			current := elements[j]

			if current < min {
				min = current
				minIndex = j
			}
		}

		tmp := elements[i]
		elements[i] = min
		elements[minIndex] = tmp
	}

	return elements
}

// Quicksort performs a quick sort algorithm on a slice haystack.
//
// It returns the sorted slice in ascending order.
//
// Quicksort algorithm is based on divide an conquer paradigm, where the complete slice sorting is
// divided into smaller sub-problems until reach base cases. Base cases for quicksort are:
//
// - Sub slice length is 1 or 0; in this case, just return the single/none element slice.
//
// - Sub slice length is 2; in this case, we can just compare elements and swap if needed in order to sort the slice.
//
// If sub slice length is 3 or greater, then recursively call quicksort.
//
// The sorts operates in O(nlogn) time complexity and O(nlogn) space complexity.
//
// The worst case really could be O(n^2) for both, time and space complexity; but this case is uncommon and only happens when,
// the slice is already sorted and the selected pivot is always on some edge.
func Quicksort[T cmp.Ordered, E ~[]T](elements E) E {
	swap := func(e E) E {
		v1 := e[0]
		v2 := e[1]

		if v1 > v2 {
			tmp := v1
			e[0] = v2
			e[1] = tmp
		}

		return e
	}

	partition := func(e E, mid T) (smallests E, repeated E, greatests E) {
		for _, elem := range e {
			if elem == mid {
				repeated = append(repeated, elem)
			}

			if elem > mid {
				greatests = append(greatests, elem)
			}

			if elem < mid {
				smallests = append(smallests, elem)
			}
		}

		return smallests, repeated, greatests
	}

	if len(elements) <= 1 {
		return elements
	}

	if len(elements) == 2 {
		swap(elements)
	}

	mid := elements[len(elements)/2]
	left, repeated, right := partition(elements, mid)

	return slices.Concat(Quicksort(left), repeated, Quicksort(right))
}
