# Graph

This package provides implementations of graph data structures and algorithms in Go.

## Representations

* **Adjacency Matrix:** Uses a 2D array.
* **Adjacency List:** Uses a map of lists.

## Algorithms

* **Depth-First Search (DFS)**
* **Breadth-First Search (BFS)**

## When to Use Graphs:**

* **Relationships:** Representing relationships between objects (social networks, family trees).
* **Networks:** Modeling networks (computer networks, transportation networks).
* **Paths and Connectivity:** Finding paths between nodes (shortest path, route planning).
* **Dependencies:** Representing dependencies between tasks.
* **Recommendations:** Suggesting related items.

## When *Not* to Use Graphs:**

* **Simple Hierarchies:** If relationships are strictly hierarchical (parent-child), a tree might be better.
* **Ordered Data:** If you just need to store data in order, an array, linked list, or BST might be more appropriate.
* **Key-Value Pairs:** If you just need key-value pairs, a hash table might be better.

## Time and Space Complexity

### Adjacency List

```markdown
| Operation      | Time Complexity | Space Complexity |
| -------------- | --------------- | ---------------- |
| AddEdge        | O(1)            | O(1)             |
| RemoveEdge     | O(degree(u))   | O(1)             |
| HasEdge        | O(degree(u))   | O(1)             |
| Neighbors      | O(degree(u))   | O(degree(u))     |
| DFS            | O(V + E)        | O(V)             |
| BFS            | O(V + E)        | O(V)             |
| **Total Space** |                 | O(V + E)         |
```

### Adjacency Matrix

```markdown
| Operation      | Time Complexity | Space Complexity |
| -------------- | --------------- | ---------------- |
| AddEdge        | O(1)            | O(1)             |
| RemoveEdge     | O(1)            | O(1)             |
| HasEdge        | O(1)            | O(1)             |
| Neighbors      | O(V)            | O(V)             |
| DFS            | O(V^2)          | O(V)             |
| BFS            | O(V^2)          | O(V)             |
| **Total Space** |                 | O(V^2)           |
```

* **Adjacency List:** Generally preferred for sparse graphs.  DFS and BFS are more efficient with adjacency lists (O(V + E)).  Space complexity is O(V + E).
* **Adjacency Matrix:** Better for dense graphs. `HasEdge` is O(1).  Space complexity is O(V^2).

## Usage

```go
package main

import (
    "fmt"
    "github.com/gaahrdner/library/internal/graph"
)

func main() {
    g := graph.AdjacencyList{}
    g.AddEdge(0, 1)
    g.AddEdge(0, 2)
    g.AddEdge(1, 2)
    g.AddEdge(2, 3)

    fmt.Println("DFS Traversal:")
    g.DFS(0, func(node int) {
    fmt.Printf("%d ", node) // Output: 0 1 2 3
 })
    fmt.Println("")
    fmt.Println("BFS Traversal")
    g.BFS(0, func(node int){
        fmt.Printf("%d ", node)
    })
    fmt.Println("")
}
