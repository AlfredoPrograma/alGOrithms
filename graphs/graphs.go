// Package graphs implements a set of algorithms related with graphs data structure.
package graphs

import (
	"github.com/alfredoprograma/algo/ds"
)

type Graph[T comparable] map[T][]T

func BFS[T comparable](graph Graph[T], start T, target T) bool {
	processed := make(map[T]bool, len(graph))
	queue := ds.NewSinglyLinkedList(graph[start]...)

	for !queue.IsEmpty() {
		current := queue.DeleteFromBeginning()

		if processed[current] {
			continue
		}

		if current == target {
			return true
		}

		for _, next := range graph[current] {
			queue.InsertAtEnd(next)
		}

		processed[current] = true
	}

	return false
}
