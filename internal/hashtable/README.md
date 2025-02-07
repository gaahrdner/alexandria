# Hash Table

This package provides an implementation of a Hash Table (Hash Map) in Go, using chaining for collision resolution.

## When to Use Hash Tables

* **Fast Lookups:** When you need very fast *average-case* performance for inserting, deleting, and searching by key (O(1) on average).
* **Unordered Data:** When you *don't* need the data to be stored in a specific order.
* **Key-Value Pairs:** When you need to associate values with unique keys.
* **Counting Occurrences:** Counting frequencies

## When *Not* to Use Hash Tables

* **Sorted Data:** When you need to maintain data in sorted order (use a BST or sorted array).
* **Range Queries:** When you need to find elements within a key range (use a BST).
* **Worst-Case Performance:** When you absolutely cannot tolerate the worst-case O(n) performance.
* **Small Datasets:** For very small datasets, the overhead of hashing might not be worth it.

## Time and Space Complexity

```markdown
| Operation | Time Complexity (Avg) | Time Complexity (Worst) | Space Complexity |
| --------- | --------------------- | ----------------------- | ---------------- |
| Insert    | O(1)                  | O(n)                    | O(1)             |
| Delete    | O(1)                  | O(n)                    | O(1)             |
| Search    | O(1)                  | O(n)                    | O(1)             |
| **Total** |                       |                         | O(n)             |
```

The key point is the difference between average-case (O(1)) and worst-case (O(n)) performance. The worst case occurs when there are many collisions, and a good hash function and resizing strategy are crucial for maintaining good performance.  The *load factor* (number of elements / number of buckets) is an important metric.

## Usage

```go
package main

import (
    "fmt"
    "github.com/gaahrdner/library/internal/hashtable"
)

func main() {
    ht := hashtable.NewHashTable(10)
    ht.Insert("apple", 1)
    ht.Insert("banana", 2)
    ht.Insert("cherry", 3)

    val, ok := ht.Search("banana")
    if ok {
        fmt.Println("Value for banana:", val) // Output: Value for banana: 2
    }
    _ = ht.Delete("apple")

}
```
