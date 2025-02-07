# Binary Search Tree (BST)

This package provides an implementation of a Binary Search Tree in Go.

## When to Use Binary Search Trees

* **Sorted Data:** When you need to store data in sorted order and perform efficient searches, insertions, and deletions.
* **Fast Lookup:** Quickly finding the minimum or maximum element.
* **Ordered Traversal:** Iterating through data in sorted order (in-order traversal).
* **Range Queries:** Finding all elements within a specific range.

## When *Not* to Use Binary Search Trees

* **Unsorted Data:** If you don't need sorted data, a hash table might be better for faster lookups (O(1) average).
* **Worst-Case Performance:** If you can't guarantee the tree will remain balanced, the worst-case performance (O(n)) can be poor. Consider self-balancing trees (AVL, Red-Black) in this case.
* **Simple Key-Value Storage:** If you only need key-value pairs without sorted order or range queries, a hash table is often simpler and faster.

## Time and Space Complexity

### Unbalanced BST (Worst Case)

```markdown
| Operation          | Time Complexity | Space Complexity |
| ------------------ | --------------- | ---------------- |
| Insert             | O(n)            | O(1)             |
| Delete             | O(n)            | O(1)             |
| Search             | O(n)            | O(1)             |
| FindMin/FindMax    | O(n)            | O(1)             |
| InOrderTraversal   | O(n)            | O(n) (recursion)   |
| PreOrderTraversal  | O(n)            | O(n) (recursion)   |
| PostOrderTraversal | O(n)            | O(n) (recursion)   |
| IsEmpty            | O(1)            | O(1)              |
| **Total Space**   |                 | O(n)             |
```

A balanced BST (e.g., AVL tree, Red-Black tree - not implemented here, but good to know about) guarantees O(log n) performance for most operations. An unbalanced BST can degenerate to O(n) in the worst case (e.g., if you insert elements in sorted order).

## Usage

```go
package main

import (
    "fmt"
    "github.com/gaahrdner/library/internal/bst"
)

func main() {
    tree := bst.BST{}
    tree.Insert(5)
    tree.Insert(2)
    tree.Insert(8)
    tree.Insert(1)
    tree.Insert(9)
    tree.Insert(4)

    fmt.Println("In-order Traversal:", tree.InOrderTraversal()) // Output: [1 2 4 5 8 9]
    node := tree.Search(8)
    if node != nil{
        fmt.Println("Found node with data:", node.Data)
    }

}
```
