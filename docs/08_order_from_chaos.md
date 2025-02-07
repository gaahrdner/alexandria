# 08: Order from Chaos (`08_order_from_chaos.md`)

## Story

The Alexandria is now equipped with a BST for efficient ISBN lookups.  However, users often want to browse the library's collection in different ways: by title, by author, by publication year. The current system lacks a robust way to sort the books.

Lillian Vance, drawing on her years of experience as a librarian, explains the need. "Patrons don't always know the exact ISBN they're looking for," she says, "Sometimes they want to see all books by a particular author, or all books published in a certain year. We need to provide different ways to *sort* the collection."

Anya agrees. "Listing all books in ISBN order using the BST's in-order traversal is fine, but it's not enough. We need to offer more flexibility."

Kai, ever the performance enthusiast, sees this as an opportunity to delve into the world of sorting algorithms. "This is where things get interesting," he says with a grin. "There are *many* ways to sort data, each with its own strengths and weaknesses. We'll need to explore different algorithms and see which ones perform best for our needs."

He starts sketching on the whiteboard, outlining various sorting algorithms: Bubble Sort, Selection Sort, Insertion Sort, Merge Sort, Quick Sort, Heap Sort. "We've got the simple but slow ones," he explains, pointing to the first three, "and the more complex but efficient ones," he indicates the latter three.

"We should aim for O(n log n) performance," Kai continues, "That's the sweet spot for general-purpose sorting.  Merge Sort and Heap Sort guarantee that, while Quick Sort is usually faster in practice but has a worst-case O(n^2) scenario."

Isabelle raises a practical point. "We also need to consider *stability*.  If we sort by author and then by title, will books by the same author remain in their original order within the title sort?"

Anya nods. "Good point, Isabelle.  A *stable* sorting algorithm preserves the relative order of equal elements.  Merge Sort is stable, but Quick Sort and Heap Sort are typically not."

The challenge is set: implement a variety of sorting algorithms and allow users to choose how they want to view the library's collection.

## Your Job

Implement several sorting algorithms and integrate them into the Alexandria.

### Tasks

* **Objective:** Implement and compare different sorting algorithms.

* **Theory:**

  * **Sorting Algorithms:** (Review the descriptions, diagrams, and time complexities of the following algorithms from the earlier lesson plan responses.  Focus on understanding *how* each algorithm works, its best/average/worst-case performance, and whether it's stable or not.)
    * **Bubble Sort:** Simple, but inefficient (O(n^2)).
    * **Selection Sort:**  Also simple, but still O(n^2).
    * **Insertion Sort:**  Efficient for small or nearly sorted data (O(n) best case), but O(n^2) on average.
    * **Merge Sort:**  Divide and conquer algorithm.  Efficient and stable (O(n log n) guaranteed).
    * **Quick Sort:**  Divide and conquer algorithm.  Often faster than Merge Sort in practice (O(n log n) average), but O(n^2) worst case.  Typically not stable.
    * **Heap Sort:**  Uses a binary heap data structure.  Guaranteed O(n log n) performance, but often slightly slower than Quick Sort.  Not stable.

  * **Comparison-Based Sorting:** All of the above algorithms are *comparison-based*: they compare elements to determine their order.

  * **Stability:** A sorting algorithm is *stable* if it preserves the relative order of equal elements.

  * **In-Place vs Out-of-Place:**
    * **In-Place** Uses a constant amount of extra memory.
    * **Out-of-Place** Uses extra memory proportional to input size.

* **Activities:**

    1. **`internal/sorting/sorting.go` (and separate files for each algorithm):**
        * Create separate `.go` files for each sorting algorithm (e.g., `bubble.go`, `merge.go`, `quick.go`, `heap.go`, `selection.go`, `insertion.go`).
        * Implement each of the six sorting algorithms listed above. Each function should take a slice of `book.Book` structs and a `less` function as input:

            ```go
            func BubbleSort(books []book.Book, less func(b1, b2 book.Book) bool) {
                // ... implementation ...
            }

            // The 'less' function defines the sorting order.  For example:
            func lessByTitle(b1, b2 book.Book) bool {
                return b1.Title < b2.Title
            }

            func lessByAuthor(b1, b2 book.Book) bool {
                return b1.Author < b2.Author
            }
            ```

            This allows you to sort by different fields (title, author, ISBN, publication year) using the *same* sorting algorithm implementations.

    2. **Update `cmd/alexandria/main.go`:**
        * Add a "List Books" option with sub-options for choosing the sorting criteria (title, author, ISBN, publication year) and the sorting algorithm.
        * When listing books:
            1. Retrieve all books from the BST. You'll need a way to get all books from the BST into a slice. You could modify your InOrderTraversal to achieve this.
            2. Create a copy of the books slice (so you don't modify the original order in the BST).
            3. Call the selected sorting function with the copied slice and the appropriate `less` function.
            4. Print the sorted slice.

    3. **Run and Test:** Compile/run. Test listing books with different sorting criteria and algorithms. Run unit tests: `go test ./internal/sorting -run TestSortingLessons/Lesson08`. Also run benchmarks to compare the performance of the different algorithms: `go test -bench=. -benchmem ./internal/sorting`

* **Edge Cases:**
  * Empty slice.
  * Slice with one element.
  * Already sorted slice.
  * Reverse sorted slice.
  * Slice with duplicate values.
