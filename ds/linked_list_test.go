package ds

import (
	"iter"
	"reflect"
	"slices"
	"testing"
)

func createNumbersSlice(end int) (numbers []int) {
	for n := range end {
		numbers = append(numbers, n)
	}

	return numbers
}

func mapNodesToNumbers(iter iter.Seq[*node[int]]) []int {
	var numbers []int

	for _, n := range slices.Collect(iter) {
		numbers = append(numbers, n.value)
	}

	return numbers
}

func TestLinkedList(t *testing.T) {
	t.Run("should create an empty singly linked list with zero initialized attributes", func(t *testing.T) {
		expected := SinglyLinkedList[any]{}
		got := NewSinglyLinkedList[any]()

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %+v but got %+v", expected, got)
		}
	})

	t.Run("should create a filled linked list", func(t *testing.T) {
		expected := createNumbersSlice(10)
		list := NewSinglyLinkedList(expected...)
		got := mapNodesToNumbers(list.Iter())

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("should create an iterator and iterate over linked list", func(t *testing.T) {
		numbers := createNumbersSlice(10)
		l := NewSinglyLinkedList(numbers...)

		for expected, got := range l.Iter2() {
			if expected != got.Value() {
				t.Errorf("expected %v but got %v", numbers, got)
			}
		}
	})

	t.Run("should insert a new element at beginning", func(t *testing.T) {
		numbers := createNumbersSlice(10)
		list := NewSinglyLinkedList(numbers...)
		expected := slices.Concat([]int{-1}, numbers)
		list.InsertAtBeginning(-1)
		got := mapNodesToNumbers(list.Iter())

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("should insert a new element at end", func(t *testing.T) {
		numbers := createNumbersSlice(10)
		list := NewSinglyLinkedList(numbers...)
		expected := slices.Concat(numbers, []int{10})
		list.InsertAtEnd(10)
		got := mapNodesToNumbers(list.Iter())

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("should insert a new element at position 3", func(t *testing.T) {
		numbers := []int{0, 1, 2, 4, 5}
		list := NewSinglyLinkedList(numbers...)
		list.InsertAtPosition(3, 3)
		expected := []int{0, 1, 2, 3, 4, 5}
		got := mapNodesToNumbers(list.Iter())

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("should return an error when try to insert an element at invalid position", func(t *testing.T) {
		numbers := []int{0, 1, 2, 4, 5}
		list := NewSinglyLinkedList(numbers...)
		err := list.InsertAtPosition(3, 10)

		if err == nil {
			t.Error("expected overflowing position error")
		}
	})

	t.Run("should delete the first element of the list", func(t *testing.T) {
		numbers := createNumbersSlice(10)
		list := NewSinglyLinkedList(numbers...)

		popped := list.DeleteFromBeginning()

		if popped != numbers[0] {
			t.Errorf("expected %v but got %v", numbers[0], popped)
		}

		if list.head.value != numbers[1] {
			t.Errorf("expected %v, but got %v", numbers[1], list.head.value)
		}
	})

	t.Run("should delete the last element of the list", func(t *testing.T) {
		numbers := createNumbersSlice(10)
		list := NewSinglyLinkedList(numbers...)

		popped := list.DeleteFromEnd()
		last := numbers[len(numbers)-1]
		preLast := numbers[len(numbers)-2]

		if popped != last {
			t.Errorf("expected %v but got %v", last, popped)
		}

		if list.tail.value != preLast {
			t.Errorf("expected %v but got %v", preLast, list.tail.value)
		}
	})

	t.Run("should delete the element at position 3", func(t *testing.T) {
		numbers := []int{0, 1, 2, 3, 4, 5}
		list := NewSinglyLinkedList(numbers...)
		popped, _ := list.DeleteAtPosition(3)
		remainingSlice := []int{0, 1, 2, 4, 5}
		remainingList := mapNodesToNumbers(list.Iter())

		if popped != numbers[3] {
			t.Errorf("expected %v but got %v", numbers[3], popped)
		}

		if !reflect.DeepEqual(remainingSlice, remainingList) {
			t.Errorf("expected %v but got %v", remainingSlice, remainingList)
		}
	})

	t.Run("should return an error when try to remove an element at invalid position", func(t *testing.T) {
		numbers := []int{0, 1, 2, 4, 5}
		list := NewSinglyLinkedList(numbers...)
		_, err := list.DeleteAtPosition(10)

		if err == nil {
			t.Error("expected overflowing position error")
		}
	})

	t.Run("should traverse every element and multiply by two", func(t *testing.T) {
		numbers := createNumbersSlice(10)
		list := NewSinglyLinkedList(numbers...)

		list.ForEach(func(index int, n *node[int]) {
			expected := numbers[index] * 2
			got := n.value * 2

			if expected != got {
				t.Errorf("expected %v but got %v", expected, got)
			}
		})
	})

	t.Run("should return the element at position 6", func(t *testing.T) {
		numbers := createNumbersSlice(10)
		list := NewSinglyLinkedList(numbers...)
		expected := numbers[6]
		got := list.Get(6)

		if expected != got.value {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}
