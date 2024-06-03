package chapter2

import (
	"errors"
)

// LinkedList is a generic linked list that supports operations such as insertion, removal, and searching.
// - The time complexity of insertion is O(1).
// - The time complexity of searching is O(n).
// - The space complexity of this data structure is O(n).
type LinkedList[T comparable] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

func (s *LinkedList[T]) GetHead() *Node[T] {
	return s.head
}
func (s *LinkedList[T]) GetTail() *Node[T] {
	return s.tail
}

func (s *LinkedList[T]) Length() int {
	return s.length
}

func (s *LinkedList[T]) InsertAtBeginning(elem T) {
	currentHead := s.GetHead()
	newHead := &Node[T]{
		Data: elem,
		Next: currentHead,
	}

	s.head = newHead
	if currentHead == nil {
		s.tail = newHead
	}
	s.length++
}

func (s *LinkedList[T]) InsertAtEnd(elem T) {
	newTail := &Node[T]{
		Data: elem,
	}
	if s.GetTail() == nil {
		s.head = newTail
		s.tail = newTail
	} else {
		s.tail.Next = newTail
		s.tail = newTail
	}
	s.length++
}

func (s *LinkedList[T]) RemoveAtBeginning() error {
	currentHead := s.GetHead()
	if currentHead == nil {
		return errors.New("empty list")
	}
	s.head = currentHead.Next
	s.length--
	if s.length == 0 {
		s.tail = nil
	}
	return nil
}
func (s *LinkedList[T]) RemoveAtEnd() error {
	tail := s.GetTail()
	if tail == nil {
		return errors.New("empty list")
	}

	toRemove := s.Length() - 2
	if toRemove < 0 {
		s.head = nil
		s.tail = nil
		s.length = 0
		return nil
	}
	penult, err := s.FindAtIndex(toRemove)
	if err != nil {
		return err
	}
	penult.Next = nil
	s.tail = penult
	s.length--
	if s.length == 1 {
		s.head = s.tail
	}

	return nil
}
func (s *LinkedList[T]) FindAtIndex(n int) (*Node[T], error) {
	if 0 > n || s.Length() <= n {
		return nil, errors.New("index out of range")
	}
	toCheck := s.GetHead()
	for range n {
		toCheck = toCheck.Next
	}
	return toCheck, nil
}

func (s *LinkedList[T]) Find(data T) *Node[T] {
	current := s.GetHead()
	for current != nil {
		if current.Data == data {
			return current
		}
		current = current.Next
	}
	return nil
}
