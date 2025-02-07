package bst

import (
	"reflect"
	"testing"
)

func TestBSTLessons(t *testing.T) {
	t.Run("Lesson07", func(t *testing.T) {
		t.Run("TestBST", TestBST)
		t.Run("TestBST_EdgeCases", TestBST_EdgeCases)
        t.Run("TestBST_Deletion_Comprehensive", TestBST_Deletion_Comprehensive)
	})
}


func TestBST(t *testing.T) {
	tree := NewBST()
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(70)
	tree.Insert(60)
	tree.Insert(80)
	result := tree.InOrderTraversal()
	expected := []int{20, 30, 40, 50, 60, 70, 80}
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Expected InOrderTraversal to have %v at %d, got %v", expected[i], i, val)
		}
	}
	if !tree.Search(40) {
		t.Errorf("Expected to find 40")
	}
	if tree.Search(99) {
		t.Errorf("Expected not to find 99")
	}
	if tree.Min() != 20 {
		t.Errorf("Expected min to be 20, got %d", tree.Min())
	}
	if tree.Max() != 80 {
		t.Errorf("Expected max to be 80, got %d", tree.Max())
	}
	tree.Delete(20)
	if tree.Search(20) {
		t.Errorf("Expected 20 to be deleted")
	}
	tree.Delete(30)
	if tree.Search(30) {
		t.Errorf("Expected 30 to be deleted")
	}
	tree.Delete(50)
	if tree.Search(50) || tree.Root.Key != 60 {
		t.Errorf("Expected 50 to be deleted and root to be 60")
	}
}

func TestBST_EdgeCases(t *testing.T) {
	t.Run("Empty Tree", func(t *testing.T) {
		tree := NewBST()
		if tree.Search(10) {
			t.Errorf("Search in empty tree should return false")
		}
		if tree.Min() != 0 { // Assuming 0 is returned for an empty tree's Min/Max
			t.Errorf("Min of empty tree should return 0")
		}
		if tree.Max() != 0 {
			t.Errorf("Max of empty tree should return 0")
		}

		//Test in order traversal on empty tree
		result := tree.InOrderTraversal()
		if len(result) != 0 {
			t.Errorf("Expected empty slice for InOrderTraversal, got: %v", result)
		}

		tree.Delete(5) // Deleting from an empty tree should not panic
	})

	t.Run("Single Node Tree", func(t *testing.T) {
		tree := NewBST()
		tree.Insert(1)

		if !tree.Search(1) {
			t.Errorf("Should be able to find key in single node tree.")
		}

		//Test Min
		if tree.Min() != 1 {
			t.Errorf("Expected min to be 1, got %d", tree.Min())
		}
		//Test Max
		if tree.Max() != 1 {
			t.Errorf("Expected max to be 1, got %d", tree.Max())
		}

		//Test InorderTraversal
		result := tree.InOrderTraversal()
		expected := []int{1}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		//Test Deletion
		tree.Delete(1)
		if tree.Search(1) {
			t.Error("Deleted only element, tree should now be empty")
		}
	})

	t.Run("Deleting Non-Existent Node", func(t *testing.T) {
		tree := NewBST()
		tree.Insert(50)
		tree.Insert(30)
		tree.Insert(70)

		tree.Delete(99) // Should not cause any errors or panic
	})

	t.Run("Inserting Duplicates", func(t *testing.T) {
		tree := NewBST()
		tree.Insert(5)
		tree.Insert(5)

		result := tree.InOrderTraversal()
		expected := []int{5, 5} //Allow duplicates for now

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestBST_Deletion_Comprehensive(t *testing.T) {
    // Test deleting a leaf node
    t.Run("Delete Leaf", func(t *testing.T) {
        tree := NewBST()
        tree.Insert(5, "root")
        tree.Insert(3, "left")
        tree.Insert(7, "right")
        tree.Delete(3)
        if tree.Search(3) {
            t.Error("Failed to delete leaf node (left child)")
        }
        tree.Delete(7)
        if tree.Search(7){
            t.Error("Failed to delete leaf node (right child)")
        }
    })

    // Test deleting a node with one child
    t.Run("Delete Node with One Child", func(t *testing.T) {
        tree := NewBST()
        tree.Insert(5, "root")
        tree.Insert(3, "left")
        tree.Insert(2, "left-left")
        tree.Delete(3) // Node 3 has only a left child

        if tree.Search(3) {
            t.Error("Failed to delete node with one left child")
        }
        if !tree.Search(2){
            t.Error("Left child should be promoted")
        }
        tree = NewBST()
        tree.Insert(5, "root")
        tree.Insert(7, "left")
        tree.Insert(8, "left-left")
        tree.Delete(7)
        if tree.Search(7){
            t.Error("Failed to delete node with one right child")
        }
        if !tree.Search(8){
            t.Error("Right child should be promoted")
        }