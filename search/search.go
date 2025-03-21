// Package search implements a set of utility functions for searching a target element within a provided collections of elements.
package search

import "cmp"

// BinarySearch performs a binary search on a sorted slice haystack to find target.
// It returns the index of target if found, and a boolean indicating whether the target was found.
//
// The elements in haystack must be sorted in ascending order for the search to work correctly.
//
// The search operates in O(log n) time complexity and O(1) space complexity.
func BinarySearch[E ~[]T, T cmp.Ordered](target T, haystack E) (index int, found bool) {
	left := 0
	right := len(haystack) - 1

	for left <= right {
		pivotIndex := (left + right) / 2
		pivotValue := haystack[pivotIndex]

		if pivotValue == target {
			return pivotIndex, true
		}

		if target < pivotValue {
			right = pivotIndex - 1
			continue
		}

		if target > pivotValue {
			left = pivotIndex + 1
			continue
		}
	}

	return 0, false
}
