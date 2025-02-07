# 13: Unlocking the Archives (`13_unlocking_the_archives.md`)

## Story

The Neo-Alexandria is a resounding success, a beacon of knowledge in the post-Forgetting world.  However, as the library grows, a new need arises.  Dr. Vance wants to generate reports: lists of all patrons, lists of patrons with overdue books, and other summaries.  The current system, while excellent for *finding* a specific patron by ID, doesn't provide an easy way to *iterate* over all patrons.

Anya realizes the limitation. "The hash table is designed for fast lookups by key," she explains, "but it doesn't inherently provide a way to access all the key-value pairs in a sequential manner. We need an *iterator*."

Kai, ever practical, adds, "And it needs to be efficient. We don't want to create a whole copy of the patron data just to iterate over it."

Isabelle, drawing on her data analysis experience, suggests, "An iterator is like a cursor, moving through the data one element at a time. It allows us to process each patron record without having to load everything into memory at once."

Anya nods. "Exactly. We'll add an `Iterator()` method to our `HashTable` struct. This method will return a *channel* of key-value pairs.  This allows us to process the patrons concurrently and efficiently."

## Your Job

Add an iterator to the `HashTable` to enable efficient iteration over all patrons.

### Tasks

* **Objective:** Implement an iterator for the `HashTable`.

* **Theory:**

  * **Iterators:** An iterator is a design pattern that provides a way to access the elements of a collection sequentially *without* exposing the underlying representation of the collection.
  * **Channels in Go:** Go's channels are a powerful mechanism for communication and synchronization between goroutines (concurrently executing functions).  We can use a channel to create an iterator that yields key-value pairs one at a time.
  * **Why Channels for Iteration?**
    * **Concurrency:**  Using a channel allows us to iterate over the hash table concurrently with other operations.
    * **Efficiency:**  We don't need to create a separate copy of all the data to iterate over it.  The channel acts as a stream, providing elements on demand.
    * **Clean Separation:** The iteration logic is encapsulated within the `Iterator()` method, keeping the main hash table implementation clean.

* **Activities:**

    1. **Update `internal/hashtable/hashtable.go`:**
        * Add an `Iterator() <-chan struct {Key string; Value interface{}}` method to the `HashTable` struct. This method should:
            * Create a *buffered* channel of a struct containing a Key, and a Value: `make(chan struct {Key string; Value interface{}}, bufferSize)`.  The buffer size controls how many key-value pairs are pre-fetched.  A small buffer (e.g., 10) is usually sufficient.
            * Start a *goroutine* (using the `go` keyword).
            * Inside the goroutine:
                * Iterate over the `buckets` array.
                * For each non-empty bucket (linked list):
                    * Iterate over the linked list.
                    * For each node in the linked list, send the key-value pair (as a struct) to the channel.
                * *Close* the channel after iterating over all buckets (using `close(channel)`). This signals to the receiver that no more data will be sent.
            * Return the channel.

            ```go
            // Example (Conceptual - fill in the details)
            func (ht *HashTable) Iterator() <-chan struct{Key string; Value interface{}} {
                ch := make(chan struct{Key string; Value interface{}}, 10) // Buffered channel

                go func() {
                    defer close(ch) // Close the channel when done

                    for _, list := range ht.buckets {
                        if list != nil {
                            current := list.Head
                            for current != nil {
                                kv := current.Value.(struct{Key string; Value interface{}}) // Type assertion
                                ch <- kv // Send the key-value pair to the channel
                                current = current.Next
                            }
                        }
                    }
                }()

                return ch
            }
            ```

    2. **Update `cmd/alexandria/main.go`:**
        * Modify the "List Patrons" logic to use the iterator:

            ```go
            for kv := range patrons.Iterator() {
                patron := kv.Value.(*patron.Patron) // Type assertion
                fmt.Println(patron.String())
                fmt.Println("---")
            }
            ```

    3. **Run and Test:** Compile/run. Test listing patrons using the iterator.  Run unit tests (including the `TestHashTable_Iterator` test you added previously): `go test ./internal/hashtable -run TestHashTable_Iterator`

* **Edge Cases:**
  * Empty hash table.
  * Hash table with multiple elements in the same bucket.
