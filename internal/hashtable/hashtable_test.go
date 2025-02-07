package hashtable

import (
	"fmt"
	"testing"
)

func TestHashTableLessons(t *testing.T) {
	t.Run("Lesson06", func(t *testing.T) {
		t.Run("TestHashTable", TestHashTable)
		t.Run("TestHashTable_EdgeCases", TestHashTable_EdgeCases)
		t.Run("TestHashTable_Resizing", TestHashTable_Resizing)
	})
	t.Run("Lesson13", func(t *testing.T) {
		t.Run("TestHashTable_Iterator", TestHashTable_Iterator)
	})
}

func TestHashTable(t *testing.T) {
	ht := NewHashTable(10)
	ht.Put("John Doe", 123)
	ht.Put("Jane Smith", 456)
	ht.Put("Peter Jones", 789)
	if !ht.Contains("John Doe") {
		t.Errorf("Expected to contain 'John Doe'")
	}
	if ht.Contains("Nonexistent") {
		t.Errorf("Expected not to contain 'Nonexistent'")
	}
	val, ok := ht.Get("Jane Smith")
	if !ok || val != 456 {
		t.Errorf("Expected Get('Jane Smith') to return (456, true), got (%v, %v)", val, ok)
	}
	_, ok = ht.Get("Nonexistent")
	if ok {
		t.Errorf("Expected Get('Nonexistent') to return false")
	}
	ht.Delete("John Doe")
	if ht.Contains("John Doe") {
		t.Errorf("Expected 'John Doe' to be deleted")
	}
	if ht.Size() != 2 {
		t.Errorf("Expected size 2, got %d", ht.Size())
	}
}

func TestHashTable_EdgeCases(t *testing.T) {
	t.Run("Empty HashTable", func(t *testing.T) {
		ht := NewHashTable(10)

		if ht.Contains("anything") {
			t.Errorf("Contains on empty table should return false")
		}

		_, ok := ht.Get("anything")
		if ok {
			t.Errorf("Get on empty table should return false")
		}

		ht.Delete("anything") // Delete on an empty table
	})

	t.Run("Deleting Non-Existent Key", func(t *testing.T) {
		ht := NewHashTable(10)
		ht.Put("key1", 1)

		ht.Delete("nonexistent")

		if ht.Size() != 1 { //should not have changed
			t.Errorf("Size changed after deleting non-existent key. Got: %d, Expected: 1", ht.Size())
		}
		if !ht.Contains("key1") {
			t.Error("Key1 should still be in the table")
		}
	})
}

func TestHashTable_Resizing(t *testing.T) {
	ht := NewHashTable(2) // Start with a very small capacity
	for i := 0; i < 10; i++ {
		ht.Put(fmt.Sprintf("key%d", i), i)
	}

	if ht.capacity <= 2 {
		t.Errorf("Expected capacity to be greater than 2 after resizing, got %d", ht.capacity)
	}
	if ht.Size() != 10 {
		t.Errorf("Expected size 10, got %d", ht.Size())
	}

	for i := 0; i < 10; i++ {
		val, ok := ht.Get(fmt.Sprintf("key%d", i))
		if !ok || val != i {
			t.Errorf("Expected Get('key%d') to return (%d, true), got (%v, %v)", i, i, val, ok)
		}
	}
}

func TestHashTable_Iterator(t *testing.T) {
	ht := NewHashTable(10)
	ht.Put("Alice", 1)
	ht.Put("Bob", 2)
	ht.Put("Charlie", 3)

	// Create a map to store the expected key-value pairs
	expected := map[string]interface{}{
		"Alice":   1,
		"Bob":     2,
		"Charlie": 3,
	}

	count := 0
	// Iterate over the hash table using the Iterator
	for kv := range ht.Iterator() {
		count++
		// Check if the key exists in the expected map
		expectedValue, ok := expected[kv.Key]
		if !ok {
			t.Errorf("Unexpected key found: %s", kv.Key)
		}

		// Check if the value matches the expected value
		if kv.Value != expectedValue {
			t.Errorf("Expected value %v for key %s, got %v", expectedValue, kv.Key, kv.Value)
		}

		// Remove the key from the expected map to ensure we don't visit it again
		delete(expected, kv.Key)
	}
	if count != ht.Size() {
		t.Errorf("Expected iterator to have size of %d, got %d", ht.Size(), count)
	}

	// If the expected map is not empty, it means we missed some keys
	if len(expected) != 0 {
		t.Errorf("Missing keys in iteration: %v", expected)
	}
}
