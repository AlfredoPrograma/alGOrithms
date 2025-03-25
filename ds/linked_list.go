package ds

import (
	"fmt"
	"iter"
	"slices"
)

func invalidPositionError(pos int) error {
	return fmt.Errorf("invalid element insertion at position %d", pos)
}

type node[T any] struct {
	value T
	next  *node[T]
}

func newNode[T any](value T) node[T] {
	return node[T]{value, nil}
}

func (n node[T]) Value() T {
	return n.value
}

func (n node[T]) Next() *node[T] {
	return n.next
}

// A SinglyLinkedList is a heap memory sequential data structure which holds a collection of
// consecutive elements linked together.
//
// Notice it is a singly linked list, so their elements are linked just in a forward fashion.
type SinglyLinkedList[T any] struct {
	size int
	head *node[T]
	tail *node[T]
}

// NewSinglyLinkedList creates a new empty SinglyLinkedList tied to type T.
func NewSinglyLinkedList[T any](elems ...T) SinglyLinkedList[T] {
	l := SinglyLinkedList[T]{}

	if len(elems) == 0 {
		return l
	}

	l.fillForward(slices.All(elems))
	return l
}

// Iter creates a single-value iterator over the elements of the list l.
func (l *SinglyLinkedList[T]) Iter() iter.Seq[*node[T]] {
	return func(yield func(*node[T]) bool) {
		node := l.head

		for node != nil {
			if !yield(node) {
				return
			}

			node = node.next
		}
	}
}

// Iter creates an indexed iterator over the elements of the list l.
func (l *SinglyLinkedList[T]) Iter2() iter.Seq2[int, *node[T]] {
	return func(yield func(int, *node[T]) bool) {
		i := 0
		node := l.head

		for node != nil {
			if !yield(i, node) {
				return
			}

			i++
			node = node.next
		}
	}
}

// InsertAtBeginning inserts a new element at the beginning of the list l.
func (l *SinglyLinkedList[T]) InsertAtBeginning(value T) {
	if l.IsEmpty() {
		l.insertOnEmptyList(value)
		return
	}

	node := newNode(value)
	head := l.head
	node.next = head
	l.head = &node
	l.size++
}

// InsertAtEnd inserts a new element at the end of the list l.
func (l *SinglyLinkedList[T]) InsertAtEnd(value T) {
	if l.IsEmpty() {
		l.insertOnEmptyList(value)
		return
	}

	node := newNode(value)
	oldTail := l.tail
	oldTail.next = &node
	l.tail = &node
	l.size++
}

// InsertAtPosition inserts a new element at the provided position. Also, reports whether provided position is invalid.
func (l *SinglyLinkedList[T]) InsertAtPosition(value T, pos int) error {
	if pos >= l.Size() || pos < 0 {
		return invalidPositionError(pos)
	}

	if pos == 0 {
		l.InsertAtBeginning(value)
		return nil
	}

	var prev *node[T]
	newNode := newNode(value)

	for i, n := range l.Iter2() {
		if i == pos-1 {
			prev = n
			continue
		}
	}

	next := prev.next
	prev.next = &newNode
	newNode.next = next
	return nil
}

// DeleteFromBeginning removes the first element of the list (element at l.head) and returns its value.
func (l *SinglyLinkedList[T]) DeleteFromBeginning() T {
	if l.Size() == 0 {
		return *new(T)
	}

	if l.Size() == 1 {
		return l.deleteOnSingleElementList()
	}

	oldHead := l.head
	l.head = oldHead.next
	l.size--

	return oldHead.value
}

// DeleteFromEnd removes the last element of the list (element at l.tail) and returns its value.
func (l *SinglyLinkedList[T]) DeleteFromEnd() T {
	if l.Size() == 0 {
		return *new(T)
	}

	if l.Size() == 1 {
		return l.deleteOnSingleElementList()
	}

	var preLast *node[T]

	for i, n := range l.Iter2() {
		if i == l.Size()-2 {
			preLast = n
		}
	}

	oldTail := l.tail
	preLast.next = nil
	l.tail = preLast
	l.size--

	return oldTail.value
}

// DeleteAtPosition removes the element at the provided position and returns its value. Also, reports whether provided position is invalid.
func (l *SinglyLinkedList[T]) DeleteAtPosition(pos int) (T, error) {
	if pos >= l.Size() || pos < 0 {
		return *new(T), invalidPositionError(pos)
	}

	if pos == 0 {
		return l.DeleteFromBeginning(), nil
	}

	if pos == l.Size()-1 {
		return l.DeleteFromEnd(), nil
	}

	var (
		prev    *node[T]
		current *node[T]
	)

	for i, n := range l.Iter2() {
		if i == pos-1 {
			prev = n
		}

		if i == pos {
			current = n
		}
	}

	node := current
	prev.next = current.next

	return node.value, nil
}

// ForEach iterates over each element of the list and performs an operation over each one.
//
// Notice, operation does not mutate the internal value of the element.
func (l *SinglyLinkedList[T]) ForEach(callback func(index int, n *node[T])) {
	for i, n := range l.Iter2() {
		callback(i, n)
	}
}

// Get returns the element at the provided position. If position is invalid or element is not found, returns nil.
func (l *SinglyLinkedList[T]) Get(pos int) *node[T] {
	for i, n := range l.Iter2() {
		if pos == i {
			return n
		}
	}

	return nil
}

// Size returns the current size of the list.
func (l *SinglyLinkedList[T]) Size() int {
	return l.size
}

// IsEmpty reports whether the list has no elements.
func (l *SinglyLinkedList[T]) IsEmpty() bool {
	return l.Size() == 0
}

// insertOnEmptyList is a private helper function to handle first element insertion on an empty list.
func (l *SinglyLinkedList[T]) insertOnEmptyList(value T) {
	node := newNode(value)

	l.head = &node
	l.tail = &node
	l.size++
}

// deleteOnSingleElementList is a private helper function to remove an element within a single element list.
func (l *SinglyLinkedList[T]) deleteOnSingleElementList() T {
	value := l.head.value
	l.head = nil
	l.tail = nil
	l.size--

	return value
}

// fillForward the SinglyLinkedList l based on the given iterator in forward way.
func (l *SinglyLinkedList[T]) fillForward(iter iter.Seq2[int, T]) {
	for _, v := range iter {
		l.InsertAtEnd(v)
	}
}
