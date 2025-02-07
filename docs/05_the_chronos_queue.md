# 05: The Chronos Queue (`05_the_chronos_queue.md`)

## Story

Tension hangs heavy in the air. The initial euphoria of the Alexandria's launch has faded, replaced by the relentless pressure of managing a growing digital library and an ever-increasing influx of data.

Isabelle, her face etched with exhaustion, slams a data chip onto the table. "Another corrupted archive," she announces, her voice tight. "This one's from a pre-Forgetting university. Thousands of research papers, potentially lost forever."

Anya looks up from her console, her brow furrowed. "We're falling behind, Isabelle. The data recovery is piling up faster than we can process it."

"It's not just the volume," Isabelle retorts, "it's the *priority*. We need a system to manage which archives get processed first. Some are more critical than others – medical research, historical records, engineering schematics..."

Elias, who has been listening quietly, strokes his beard. "She's right. We need a queue. A way to prioritize tasks, to ensure that the most important data is recovered first."

Kai, who has been unusually silent, speaks up. "A queue, yes. But what *kind* of queue? We need to be precise. A simple first-in, first-out queue might not be enough. What if a critical archive arrives *after* a less important one?"

Anya raises an eyebrow. "You're thinking of a priority queue, Kai? We could assign a priority level to each archive and process them accordingly."

"Potentially," Kai admits, "but let's start with the basics. A standard FIFO queue is a good foundation. We can always add priority later if needed." He looks pointedly at Anya. "And let's not overengineer it from the start. Remember the K.I.S.S. principle – Keep It Simple, Stupid."

Anya bristles slightly at Kai's tone. Their working relationship has been strained lately. Anya feels Kai is overly critical and dismissive of her ideas, while Kai believes Anya sometimes rushes into complex solutions without considering simpler alternatives.

"Fine," Anya says, her voice clipped. "A standard queue it is. But we need to implement it efficiently. And we *will* need to revisit the priority issue later."

The unspoken tension between them fills the room. Elias, sensing the friction, tries to lighten the mood. "Think of it as a time machine, Anya, Kai," he says with a smile. "A queue that brings the past back to life, one piece at a time."

You, the ever-reliable engineer, step in to defuse the situation. "I can implement the queue," you say. "We'll use a slice-based implementation with a circular buffer for efficiency. It'll be fast and reliable."

## Your Job

Implement a queue data structure (using a slice and circular buffer) and understand its applications. Although it's not integrated directly into the main library features *yet*, this lesson focuses on understanding the queue data structure itself.

### Tasks

* **Objective:** Implement a queue and understand its properties.

* **Theory:**

  * **Queues: The First-In, First-Out (FIFO) Principle**

        A queue is a linear data structure that follows the **FIFO** principle. Think of it like a line at a store, or a waiting list:

    * New elements are added to the *rear* (or *tail*) of the queue (this is called "enqueuing").
    * Elements are removed from the *front* (or *head*) of the queue (this is called "dequeuing").
    * The *first* element added is the *first* one to be removed.

        ```bash
        Dequeue  <-- [ Element 1 | Element 2 | Element 3 ] <-- Enqueue
         (Head)                                            (Tail)
        ```

  * **Key Operations and Time Complexity:**

    * **`Enqueue(value interface{})`:** Adds an element to the rear. O(1) (for efficient implementations).
    * **`Dequeue() (interface{}, error)`:** Removes and returns the front element. Returns an error if empty. O(1) (for efficient implementations).
    * **`Peek() (interface{}, error)`:** Returns the front element *without* removing it. Returns an error if empty. O(1).
    * **`IsEmpty() bool`:** Checks if the queue is empty. O(1).
    * **`Size() int`:** Returns the number of elements. O(1).

  * **Implementation Options:**

    * **Linked Lists:** You *could* use a linked list (with a tail pointer). Enqueuing would be adding to the tail, and dequeuing would be removing from the head. This is a perfectly valid approach.
    * **Slices with Circular Buffer (Recommended):** This is a more efficient approach for queues in Go.  It uses a slice, but with a clever trick to avoid shifting elements when dequeuing.

      * **Circular Buffer:** Imagine the slice as a circle.  We have two indices: `head` and `tail`.
        * `head`:  Points to the index of the *next* element to be dequeued.
        * `tail`: Points to the index where the *next* element will be enqueued.

                ```
                [  ,  , 4, 5, 6,  ,  ]  (Slice with capacity 7)
                 ^     ^
                 |     |
                Head  Tail
                ```

        * When we enqueue, we add the element at the `tail` index and increment `tail`.
        * When we dequeue, we return the element at the `head` index and increment `head`.
        * **Crucially:** When `head` or `tail` reach the end of the slice, they "wrap around" to the beginning (using the modulo operator `%`). This is what makes it "circular."

                ```go
                // Example: Enqueuing a new element (assuming tail is not at capacity)
                tail = (tail + 1) % capacity

                // Example: Dequeuing
                head = (head + 1) % capacity
                ```

        * You'll also need to handle resizing the slice when it becomes full.

* **Activities:**

    1. **`internal/stackqueue/queue.go`:** Implement a queue using a slice and a circular buffer:
        * Create a `Queue` struct:

            ```go
            type Queue struct {
                items []interface{}
                head  int
                tail  int
                size  int
                cap int //capacity
            }
            ```

        * Create a constructor function:

            ```go
            func NewQueue() *Queue{
                return &Queue{
                    items: make([]interface{}, 8), //Initial Capacity
                    head: 0,
                    tail: 0,
                    size: 0,
                    cap: 8,
                }
            }
            ```

        * Implement the following methods:
            * `Enqueue(value interface{})`: Add to the tail, handle wrapping, and resize if necessary (double the capacity when full).
            * `Dequeue() (interface{}, error)`: Remove from the head, handle wrapping, return an error if empty.
            * `Peek() (interface{}, error)`: Return the head element without removing it, return an error if empty.
            * `IsEmpty() bool`: Check if the queue is empty.
            * `Size() int`: Return the number of elements.

    2. **Run and Test:** Compile and run. Test all queue operations.  Run unit tests: `go test ./internal/stackqueue -run TestQueueLessons/Lesson05`

* **Edge Cases:**
  * Empty queue (dequeuing/peeking).
  * Full queue (enqueuing and resizing).
  * Wrapping around the circular buffer.
