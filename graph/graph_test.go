package graph

import (
	"reflect"
	"sort"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph()

	vertices := []string{"A", "B", "C"}
	for _, v := range vertices {
		g.AddVertex(v)
	}

	g.AddEdge("A", "B")
	g.AddEdge("A", "C")

	got := g.GetNeighbors("A")
	sort.Strings(got)
	want := []string{"B", "C"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNeighbors(A) = %v; want %v", got, want)
	}

	gotVertices := g.GetVertices()
	sort.Strings(gotVertices)

	if !reflect.DeepEqual(gotVertices, vertices) {
		t.Errorf("GetVertices() = %v; want %v", gotVertices, vertices)
	}
}
