# Stack and Queue

This package provides implementations of Stack and Queue data structures in Go (array-based and linked-list-based).

## Data Structures

* **Stack:** LIFO (Last-In, First-Out).
* **Queue:** FIFO (First-In, First-Out).

## When to Use Stacks

* **Function Call Stack:** Managing function calls (recursion).
* **Undo/Redo:** Storing history of actions.
* **Expression Evaluation:** Evaluating mathematical expressions.
* **Depth-First Search (DFS):** In graph traversal.
* **Backtracking Algorithms:** Exploring solutions and undoing choices.

## When *Not* to Use Stacks

* When you need FIFO behavior (use a queue).
* When you need random access to elements.

## When to Use Queues

* **Breadth-First Search (BFS):** In graph traversal.
* **Print Queues:** Managing print jobs.
* **Task Scheduling:** Handling tasks in arrival order.
* **Buffering:** Storing data temporarily.
* **Asynchronous Communication:** Between processes/threads.

## When *Not* to Use Queues

* When you need LIFO behavior (use a stack).
* When you need random access to elements.

## Time and Space Complexity

### Stack (Array-Based)

```markdown
| Operation | Time Complexity | Space Complexity |
| --------- | --------------- | ---------------- |
| Push      | O(1) (amortized) | O(1)             |
| Pop       | O(1)            | O(1)             |
| Peek      | O(1)            | O(1)             |
| IsEmpty   | O(1)            | O(1)             |
| IsFull    | O(1)            | O(1)             |
| **Total** |                 | O(n)             |
```

### Stack (Linked-List-Based)

```markdown
| Operation | Time Complexity | Space Complexity |
| --------- | --------------- | ---------------- |
| Push      | O(1)            | O(1)             |
| Pop       | O(1)            | O(1)             |
| Peek      | O(1)            | O(1)             |
| IsEmpty   | O(1)            | O(1)             |
| **Total** |                 | O(n)             |
```

### Queue (Array-Based - Circular Buffer)

```markdown
| Operation | Time Complexity | Space Complexity |
| --------- | --------------- | ---------------- |
| Enqueue   | O(1) (amortized) | O(1)             |
| Dequeue   | O(1)            | O(1)             |
| Peek      | O(1)            | O(1)             |
| IsEmpty   | O(1)            | O(1)             |
| IsFull    | O(1)            | O(1)             |
| **Total** |                 | O(n)             |
```markdown

### Queue (Linked-List-Based)

```markdown
| Operation | Time Complexity | Space Complexity |
| --------- | --------------- | ---------------- |
| Enqueue   | O(1)            | O(1)             |
| Dequeue   | O(1)            | O(1)             |
| Peek      | O(1)            | O(1)             |
| IsEmpty   | O(1)            | O(1)             |
| **Total** |                 | O(n)             |
```

For both stacks and queues, the core operations (Push/Pop for stacks, Enqueue/Dequeue for queues) are typically O(1), *especially* for linked-list-based implementations.  Array-based implementations might have amortized O(1) for operations that could require resizing the underlying array.

## Usage

```go
package main

import (
   "fmt"
   "github.com/gaahrdner/library/internal/stackqueue"
)

func main() {
    s := stackqueue.StackArray{}
    _ = s.Push(1)
    _ = s.Push(2)
    val, _ := s.Pop()
    fmt.Println(val) // Output: 2

    q := stackqueue.QueueArray{}
    _ = q.Enqueue(1)
    _ = q.Enqueue(2)
    val, _ = q.Dequeue()
    fmt.Println(val) // Output: 1
}
```
