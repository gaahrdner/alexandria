# Sorting Algorithms

This package provides implementations of various sorting algorithms in Go.

## Implemented Algorithms

### Bubble Sort

* Simple and educational but inefficient.
* **Time Complexity:** O(n^2)
* **Space Complexity:** O(1) (in-place)
* **Stability:** Stable

### Selection Sort

* Finds the minimum repeatedly and swaps it.
* **Time Complexity:** O(n^2)
* **Space Complexity:** O(1) (in-place)
* **Stability:** Not stable

### Insertion Sort

* Inserts each element into its correct position in the sorted portion.
* **Time Complexity:** O(n^2), O(n) if nearly sorted
* **Space Complexity:** O(1) (in-place)
* **Stability:** Stable

### Merge Sort

* Divide-and-conquer algorithm that splits, sorts, and merges.
* **Time Complexity:** O(n log n)
* **Space Complexity:** O(n)
* **Stability:** Stable

### Quick Sort

* Divide-and-conquer using a pivot to partition.
* **Time Complexity:** O(n log n) average, O(n^2) worst
* **Space Complexity:** O(log n) average
* **Stability:** Not stable

### Heap Sort

* Uses a binary heap structure.
* **Time Complexity:** O(n log n)
* **Space Complexity:** O(1) (in-place)
* **Stability:** Not stable

## Usage

```go
package main

import (
    "fmt"
    "github.com/gaahrdner/library/internal/sorting"
)

func main() {
    arr := []int{5, 2, 8, 1, 9, 4}
    sorting.QuickSort(arr) // Or any other sorting function
    fmt.Println(arr)
}
```

## Testing

Run tests with:

```bash
go test ./...
```

Run benchmarks with:

```bash
go test -bench=. -benchmem
```
