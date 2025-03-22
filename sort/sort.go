// Package sort implements a set of utility functions for sorting collections of elements.
package sort

import "cmp"

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
