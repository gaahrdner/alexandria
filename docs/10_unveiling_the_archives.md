# 10: Unveiling the Archives (`10_unveiling_the_archives.md`)

## Story

The Alexandria is functional, but finding specific information within its vast collection can still be cumbersome. Searching by ISBN (using the BST) is efficient, but patrons often don't know the ISBN. They need to be able to search by *title* and, ideally, with partial matches.

Anya Sharma calls a meeting. "We've built the structure," she says, "but the library isn't truly accessible until patrons can easily find what they're looking for. We need to implement title search."

Kai Tanaka, ever the performance watchdog, is wary. "Title search... that sounds like a linear search through the entire book collection.  O(n).  That won't scale."

"We can't use the BST for title search directly," Anya acknowledges, "because the BST is ordered by ISBN.  We could build a *separate* BST ordered by title, but that would mean duplicating the book data, which is inefficient and could lead to inconsistencies."

Dr. Vance suggests, "Perhaps we can adapt existing search techniques. A simple linear search, but optimized for our needs?"

Isabelle Dubois, who has experience with information retrieval, offers a suggestion. "We could use a technique called *prefix search*.  If the user types 'Hist', we show all books whose titles start with 'Hist', like 'History of Rome,' 'Historical Linguistics,' etc.  That would be more efficient than searching the entire title string every time."

Marcus Cole, ever concerned about security, adds a note of caution. "Whatever we do, we need to be mindful of potential vulnerabilities.  Search queries could be manipulated to cause performance issues or even crashes.  We need to sanitize and validate user input carefully."

Anya nods. "Good point, Marcus. Security is paramount." She turns to you. "Your task is to implement title search, focusing on efficiency and security. Start with a linear search, but think about how we can optimize it. Consider Isabelle's suggestion about prefix search."

## Your Job

Implement title search (partial match) within the Alexandria.

### Tasks

* **Objective:** Implement title search (partial match).

* **Theory:**

  * **Linear Search (for now):**  For this initial implementation, we'll use a linear search through the slice of books. While not the most efficient for very large datasets, it's a good starting point and allows us to focus on the integration with the existing system.
    * Iterate through the slice of books.
    * For each book, compare the search query (title) with the book's title.
    * Use `strings.Contains` or `strings.HasPrefix` (for prefix search) for the comparison. Consider case-insensitivity (`strings.ToLower`).
    * If a match is found, add the book to a result list.
    * Return the list of matching books.

  * **Optimization Considerations (Future):**
    * **Trie (Prefix Tree):**  A trie is a tree-like data structure specifically designed for efficient prefix search.  This would be a significant improvement for a large library.
    * **Inverted Index:**  An inverted index maps words to the documents (books) that contain them.  This is used by search engines and would be highly effective for full-text search.
    * **Database:**  For a very large collection, a dedicated database (like PostgreSQL, MySQL, or Elasticsearch) would provide powerful indexing and search capabilities.

  * **Security Considerations:**
    * **Input Sanitization:**  Always sanitize user input to prevent potential injection attacks.  Limit the length of search queries, escape special characters, and consider using a whitelist of allowed characters.
    * **Rate Limiting:**  Limit the number of search requests per user per unit of time to prevent denial-of-service attacks.

* **Activities:**

    1. **Update `cmd/alexandria/main.go`:**
        * Add a "Search by Title" option to the main menu.
        * Prompt the user for a search query (title or part of a title).
        * Implement a `searchByTitle` function (you can put this in `book.go` or `main.go` – your choice):
            * This function should take the search query as input.
            * Iterate over all books in the collection (You'll have to retrieve the books from your BST using in-order traversal.)
            * For each book, use `strings.Contains` or `strings.HasPrefix` (and `strings.ToLower` for case-insensitivity) to check if the book's title matches the query.
            * If it matches, add the book to a result list.
            * Return the list of matching books.
        * Display the search results to the user.
        * **Important:** Sanitize the user input *before* performing the search.

    2. **Run and Test:** Compile/run. Test title search with various queries, including partial matches and edge cases.  You'll need to add some tests to a new `TestLessons` TestSuite in main_test.go (create this file in the cmd/alexandria folder)

* **Edge Cases:**
  * Empty search query.
  * Search query with special characters.
  * Search query that matches no books.
  * Search query that matches many books.
