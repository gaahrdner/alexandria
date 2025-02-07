# 07: The Knowledge Tree (`07_the_knowledge_tree.md`)

## Story

Patron lookups are fast, but finding books by ISBN is slow.

Dr. Vance suggests a hierarchy. "Like branches of a tree. Each ISBN guides us."

Anya: "A binary search tree! Smaller ISBNs left, larger right. Efficient searching."

Kai: "Efficient *on average*. A degenerate tree is as slow as a linear search."

Anya: "True. We'll watch for balance. We could use self-balancing trees later. ISBN distribution should help."

Isabelle: "A BST also gives us *order*. We can list books in ISBN order easily."

Anya: "Exactly. Structure *and* speed."

## Your Job

Implement a Binary Search Tree (BST) for ISBN lookup.

### Tasks

* **Objective:** Implement a BST for efficient book lookup by ISBN.

* **Theory:**

  * **Binary Search Trees (BSTs):**
    * Each *node* has at most two children: *left* and *right*.
    * **BST invariant:**
      * Left subtree keys < node's key.
      * Right subtree keys > node's key.

        ```bash
            50  <-- Root
           /  \
          30   70
         /  \ /  \
        20  40 60 80   (Example BST - ISBNs as keys)
        ```

  * **Key Operations:**
    * **`Insert(key int, value interface{})`:**
      * Start at root.
      * Compare key with current node's key.
      * Smaller: go left. Larger: go right.
      * Repeat until empty spot, insert new node.
      * Average: O(log n). Worst: O(n) (degenerate tree).
    * **`Search(key int) (interface{}, bool)`:**
      * Start at root.
      * Compare key with current node's key.
      * Equal: return value, `true`.
      * Smaller: go left. Larger: go right.
      * Repeat until found or reach `nil` (return `nil`, `false`).
      * Average: O(log n). Worst: O(n).
    * **`Delete(key int)`:** *Most complex*. Three cases:
      * **Leaf (no children):** Remove.
      * **One child:** Replace with child.
      * **Two children:** Find *inorder successor* (smallest in right subtree) or *predecessor* (largest in left subtree). Replace node's key/value with successor/predecessor's, then *delete* the successor/predecessor.
      * Average: O(log n). Worst: O(n).
    * **`Min() (int, error)`:** Leftmost node.
    * **`Max() (int, error)`:** Rightmost node.
    * **`InOrderTraversal() []int`:** Visits nodes in ascending order.

  * **Balanced vs. Unbalanced:**
    * **Balanced:** Height ≈ log(n). Good performance (O(log n)).
    * **Unbalanced:** Height close to n. Worst case (degenerate tree): O(n).
    * Self-balancing trees (AVL, Red-Black) maintain balance (more complex).

* **Activities:**

    1. **`internal/bst/bst.go`:** Implement a BST:
        * `Node` struct: `Key` (`int`), `Value` (`interface{}`), `Left` (`*Node`), `Right` (`*Node`).
        * `BST` struct: `Root` (`*Node`).
        * Methods:
            * `Insert(key int, value interface{})`
            * `Search(key int) (interface{}, bool)`
            * `Delete(key int)` (Most challenging!)
            * `Min() (int, error)`
            * `Max() (int, error)`
            * `InOrderTraversal() []int` (For testing)
            * Consider recursive helper functions.

    2. **Update `internal/book/book.go`:**
        * Method to convert ISBNs to integers.

    3. **Update `cmd/alexandria/main.go`:**
        * Replace book slice (`[]book.Book`) with BST (`*bst.BST`).
        * "Add Book": Use `bst.Insert(bookISBNInt, newBook)`.
        * "Return/Check Out Book": Use `bst.Search(isbnInt)`. Handle `bool` return.
        * Temporarily remove or modify "List Books" (use `InOrderTraversal`).

    4. **Run and Test:** Compile/run. Test adding, searching, deleting. Run unit tests: `go test ./internal/bst -run TestBSTLessons/Lesson07`

* **Edge Cases:**
  * Empty tree.
  * Inserting duplicates (decide: allow or disallow).
  * Deleting non-existent key.
  * Deleting nodes with 0, 1, or 2 children.
