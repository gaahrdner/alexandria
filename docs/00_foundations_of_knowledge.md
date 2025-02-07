# 00: Foundations of Knowledge (`00_foundations_of_knowledge.md`)

Before we embark on the journey of building the Alexandria, let's establish some fundamental concepts that will be crucial throughout this project. Think of these as the core principles, the building blocks upon which all our data structures and algorithms will be constructed.

## Big O Notation: Measuring Efficiency

In the post-Forgetting world, resources are scarce, and efficiency is paramount.  We can't afford to waste time or processing power on slow, inefficient algorithms.  That's where **Big O notation** comes in.

Big O notation is a way to describe the *asymptotic* performance of an algorithm – how its runtime or memory usage grows as the input size grows. It focuses on the *dominant* term, ignoring constant factors and lower-order terms.  It's like describing the *overall shape* of a growth curve, rather than its precise values.

**Common Big O Notations (from fastest to slowest):**

* **O(1) - Constant Time:** The algorithm takes the same amount of time regardless of the input size.  Think of looking up a value in a hash table (on average) or accessing an element in an array by its index.
* **O(log n) - Logarithmic Time:** The runtime grows logarithmically with the input size.  Each step effectively halves the remaining data.  Binary search is a classic example.
* **O(n) - Linear Time:** The runtime grows linearly with the input size.  If you double the input size, you double the runtime.  Searching for an element in an unsorted array is O(n).
* **O(n log n) - Log-linear Time:**  A combination of linear and logarithmic growth.  Many efficient sorting algorithms (Merge Sort, Heap Sort, Quick Sort – on average) fall into this category.
* **O(n^2) - Quadratic Time:** The runtime grows proportionally to the square of the input size.  Nested loops often lead to O(n^2) complexity.  Simple sorting algorithms like Bubble Sort and Selection Sort are O(n^2).
* **O(2^n) - Exponential Time:** The runtime doubles with each addition to the input.  This is very slow and becomes impractical for even moderately sized inputs.
* **O(n!) - Factorial Time:**  The runtime grows factorially.  This is *extremely* slow and is rarely encountered in practical algorithms.

**Visualizing Growth Rates:**

Growth Rates of Common Big O Notations

```bash
|                       /
|                      / O(n!)
|                     /
|                    /  O(2^n)
|                   /
|                  /   O(n^2)
|                 /
|                /    O(n log n)
|               /
|              /     O(n)
|             /
|            /      O(log n)
|           /
|          /       O(1)
|_________/____________________
      Input Size (n)
```

**Why Big O Matters:**

* **Scalability:**  Big O helps us predict how an algorithm will perform as the data grows. An O(n^2) algorithm might work fine for 10 books but become unusable for 10,000.
* **Resource Optimization:**  Choosing efficient algorithms (lower Big O) saves time and processing power.
* **Trade-offs:**  Sometimes, we have to choose between algorithms with different Big O complexities based on the specific needs of the application (e.g., prioritizing insertion speed over search speed).

## Calculating Time and Space Complexity

* **Time Complexity:** Analyze the number of operations an algorithm performs as a function of the input size (n).
  * **Count Basic Operations:** Identify the fundamental operations (comparisons, assignments, arithmetic operations, etc.).
  * **Focus on Loops:** Loops (and nested loops) are often the primary drivers of time complexity.
  * **Recursive Functions:** Analyze the number of recursive calls and the work done in each call.

* **Space Complexity:** Analyze the amount of memory an algorithm uses as a function of the input size.
  * **Variables:** Consider the space used by variables, data structures, and function call stacks.
  * **Auxiliary Space:**  Focus on the *extra* space used by the algorithm, *in addition* to the space occupied by the input itself.

### Example: Linear Search

```go
func linearSearch(arr []int, target int) int {
    for i := 0; i < len(arr); i++ { // Loop iterates 'n' times (where n = len(arr))
        if arr[i] == target {      // Comparison (constant time)
            return i             // Return (constant time)
        }
    }
    return -1                  // Return (constant time)
}
```

Time Complexity: O(n) – The loop iterates, in the worst case, through all n elements of the array.

Space Complexity: O(1) – The algorithm uses a fixed amount of extra space (the variable i), regardless of the input size.

### Example: Nested Loops

```go
func nestedLoop(n int) {
    for i := 0; i < n; i++ {     // Outer loop: 'n' iterations
        for j := 0; j < n; j++ { // Inner loop: 'n' iterations for each outer iteration
            fmt.Println(i, j)   // Constant time operation
        }
    }
}
```

Time Complexity: O(n^2) – The inner loop executes n times for each iteration of the outer loop, resulting in n * n = n^2 operations.

Space Complexity: O(1)

## Key Takeaways

Big O notation is a crucial tool for understanding and comparing algorithm efficiency.

Focus on the dominant term and ignore constant factors.

Strive for the lowest possible Big O complexity (while considering trade-offs).

Analyze both time and space complexity.

This lesson provides a foundational understanding of algorithm analysis. We'll refer back to these concepts throughout our journey as we build the Alexandria. These concepts will be revisited by the characters throughout the story.
