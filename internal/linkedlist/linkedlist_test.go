package linkedlist

import "testing"

func TestLinkedListLessons(t *testing.T) {
	t.Run("Lesson03", func(t *testing.T) {
		t.Run("TestLinkedList", TestLinkedList) // Contains all the original tests
		t.Run("TestLinkedList_EdgeCases", TestLinkedList_EdgeCases)
	})
}

func TestLinkedList(t *testing.T) {
	list := NewLinkedList()
	if !list.IsEmpty() {
		t.Errorf("Expected empty list, got non-empty")
	}
	list.Add("A")
	list.Add("B")
	list.Add("C")
	if val, _ := list.Peek(); val != "A" {
		t.Errorf("Expected Peek to return 'A', got '%v'", val)
	}
	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}
	list.Remove()
	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}
	list.Remove()
	list.Remove()
	if !list.IsEmpty() {
		t.Errorf("Expected empty list after removals, got non-empty")
	}
	err := list.Remove()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	_, err = list.Peek()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestLinkedList_EdgeCases(t *testing.T) {
	list := NewLinkedList()

	// Test removing from an already empty list
	err := list.Remove()
	if err == nil {
		t.Errorf("Expected error when removing from an empty list, got nil")
	}

	// Test peeking into an already empty list
	_, err = list.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking into an empty list, got nil")
	}

	// Add and then remove a single element.
	list.Add("X")
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}
	list.Remove()
	if !list.IsEmpty() {
		t.Error("Expected to be empty")
	}
}
