package sorting

import (
	"reflect"
	"testing"
)

func TestSortingLessons(t *testing.T) {
	t.Run("Lesson08", func(t *testing.T) {
		t.Run("TestBubbleSort", func(t *testing.T) { testSort(t, BubbleSort, "BubbleSort") })
		t.Run("TestSelectionSort", func(t *testing.T) { testSort(t, SelectionSort, "SelectionSort") })
		t.Run("TestInsertionSort", func(t *testing.T) { testSort(t, InsertionSort, "InsertionSort") })
		t.Run("TestMergeSort", func(t *testing.T) { testSort(t, MergeSort, "MergeSort") })
		t.Run("TestQuickSort", func(t *testing.T) { testSort(t, QuickSort, "QuickSort") })
		t.Run("TestHeapSort", func(t *testing.T) { testSort(t, HeapSort, "HeapSort") })
		t.Run("TestSort_EdgeCases", TestSort_EdgeCases)
		t.Run("TestStability", TestSort_Stability)
	})
}

func testSort(t *testing.T, sortFunc func([]int), sortName string) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 8, 1, 9, 4, 7, 3, 6}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
	}

	for _, tc := range testCases {
		inputCopy := make([]int, len(tc.input))
		copy(inputCopy, tc.input)

		sortFunc(inputCopy)
		if !reflect.DeepEqual(inputCopy, tc.expected) {
			t.Errorf("%s: Expected %v, got %v (original: %v)", sortName, tc.expected, inputCopy, tc.input)
		}
	}
}

func benchmarkSort(b *testing.B, sortFunc func([]int)) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = len(data) - i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dataCopy := make([]int, len(data))
		copy(dataCopy, data)
		sortFunc(dataCopy)
	}
}

func BenchmarkBubbleSort(b *testing.B)    { benchmarkSort(b, BubbleSort) }
func BenchmarkSelectionSort(b *testing.B) { benchmarkSort(b, SelectionSort) }
func BenchmarkInsertionSort(b *testing.B) { benchmarkSort(b, InsertionSort) }
func BenchmarkMergeSort(b *testing.B)     { benchmarkSort(b, MergeSort) }
func BenchmarkQuickSort(b *testing.B)     { benchmarkSort(b, QuickSort) }
func BenchmarkHeapSort(b *testing.B)      { benchmarkSort(b, HeapSort) }

func TestSort_EdgeCases(t *testing.T) {
	sortFuncs := map[string]func([]int){
		"BubbleSort":    BubbleSort,
		"SelectionSort": SelectionSort,
		"InsertionSort": InsertionSort,
		"MergeSort":     MergeSort,
		"QuickSort":     QuickSort,
		"HeapSort":      HeapSort,
	}

	for name, sortFunc := range sortFuncs {
		t.Run(name, func(t *testing.T) {
			t.Run("Empty Slice", func(t *testing.T) {
				input := []int{}
				expected := []int{}
				sortFunc(input)
				if !reflect.DeepEqual(input, expected) {
					t.Errorf("Expected %v, got %v", expected, input)
				}
			})

			t.Run("Single Element Slice", func(t *testing.T) {
				input := []int{5}
				expected := []int{5}
				sortFunc(input)
				if !reflect.DeepEqual(input, expected) {
					t.Errorf("Expected %v, got %v", expected, input)
				}
			})
			t.Run("Already Sorted Slice", func(t *testing.T) {
				input := []int{1, 2, 3, 4, 5}
				expected := []int{1, 2, 3, 4, 5}
				sortFunc(input)

				if !reflect.DeepEqual(input, expected) {
					t.Errorf("Expected: %v, got: %v", expected, input)
				}
			})
			t.Run("Reverse Sorted Slice", func(t *testing.T) {
				input := []int{5, 4, 3, 2, 1}
				expected := []int{1, 2, 3, 4, 5}
				sortFunc(input)

				if !reflect.DeepEqual(input, expected) {
					t.Errorf("Expected: %v, got: %v", expected, input)
				}
			})
			t.Run("Slice with Duplicate Values", func(t *testing.T) {
				input := []int{5, 2, 8, 2, 9, 5}
				expected := []int{2, 2, 5, 5, 8, 9}
				sortFunc(input)

				if !reflect.DeepEqual(input, expected) {
					t.Errorf("Expected: %v, got: %v", expected, input)
				}
			})
		})
	}
}

