# 06: The Identity Cipher (`06_the_identity_cipher.md`)

## Story

The Neo-Alexandria is growing, and finding patrons by ID is slow.

Marcus, the security expert, is concerned. "The system is crawling. It takes *seconds* to find a patron. Unacceptable."

Anya sighs. "We're using a slice and searching linearly. It's O(n)."

Kai nods. "It's getting worse every day."

Dr. Vance suggests binary search, but Kai shakes his head. "O(log n), and we'd have to keep it sorted."

Anya has an idea. "A *hash table*! O(1) lookups on average!"

Marcus is skeptical. "Hash table? What's that?"

Anya explains: "A function – a *hash function* – maps keys (patron IDs) to indices in an array. If it's designed well, we get near-constant-time lookups."

Kai is intrigued. "O(1)? But what about collisions? Two IDs hashing to the same index?"

"We'll handle collisions," Anya says. "We'll use separate chaining – each index will be a linked list."

Isabelle adds, "Hash tables are used extensively in security and data analysis. They're fundamental for fast lookups."

## Your Job

Implement a hash table for patron lookup.

### Tasks

* **Objective:** Implement a hash table for efficient patron lookup.

* **Theory:**

  * **Hash Tables: The Core Idea**
        Store and retrieve data quickly based on a *key*. Like a dictionary: word (key) -> definition (value).

        Components:

        1. **Hash Function:** Takes a key, returns an integer (hash value). Used as the index.
            * **Good Hash Function:**
                * **Deterministic:** Same key *always* -> same hash.
                * **Uniform Distribution:** Distributes keys evenly to minimize collisions.
                * **Fast to Compute:**
        2. **Array (Buckets):** Underlying storage. Each element is a "bucket."
        3. **Collision Handling:** What happens when keys hash to the same index?
            * **Separate Chaining:** Each bucket is a linked list.  If collision, add to the linked list. (We'll use this).
            * **Open Addressing:** Try a different index (linear probing, quadratic probing, double hashing).

  * **Example (Separate Chaining):**

        ```
        // Hash Table with Separate Chaining

        Bucket 0: [Key1, Value1] -> [Key4, Value4] -> nil
        Bucket 1: [Key2, Value2] -> nil
        Bucket 2: [Key3, Value3] -> [Key5, Value5] -> [Key6, Value6] -> nil
        ...
        ```

  * **Time Complexity:**
    * **Average Case:** O(1) for `Get`, `Put`, `Delete`.
    * **Worst Case:** O(n) (if *all* keys hash to the same bucket - very unlikely).

  * **Load Factor and Resizing:**
    * **Load Factor:** `number of elements / number of buckets`
    * High load factor -> degraded performance (more collisions).
    * **Resizing:** Increase number of buckets when load factor is too high (e.g., > 0.75). Rehash all elements.

* **Activities:**

    1. **`internal/hashtable/hashtable.go`:** Implement with separate chaining:
        * `HashTable` struct:

            ```go
            type HashTable struct {
                buckets []*linkedlist.LinkedList // Array of linked lists
                size    int
                capacity int
            }
            ```

        * Constructor Function.
        * `hash(key string) int` method: Simple hash function (sum ASCII values % capacity). Experiment!
        * Methods:
            * `Put(key string, value interface{})`:
                * Calculate hash.
                * Get linked list at that index.
                * Add key-value pair (store as a struct in the linked list).
                * Increment `size`.
                * Resize if necessary, if `size/capacity` is over a threshold. Create a new `HashTable` with double capacity, add all the old values to this new table, then replace.
            * `Get(key string) (interface{}, bool)`:
                * Calculate hash.
                * Get linked list.
                * Search list for key. Return value, `true` if found; `nil`, `false` if not.
            * `Delete(key string)`:
                * Calculate hash.
                * Get linked list.
                * Search list for key. Remove if found, decrement `size`.
            * `Contains(key string) bool`: Like `Get`, but only returns `true`/`false`.
            * `Size() int`: Returns `size`.

    2. **Update `cmd/alexandria/main.go`:**
        * Replace patron slice (`[]patron.Patron`) with hash table (`*hashtable.HashTable`).
        * "Add Patron": Use `ht.Put(strconv.Itoa(newPatron.ID), newPatron)` (ID as string key, `Patron` as value).
        * "List Patrons," "Check Out/Return Book": Use `ht.Get(patronIDStr)`. Handle `bool` return.
        * Update logic that iterated over patron slice. You may want to skip this for now.

    3. **Run and Test:** Compile/run. Test. Run unit tests: `go test ./internal/hashtable -run TestHashTableLessons/Lesson06`

* **Edge Cases:**
  * Empty hash table.
  * Deleting non-existent key.
  * Collisions.
  * Resizing.
