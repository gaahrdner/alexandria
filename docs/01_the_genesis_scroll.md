# 01: The Genesis Scroll (`01_the_genesis_scroll.md`)

## Story

The flickering holographic display illuminates Anya Sharma's face as she gestures towards the skeletal framework of code projected in the air. "Alright," she begins, her voice echoing slightly in the cavernous, still-under-construction server room, "we're starting with the foundation: representing books. Think of each book as a digital scroll – title, author, ISBN, publication year, and whether it's available. We'll use a Go `struct` for this."

You nod, familiar with structs. Anya continues, "Structs are our basic building blocks. They're like defining the schema for each scroll – the essential data we need to track."

Dr. Elias Vance, his eyes twinkling with historical perspective, interjects, "Indeed! Just as the ancient Library of Alexandria meticulously cataloged its scrolls, we must carefully structure our digital records. Each field in our `Book` struct is like a label on a scroll, identifying its contents."

Anya taps on her wrist-mounted console. "We'll need fields for the title, author, ISBN, and publication year. And a boolean to track availability. Consider the types carefully – strings for text, an integer for the year..."

```go
// Example: A struct to represent a similar concept
type Document struct {
    Name        string
    Author      string
    ID          string
    CreatedYear int
    IsPublic    bool
}
```

"See? Simple, clear. Now, we need a way to store these books. We could use an array, but..." Anya pauses, and Kai Tanaka, leaning against a server rack, jumps in.

"Arrays are fixed-size," Kai states, his voice a low, resonant hum. "We don't know how many books we'll eventually recover. We need something dynamic, something that can grow. That's where slices come in."

Anya nods in agreement. "Exactly. Slices are Go's dynamic arrays. They're built on top of arrays, but they can resize automatically. Think of them as magical shelves that expand as we add more books."

She draws a diagram in the air, the holographic projector bringing it to life:

```bash
Slice: [ptr] -> [ Element 1 | Element 2 | Element 3 | ... | Element N ]
        len      [  Capacity (may be larger than len)          ]
```

"A slice has three parts: a pointer to the underlying array, the length (number of elements currently used), and the capacity (the total allocated space). Go handles the resizing for us, which is incredibly convenient. Appending is usually O(1), but becomes O(n) when the capacity is reached and a new, larger array must be allocated and the data copied."

Dr. Vance smiles. "Like adding new wings to the library as our collection grows!"

Anya continues, "To make creating new books easier, we'll create a constructor function. This function will take the book's details as arguments and return a pointer to a newly created Book struct. Think of it as a function that takes the raw data and 'binds' it into a digital scroll."

```go
// Example Constructor Function (for a different struct)
func NewDocument(name, author, id string, createdYear int) *Document {
    return &Document{ /* ... initialization ... */ }
}
```

"Using a pointer is more efficient – we avoid copying the entire struct. Also, consider setting the initial availability of the book to true within this constructor."

"And a String() method is essential," Isabelle Dubois calls out, "for debugging and displaying book information."

Anya agrees. "Absolutely. The String() method should return a nicely formatted string representation of the book's data. You could use fmt.Sprintf to create the string, or a string builder, to avoid extra allocations."

```go
//Example String() method
func (d *Document) String() string{
    return fmt.Sprintf("Name %s\nAuthor: %s\n", d.Name, d.Author)
}
```

"Now, for the user interface," Anya says, turning to you. "We'll start with a command-line interface. We'll use the fmt package for printing, and bufio for reading input, specifically, bufio.Scanner to read line by line."

"Think of fmt as our announcer," Dr. Vance explains, "and bufio as our scribe."

"And strconv," Kai adds, "is our translator. String input needs conversion to numbers."

Anya emphasizes, "And remember error handling. Check for errors and handle them. What if the user enters text instead of a number for the year? The application shouldn't crash."

## Your job

Set up the project, represent books, and create a command-line menu.

### Tasks

* **Objective:** Set up the project, represent books, and create a command-line menu.

* **Activities:**

    1. **Project Structure:** (note the new files)

        ```bash
        alexandria/
        ├── cmd/
        │   └── alexandria/
        │       └── main.go
        ├── internal/
        │   ├── book/
        │   │   ├── book.go
        │   │   └── book_test.go
        │   ├── patron/
        │   │   ├── patron.go
        │   │   └── patron_test.go
        ├── go.mod
        └── go.sum
        ```

    2. **`internal/book/book.go`:**
        * Define the `Book` struct. Choose appropriate types for each field.
        * Create a `NewBook` constructor function. It should take arguments for each field and return a *pointer* to a new `Book` instance.
        * Create a String() Method to return a formatted string of the Book.

    3. **`cmd/alexandria/main.go`:**
        * Create a loop that presents a menu: Add Book, List Books, Exit.
        * Use `fmt` and `bufio` for I/O. Use a `bufio.Scanner` to read user input line by line.
        * Use `strconv` to convert string input to numbers (e.g., for the publication year).
        * Implement "Add Book":
            * Prompt the user for the title, author, ISBN, and publication year.
            * Use `NewBook` to create a new `Book` instance.
            * Append the new book to a slice of books.
        * Implement "List Books":
            * Iterate over the slice of books.
            * Print the details of each book using its `String()` method.
        * Implement "Exit": Terminate the program.
        * Handle errors gracefully (e.g., invalid input).

    4. **Run and Test:** Build and run. Add and list books. Run tests: `go test ./internal/book -run TestLessons/Lesson01`

* **Edge Cases:**
* Empty strings for title, author, and ISBN.
 A publication year of 0.

"This is just the beginning," Anya says, "but it's a crucial first step. Let's build this foundation strong!"
