package graph

import (
	"reflect"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph()

	// Add vertices
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")

	// Add edges
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "E")

	//Test GetNeighbors.
	neighbors := g.GetNeighbors("A")
	expectedNeighbors := []string{"B", "C"}

	if !reflect.DeepEqual(neighbors, expectedNeighbors) {
		t.Errorf("Expected neighbors of A to be %v, got %v", expectedNeighbors, neighbors)
	}

	// Test BFS
	bfsResult := g.BFS("A")
	expectedBFS := []string{"A", "B", "C", "D", "E"}
	if !reflect.DeepEqual(bfsResult, expectedBFS) {
		t.Errorf("BFS: Expected %v, got %v", expectedBFS, bfsResult)
	}

	// Test DFS
	dfsResult := g.DFS("A")
	expectedDFS := []string{"A", "B", "D", "C", "E"} // One possible DFS result
	if !reflect.DeepEqual(dfsResult, expectedDFS) {
		t.Errorf("DFS: Expected %v, got %v", expectedDFS, dfsResult)
	}
}
