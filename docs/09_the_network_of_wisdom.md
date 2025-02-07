# 09: The Network of Wisdom (`09_the_network_of_wisdom.md`)

## Story

The Neo-Alexandria is a vast repository, but more than just individual books. Connections exist: sequels, related topics, same author.

Isabelle proposes a new idea: "We need to map the relationships. A network, a web of knowledge. Like গোয়েন্দা work – connecting the dots."

Anya: "A graph! Books as nodes, relationships as edges."

Kai: "What *kind* of relationships? Sequels? Topics? And how to represent the graph? Adjacency matrix? Adjacency list?"

Dr. Vance: "Imagine a network of scrolls, linked by threads of meaning."

Anya: "We could recommend books. If they like one, suggest connected ones."

## Your Job

Implement a graph to represent relationships between books.

### Tasks

* **Objective:** Implement a graph for book relationships.

* **Theory:**

  * **Graphs:** Represent relationships.
    * **Vertices (Nodes):** Objects (books).
    * **Edges:** Relationships.

        Edges can be:
    * **Directed:** One-way (Book A -> sequel -> Book B).
    * **Undirected:** Two-way (Book A <-> similar topic <-> Book B).
    * **Weighted:** Edge has a value (similarity score).
    * **Unweighted:** Simple connection.

        ```bash
        Undirected Graph:
        A -- B
        |    |
        C -- D

        Directed Graph:
        A -> B
        ^    |
        |    v
        C <- D
        ```

  * **Representations:**
    * **Adjacency Matrix:** 2D array. `matrix[i][j] = 1` if edge between vertex i and j, 0 otherwise. Good for *dense* graphs.

            ```bash
               A  B  C  D
            A  0  1  1  0
            B  1  0  0  1
            C  1  0  0  1
            D  0  1  1  0
            ```

    * **Adjacency List:** Each vertex stores a list of its neighbors.  More space-efficient for *sparse* graphs (few edges).

            ```bash
            A: [B, C]
            B: [A, D]
            C: [A, D]
            D: [B, C]
            ```

  * **Traversals:**
    * **Depth-First Search (DFS):** Explore as far as possible along each branch before backtracking. Uses a stack.
    * **Breadth-First Search (BFS):** Explore all neighbors at current depth before moving to next level. Uses a queue.

* **Activities:**

    1. **`internal/graph/graph.go`:** Implement a graph (choose *adjacency list* or *adjacency matrix* - adjacency list is generally recommended):
        * `Graph` struct:
            * **Adjacency List:**

                ```go
                type Graph struct {
                    vertices map[string][]string // Key: Book ISBN, Value: List of connected ISBNs
                }
                ```

            * **Adjacency Matrix:** (Less recommended)

                ```go
                type Graph struct {
                    vertices []string // List of ISBNs
                    edges    [][]bool // 2D array
                }
                ```

        * Methods:
            * `AddVertex(isbn string)`: Add new book.
            * `AddEdge(isbn1, isbn2 string)`: Add *undirected* edge.
            * `GetNeighbors(isbn string) []string`: Return list of connected books.
            * `BFS(startISBN string) []string`: Breadth-First Search.
            * `DFS(startISBN string) []string`: Depth-First Search.

    2. **Update `cmd/alexandria/main.go`:**
        * Add menu options:
            * "Add Book Relationship":  Prompt for two ISBNs, add undirected edge.
            * "List Similar Books": Prompt for ISBN, use BFS/DFS to find connected books.

    3. **Run and Test:** Compile/run. Test adding vertices/edges, listing neighbors. Run unit tests: `go test ./internal/graph -run TestGraphLessons/Lesson09`

* **Edge Cases:**
  * Empty graph.
  * Duplicate vertices/edges.
  * Self-loops.
  * Non-existent vertex.
