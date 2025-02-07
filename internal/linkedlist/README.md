# Linked List

This package provides implementations of singly and doubly linked lists in Go.

## Data Structures

* **Singly Linked List:** Each node points to the next node.
* **Doubly Linked List:** Each node points to both the next and previous nodes.

## When to Use Linked Lists

* **Frequent Insertions/Deletions:** Especially in the middle of the list.  These are O(1) (once you have a pointer to the node) in a linked list, compared to O(n) for arrays.
* **Dynamic Size:** When you don't know the size in advance and it needs to grow/shrink frequently.
* **Non-Contiguous Memory:** When you can't allocate a large contiguous block of memory.

## When *Not* to Use Linked Lists

* **Random Access:** When you need to frequently access elements by index (use arrays - O(1)).
* **Memory Overhead:** When memory is tight, and the extra space for pointers is a concern.
* **Simple Iteration:** If you *only* need sequential iteration, arrays are often just as efficient.
* **Cache Locality**: Arrays tend to be faster

## Time and Space Complexity

### Singly Linked List

```markdown
| Operation             | Time Complexity | Space Complexity |
| --------------------- | --------------- | ---------------- |
| InsertAtBeginning     | O(1)            | O(1)             |
| InsertAtEnd           | O(n)            | O(1)             |
| InsertAtPosition      | O(n)            | O(1)             |
| DeleteAtBeginning     | O(1)            | O(1)             |
| DeleteAtEnd           | O(n)            | O(1)             |
| DeleteByValue         | O(n)            | O(1)             |
| DeleteAtPosition      | O(n)            | O(1)             |
| Find                  | O(n)            | O(1)             |
| Traverse              | O(n)            | O(1)             |
| Reverse (Iterative)   | O(n)            | O(1)             |
| IsEmpty               | O(1)            | O(1)             |
| **Total Space**       |                 | O(n)             |
```

### Doubly Linked List

```markdown
| Operation             | Time Complexity | Space Complexity |
| --------------------- | --------------- | ---------------- |
| InsertAtBeginning     | O(1)            | O(1)             |
| InsertAtEnd           | O(1)            | O(1)             |
| InsertAtPosition      | O(n)            | O(1)             |
| DeleteAtBeginning     | O(1)            | O(1)             |
| DeleteAtEnd           | O(1)            | O(1)             |
| DeleteByValue         | O(n)            | O(1)             |
| DeleteAtPosition      | O(n)            | O(1)             |
| Find                  | O(n)            | O(1)             |
| Traverse              | O(n)            | O(1)             |
| Reverse               | O(n)            | O(1)             |
| IsEmpty               | O(1)            | O(1)             |
| **Total Space**      |                 | O(n)             |
```

All operations that involve searching for a specific node or position are O(n) in the worst case (you might have to traverse the entire list).  Insertions and deletions at the beginning (and end, for doubly linked lists) are O(1).

## Usage

```go
package main

import (
    "fmt"
    "github.com/gaahrdner/library/internal/linkedlist"
)

func main() {
    l := linkedlist.SinglyLinkedList{}
    l.InsertAtEnd(1)
    l.InsertAtEnd(2)
    l.InsertAtBeginning(0)
    fmt.Println(l.Traverse()) // Output: [0 1 2]
}
```

## Testing

Run tests with: go test ./...

Run benchmarks with: go test -bench=. -benchmem
