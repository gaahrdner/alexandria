# 04: Reversing the Forgetting (`04_reversing_the_forgetting.md`)

## Story

The Alexandria is growing, and the core functionality is in place.  But with increased activity comes the increased risk of mistakes.  A librarian accidentally adds the wrong book, a patron checks out a book to the wrong ID, a data entry error corrupts a record.

Dr. Vance, ever the historian, draws a parallel to ancient scribal practices. "Even the most skilled scribes made errors," he explains, "They developed techniques to correct their mistakes – scraping away ink, using special erasing fluids. We need a digital equivalent – a way to *undo* actions."

Anya Sharma nods. "He's right. We need an undo/redo system. It's essential for data integrity and user experience. Imagine the frustration of making a simple mistake and having to start all over again."

Kai Tanaka, ever focused on performance, adds his perspective. "It can't be slow. Undoing and redoing actions should be *instantaneous*. We can't have the system freezing up every time someone clicks 'undo'."

Isabelle Dubois, ever practical, points out, "We need to think carefully about *what* actions we want to be able to undo. Adding a book? Checking out a book?  We need to define the scope."

Anya turns to you. "This is where *stacks* come in. They're the perfect data structure for implementing undo/redo.  It's your task to build it."

## Your Job

Implement a stack data structure and use it to add undo/redo functionality.

### Tasks

* **Objective:** Implement a stack data structure and use it for undo/redo.

* **Theory:**

  * **Stacks: The Last-In, First-Out (LIFO) Principle**

        A stack is a linear data structure that follows the **LIFO** principle. Think of it like a stack of plates, or a stack of books (fittingly!):

    * You add new items to the *top* of the stack (this is called "pushing").
    * You remove items from the *top* of the stack (this is called "popping").
    * The *last* item you added is the *first* one you remove.

        ```bash
        |  Top Element  |  <-- Pushed last, Popped first
        |---------------|
        |  Element 2    |
        |---------------|
        |  Bottom Element|
        ```

  * **Key Operations and Time Complexity:**

    * **`Push(value interface{})`:** Adds an element to the top of the stack.  O(1) – constant time.
    * **`Pop() (interface{}, error)`:** Removes and returns the top element. Returns an error if the stack is empty. O(1).
    * **`Peek() (interface{}, error)`:** Returns the top element *without* removing it. Returns an error if empty. O(1).
    * **`IsEmpty() bool`:** Checks if the stack is empty. O(1).
    * **`Size() int`:** Returns the number of elements.  O(1).

  * **Implementation Options:**

        You can implement a stack using either:

    * **Slices:**  Go's slices are a natural fit for stacks.  You can use `append` to push and slicing to pop. This is the most common and generally recommended approach.
    * **Linked Lists:**  You *could* also use a linked list (which you just implemented!).  Pushing would be adding to the head, and popping would be removing from the head.  However, slices are usually more efficient for stacks in Go.

  * **Why Stacks for Undo/Redo?**

        The LIFO nature of stacks perfectly mirrors the undo/redo process:

        1. **Each user action is an "operation":** Adding a book, checking out a book, etc.
        2. **As the user performs actions, we "push" these operations onto the stack.**  The most recent action is always at the top.
        3. **When the user clicks "Undo," we "pop" the top operation from the stack and *reverse* it.**
        4. **For "Redo," we can use a *second* stack.**  When we undo an action, we push it onto the redo stack.  When we redo, we pop from the redo stack and push back onto the undo stack.

* **Activities:**

    1. **`internal/stackqueue/stack.go`:** Implement a stack:
        * Create a `Stack` struct.  You can use a slice (`[]interface{}`) as the underlying storage:

            ```go
            type Stack struct {
                items []interface{}
            }
            ```

        * Implement the following methods:
            * `Push(value interface{})`:  Append the value to the `items` slice.
            * `Pop() (interface{}, error)`:  Check if the stack is empty (return an error if it is).  Otherwise, get the last element of the slice, remove it from the slice (using slicing), and return the element.
            * `Peek() (interface{}, error)`:  Check if empty.  Otherwise, return the last element of the slice *without* removing it.
            * `IsEmpty() bool`:  Check if the length of the `items` slice is 0.
            * `Size() int`: Return the length of the `items` slice.

    2. **Define an `Action` struct (in `main.go`):**
        * You need a way to represent *what* action was performed, and store the necessary data to reverse the action. This needs to happen inside the `main.go`.

          ```go
          type Action struct {
              Name string // "AddBook", "CheckOutBook", "ReturnBook", "AddPatron"
              // Add fields to store the *data* needed to undo the action.  Examples:
              BookData   *book.Book     // For AddBook (store the entire book)
              PatronData *patron.Patron // For AddPatron
              PatronID   int          // For CheckOutBook/ReturnBook
              BookISBN   string       // For CheckOutBook/ReturnBook
          }
          ```

    3. **Create the stacks**
        * Add Undo and Redo stack variables to your main.go

        ```go
        undoStack := stackqueue.NewStack()
        redoStack := stackqueue.NewStack()
        ```

    4. **Update `cmd/alexandria/main.go`:**

        * **Modify `AddBook`, `AddPatron`, `CheckOutBook`, and `ReturnBook`:**
            * *Before* performing the action, create an `Action` struct and populate it with the necessary data.
            * `Push` the `Action` onto the `undoStack`.
        * **Add "Undo" and "Redo" options to the main menu.**
        * **Implement "Undo":**
            * Check if the `undoStack` is empty. If so, print a message ("Nothing to undo").
            * If not empty, `Pop` an `Action` from the `undoStack`.
            * Based on the `Action.Name`, perform the *reverse* operation:
                * `"AddBook"`:  Remove the book from the `books` slice (you'll need to find it by its ISBN – consider adding a helper function for this).
                * `"AddPatron"`: Remove the patron from the `patrons` slice (find by ID).
                * `"CheckOutBook"`:  Mark the book as available (`b.Return()`) and remove the book from the patron's checked-out list.
                * `"ReturnBook"`:  Mark the book as checked out (`b.CheckOut()`) and add the book back to the patron's checked-out list. You might need to add a helper method to your `Patron` struct to *force* a checkout without checking availability, since you're undoing a return.
            * Push the popped `Action` to the `redoStack`.
        * **Implement "Redo":**
            * Check if the `redoStack` is empty. If so, print "Nothing to redo".
            * If not empty, `Pop` an `Action` from the `redoStack`.
            * Based on the `Action.Name`, *re-perform* the original operation (essentially the same code you had before you added undo/redo).
            * Push the popped `Action` onto the `undoStack`.

    5. **Run and Test:** Compile/run. Thoroughly test undo/redo. Run unit tests: `go test ./internal/stackqueue -run TestStackLessons/Lesson04`

* **Edge Cases:**
  * Empty undo/redo stacks.
  * Multiple undos/redos in a row.
  * Undoing/redoing actions in different combinations.
