package stackqueue

import "testing"

func TestQueueLessons(t *testing.T) {
	t.Run("Lesson05", func(t *testing.T) {
		t.Run("TestQueue", TestQueue) // Contains all original tests
		t.Run("TestQueue_EdgeCases", TestQueue_EdgeCases)
	})
}

func TestQueue(t *testing.T) {
	queue := NewQueue()
	if !queue.IsEmpty() {
		t.Errorf("Expected empty queue, got non-empty")
	}
	queue.Enqueue("A")
	queue.Enqueue("B")
	queue.Enqueue("C")
	if val, _ := queue.Peek(); val != "A" {
		t.Errorf("Expected Peek to return 'A', got '%v'", val)
	}
	if queue.Size() != 3 {
		t.Errorf("Expected size 3, got %d", queue.Size())
	}
	val, _ := queue.Dequeue()
	if val != "A" {
		t.Errorf("Expected dequeued value 'A', got '%v'", val)
	}
	if queue.Size() != 2 {
		t.Errorf("Expected size 2, got %d", queue.Size())
	}
	queue.Dequeue()
	queue.Dequeue()
	if !queue.IsEmpty() {
		t.Errorf("Expected empty queue after dequeues, got non-empty")
	}
	_, err := queue.Dequeue()
	if err == nil {
		t.Errorf("Expected error from dequeuing an empty queue")
	}
	_, err = queue.Peek()
	if err == nil {
		t.Errorf("Expected error from peeking an empty queue")
	}
}

func TestQueue_EdgeCases(t *testing.T) {
	queue := NewQueue()

	//Test dequeue on empty queue
	_, err := queue.Dequeue()
	if err == nil {
		t.Errorf("Expected error when dequeuing from an empty queue, got nil")
	}

	//Test peek on empty queue
	_, err = queue.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking into an empty queue, got nil")
	}
}
