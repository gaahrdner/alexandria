# 02: Guardians of Knowledge (`02_guardians_of_knowledge.md`)

## Story

The digital shelves of the Alexandria are beginning to fill with books, thanks to your work on the `Book` struct and the basic command-line interface. But a library isn't just about books; it's about the *people* who use them.

Anya Sharma calls a meeting in the main hall, which is slowly taking shape amidst the ongoing construction. Dr. Vance stands beside her, while Kai Tanaka leans against a nearby pillar, arms crossed, a skeptical look on his face. Isabelle Dubois sits quietly, cleaning dust off a recovered hard drive.

"We've got books," Anya announces, "now we need *patrons*. We need a way to represent them in the system, track who has which book, and manage the checkout and return process."

Dr. Vance nods enthusiastically. "Indeed! The guardians of knowledge, the seekers of wisdom! We must make their interaction with the library as seamless as possible."

Kai pushes himself off the pillar. "And efficient," he adds. "We can't have the system bogging down when hundreds, thousands, maybe even millions of patrons try to access the library simultaneously."

Isabelle looks up. "And we need to handle errors. What if someone tries to return a book they never checked out? What if the system crashes mid-transaction?"

Anya turns to you. "That's where you come in. We need you to implement the `Patron` struct and the core checkout/return logic. Think about how a patron interacts with a book. They check it out, they return it. We need methods for those actions, on both the `Book` and `Patron` structs."

## Your Job

Implement patron representation, add checkout/return functionality, and update the command-line interface.

### Tasks

* **Objective:** Implement patron representation, add checkout/return functionality, and update the command-line interface.

* **Activities:**

    1. **`internal/patron/patron.go`:**
        * Define the `Patron` struct. Consider what information you need to store about each patron: their name, a unique ID, and a list of the books they currently have checked out.
        * Create a `NewPatron` constructor function, similar to `NewBook`.
        * Implement `CheckOutBook(isbn string)`: This method should add a book's ISBN to the patron's list of checked-out books. Consider: Should a patron be able to check out the same book twice?
        * Implement `ReturnBook(isbn string) error`: This method should remove a book's ISBN from the patron's list. It should return an `error` if the patron *doesn't* have that book checked out.
        * Implement `HasCheckedOutBook(isbn string) bool`: A helper method to check if a patron has a specific book.
        * Implement a `String()` method for easy display of patron information, including their checked-out books.

    2. **Update `internal/book/book.go`:**
        * Add a `CheckOut()` method. This should simply update the book's `Available` status.
        * Add a `Return()` method. This should update the book's `Available` status.
        * Add an `IsAvailable() bool` method.

    3. **Update `cmd/alexandria/main.go`:**
        * Add menu options:
            * Add Patron: Prompt for the patron's name and generate a unique ID.
            * List Patrons: Display the details of all registered patrons.
            * Check Out Book: Prompt for the patron's ID and the book's ISBN.  Find both the patron and the book. Call the appropriate `CheckOutBook` methods on both.
            * Return Book: Prompt for the patron's ID and the book's ISBN. Find both, and call the appropriate `ReturnBook` methods.
        * Handle cases where the patron or book is not found.
        * Display appropriate error messages returned by `ReturnBook`. Use `errors.Is` if you use custom errors.
        * Use `strings.EqualFold` when comparing entered ISBNs to those in the system, case-insensitively.

    4. **Run and Test:** Compile and run. Extensively test adding patrons, checking out books, returning books, and listing both books and patrons. Use the subtest running capabilities:
          * `go test ./internal/patron -run TestPatronLessons/Lesson02`
          * `go test ./internal/book -run TestLessons/Lesson02`

* **Edge Cases to Consider in tests**
  * `internal/patron/patron_test.go`
    * Empty Name
    * ID 0
    * Return book not checked out
    * Check out same book twice
  * `internal/book/book_test.go`
    * Already checked out
    * Already returned
