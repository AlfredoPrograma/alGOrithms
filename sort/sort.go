// Package sort implements a set of utility functions for sorting collections of elements.
package sort

import "cmp"

type SortFn[T cmp.Ordered, E ~[]T] func(elements E) E

// BubbleSort performs a bubble sort on a slice haystack in place.
// It returns the sorted array in ascending order.
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
