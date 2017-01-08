package groph

import (
	"fmt"
	"testing"
)

func TestBreadthFirst(t *testing.T) {
	graph := getMockedGraph()

	var found bool
	foundFunc := func(v *Vertex) {
		fmt.Printf("Found %#v\n", v)
		found = true
	}

	graph.BreadthFirst("finish", foundFunc)
	if !found {
		t.Fail()
	}

	found = false
	graph.BreadthFirst("not-found", foundFunc)
	if found {
		t.Fail()
	}
}
