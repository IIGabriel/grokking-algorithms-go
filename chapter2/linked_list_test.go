package chapter2

import (
	"testing"
)

func createLinkedListWithElements[T comparable](elements ...T) LinkedList[T] {
	ll := LinkedList[T]{}
	for _, elem := range elements {
		ll.InsertAtEnd(elem)
	}
	return ll
}

func TestInsertAtBeginning(t *testing.T) {
	ll := LinkedList[int]{}
	ll.InsertAtBeginning(1)

	if ll.GetTail() != ll.GetHead() {
		t.Errorf("expected head and tail to be the same, got different")
	}

	if ll.Length() != 1 {
		t.Errorf("expected length 1, got %d", ll.Length())
	}

	ll.InsertAtBeginning(2)

	if ll.GetTail() == ll.GetHead() {
		t.Errorf("expected head and tail to be different, got the same")
	}

	if ll.Length() != 2 {
		t.Errorf("expected length 2, got %d", ll.Length())
	}

	ll.InsertAtBeginning(3)

	if ll.Length() != 3 {
		t.Errorf("expected length 3, got %d", ll.Length())
	}
	if ll.GetHead().Data != 3 {
		t.Errorf("expected head 3, got %d", ll.GetHead().Data)
	}
	if ll.GetTail().Data != 1 {
		t.Errorf("expected tail 1, got %d", ll.GetTail().Data)
	}
}

func TestInsertAtEnd(t *testing.T) {
	ll := LinkedList[int]{}
	ll.InsertAtEnd(1)

	if ll.GetTail() != ll.GetHead() {
		t.Errorf("expected head and tail to be the same, got different")
	}

	if ll.Length() != 1 {
		t.Errorf("expected length 1, got %d", ll.Length())
	}

	ll.InsertAtEnd(2)

	if ll.GetTail() == ll.GetHead() {
		t.Errorf("expected head and tail to be different, got the same")
	}

	if ll.Length() != 2 {
		t.Errorf("expected length 2, got %d", ll.Length())
	}

	ll.InsertAtEnd(3)

	if ll.Length() != 3 {
		t.Errorf("expected length 3, got %d", ll.Length())
	}
	if ll.GetHead().Data != 1 {
		t.Errorf("expected head 1, got %d", ll.GetHead().Data)
	}
	if ll.GetTail().Data != 3 {
		t.Errorf("expected tail 3, got %d", ll.GetTail().Data)
	}
}

func TestRemoveAtBeginning(t *testing.T) {
	t.Run("EmptyList", func(t *testing.T) {
		ll := LinkedList[int]{}
		err := ll.RemoveAtBeginning()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("SingleElementList", func(t *testing.T) {
		ll := LinkedList[int]{}
		ll.InsertAtEnd(5)
		err := ll.RemoveAtBeginning()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ll.GetTail() != nil {
			t.Errorf("expected tail to be nil, got %v", ll.GetTail())
		}
		if ll.GetHead() != nil {
			t.Errorf("expected head to be nil, got %v", ll.GetHead())
		}
	})

	t.Run("MultipleElementsList", func(t *testing.T) {
		ll := createLinkedListWithElements(1, 2, 3)
		err := ll.RemoveAtBeginning()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ll.GetHead().Data != 2 {
			t.Errorf("expected head 2, got %d", ll.GetHead().Data)
		}
		if ll.Length() != 2 {
			t.Errorf("expected length 2, got %d", ll.Length())
		}
	})
}

func TestRemoveAtEnd(t *testing.T) {
	t.Run("EmptyList", func(t *testing.T) {
		ll := LinkedList[int]{}
		err := ll.RemoveAtEnd()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("SingleElementList", func(t *testing.T) {
		ll := LinkedList[int]{}
		ll.InsertAtEnd(1)
		err := ll.RemoveAtEnd()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ll.GetHead() != nil {
			t.Errorf("expected head to be nil, got %v", ll.GetHead())
		}
		if ll.GetTail() != nil {
			t.Errorf("expected tail to be nil, got %v", ll.GetTail())
		}
	})
	t.Run("TwoElementList", func(t *testing.T) {
		ll := LinkedList[int]{}
		ll.InsertAtEnd(1)
		ll.InsertAtEnd(2)
		err := ll.RemoveAtEnd()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if ll.GetHead().Data != 1 {
			t.Errorf("expected head to be 1, got %v", ll.GetHead())
		}
		if ll.GetTail().Data != 1 {
			t.Errorf("expected tail to be 1, got %v", ll.GetTail())
		}
	})

	t.Run("MultipleElementsList", func(t *testing.T) {
		ll := createLinkedListWithElements(1, 2, 3)
		err := ll.RemoveAtEnd()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if ll.GetTail().Data != 2 {
			t.Errorf("expected tail 2, got %d", ll.GetTail().Data)
		}
		if ll.Length() != 2 {
			t.Errorf("expected length 2, got %d", ll.Length())
		}
	})
}

func TestFindAtIndex(t *testing.T) {
	t.Run("IndexOutOfBounds", func(t *testing.T) {
		ll := createLinkedListWithElements(1, 2, 3)
		_, err := ll.FindAtIndex(3)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("ValidIndex", func(t *testing.T) {
		ll := createLinkedListWithElements(1, 2, 3)
		node, err := ll.FindAtIndex(1)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if node.Data != 2 {
			t.Errorf("expected data 2, got %d", node.Data)
		}
	})
}

func TestFind(t *testing.T) {
	ll := createLinkedListWithElements(1, 2, 3)
	node := ll.Find(2)
	if node == nil {
		t.Errorf("expected node with data 2, got nil")
	}
	if node != nil && node.Data != 2 {
		t.Errorf("expected data 2, got %d", node.Data)
	}

	t.Run("NotFound", func(t *testing.T) {
		node := ll.Find(4)
		if node != nil {
			t.Errorf("expected nil, got node with data %d", node.Data)
		}
	})
}