// --- Stability Tests ---
type sortTest struct {
	value         int
	originalIndex int //To test stability
}

func TestSort_Stability(t *testing.T) {
	stableSorts := map[string]func([]sortTest, func(sortTest, sortTest) bool){
		"MergeSort":     GenericMergeSort, // Assuming you have a generic version
		"BubbleSort":    GenericBubbleSort,
		"InsertionSort": GenericInsertionSort,
	}

	unstableSorts := map[string]func([]sortTest, func(sortTest, sortTest) bool){
		"QuickSort":     GenericQuickSort, //Assuming you have a generic version
		"HeapSort":      GenericHeapSort,
		"SelectionSort": GenericSelectionSort,
	}

	less := func(a, b sortTest) bool {
		return a.value < b.value
	}

	// Test cases
	data := []sortTest{
		{5, 0},
		{2, 1},
		{8, 2},
		{2, 3}, // Duplicate value
		{9, 4},
		{5, 5}, // Duplicate Value
	}
	expectedStable := []sortTest{
		{2, 1}, {2, 3}, {5, 0}, {5, 5}, {8, 2}, {9, 4},
	}

	for name, sortFunc := range stableSorts {
		t.Run(name, func(t *testing.T) {
			dataCopy := make([]sortTest, len(data))
			copy(dataCopy, data)
			sortFunc(dataCopy, less)
			if !reflect.DeepEqual(dataCopy, expectedStable) {
				t.Errorf("Expected %v, got %v", expectedStable, dataCopy)
			}
		})
	}

	for name, sortFunc := range unstableSorts {
		t.Run(name, func(t *testing.T) {
			dataCopy := make([]sortTest, len(data))
			copy(dataCopy, data)
			sortFunc(dataCopy, less)
			if reflect.DeepEqual(dataCopy, expectedStable) {
				t.Errorf("%s is claimed to be unstable, but produced a stable result. This doesn't *guarantee* instability, but is suspicious.", name)
			}
		})
	}
}

// Generic versions of sorting algorithms (needed for stability test)
func GenericBubbleSort(data []sortTest, less func(a, b sortTest) bool) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if less(data[j+1], data[j]) {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func GenericSelectionSort(data []sortTest, less func(a, b sortTest) bool) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if less(data[j], data[minIdx]) {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
}

func GenericInsertionSort(data []sortTest, less func(a, b sortTest) bool) {
	n := len(data)
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1

		for j >= 0 && less(key, data[j]) {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}
}
func GenericMergeSort(data []sortTest, less func(a, b sortTest) bool) {
	if len(data) <= 1 { //Base case
		return
	}

	mid := len(data) / 2
	left := make([]sortTest, mid)
	right := make([]sortTest, len(data)-mid)

	copy(left, data[:mid])
	copy(right, data[mid:])

	GenericMergeSort(left, less)  //Sort left
	GenericMergeSort(right, less) //Sort right

	//Merge
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if less(left[i], right[j]) {
			data[k] = left[i]
			i++
		} else {
			data[k] = right[j]
			j++
		}
		k++
	}

	//Copy remaining elements
	for i < len(left) {
		data[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		data[k] = right[j]
		j++
		k++
	}
}
func GenericQuickSort(data []sortTest, less func(a, b sortTest) bool) {
	if len(data) < 2 {
		return
	}
	pivot := data[0]
	i := 1
	j := len(data) - 1
	for {
		for ; i <= j && less(data[i], pivot); i++ {
		}
		for ; i <= j && less(pivot, data[j]); j-- {
		}

		if i > j {
			break
		}
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
	data[0], data[j] = data[j], data[0]
	GenericQuickSort(data[:j], less)
	GenericQuickSort(data[j+1:], less)
}

func GenericHeapSort(data []sortTest, less func(a, b sortTest) bool) {
	//Not implemented yet, not needed for this project
}
