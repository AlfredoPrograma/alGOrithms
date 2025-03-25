// Package graphs implements a set of algorithms related with graphs data structure.
package graphs

import (
	"github.com/alfredoprograma/algo/ds"
)

type Graph[T comparable] map[T][]T

// BFS (Breadth-First algorithm) reports whether a target is found within a graphs network.
//
// Basically it searches over the layers and process nodes one by one; checking if current node is the needed target.
// Each time a node is processed, its neighbors are enqueued to be processed later.
// This process is repeated until the target is found or the processing queue is empty.
func (g Graph[T]) BFS(start T, target T) bool {
	processed := make(map[T]bool, len(g))
	queue := ds.NewSinglyLinkedList(g[start]...)

	for !queue.IsEmpty() {
		current := queue.DeleteFromBeginning()

		if processed[current] {
			continue
		}

		if current == target {
			return true
		}

		for _, next := range g[current] {
			queue.InsertAtEnd(next)
		}

		processed[current] = true
	}

	return false
}
