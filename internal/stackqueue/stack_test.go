package stackqueue

import "testing"

func TestStackLessons(t *testing.T) {
	t.Run("Lesson04", func(t *testing.T) {
		t.Run("TestStack", TestStack)
		t.Run("TestStack_EdgeCases", TestStack_EdgeCases)
	})
}

func TestStack(t *testing.T) {
	stack := NewStack()
	if !stack.IsEmpty() {
		t.Errorf("Expected empty stack, got non-empty")
	}
	stack.Push("A")
	stack.Push("B")
	stack.Push("C")
	if val, _ := stack.Peek(); val != "C" {
		t.Errorf("Expected Peek to return 'C', got '%v'", val)
	}
	if stack.Size() != 3 {
		t.Errorf("Expected size 3, got %d", stack.Size())
	}
	val, _ := stack.Pop()
	if val != "C" {
		t.Errorf("Expected popped value 'C', got '%v'", val)
	}
	if stack.Size() != 2 {
		t.Errorf("Expected size 2, got %d", stack.Size())
	}
	stack.Pop()
	stack.Pop()
	_, err := stack.Pop()
	if err == nil {
		t.Errorf("Expected error from popping an empty stack")
	}
	_, err = stack.Peek()
	if err == nil {
		t.Errorf("Expected error from popping an empty stack")
	}
}

func TestStack_EdgeCases(t *testing.T) {
	stack := NewStack()

	// Test popping from an empty stack
	_, err := stack.Pop()
	if err == nil {
		t.Errorf("Expected error when popping from an empty stack, got nil")
	}

	// Test peeking into an empty stack
	_, err = stack.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking into an empty stack, got nil")
	}
}
