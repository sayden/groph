package groph

import (
	"fmt"
	"testing"
)

func TestVertex_Inner(t *testing.T) {
	graph := getMockedGraph()

	edges, err := graph.Inner(graph.StartVertex)
	if err != nil {
		t.Fatal(err)
	}

	if len(edges) != 1 {
		t.Errorf("Unexpected number of edges %d != %d\n", len(edges), 1)
	}

	for _, e := range edges {
		if e.PointsTo.GetID() != "start" {
			t.Errorf("%s != %s", e.PointsTo.GetID(), "start")
		}
	}
}

func TestGraph_InnerWhereEdge(t *testing.T) {
	graph := getMockedGraph()

	edges, err := graph.InnerWhereEdge(func(e *Edge) bool {
		return e.PointsTo.GetID() == "start"
	})

	if err != nil {
		t.Fatal(err)
	}

	if edges[0].From.GetID() != "finish" || edges[0].PointsTo.GetID() != "start" {
		fmt.Printf("%#v\n", edges[0])
		t.Fail()
	}
}

func TestGraph_InnerWhereVertex(t *testing.T) {
	graph := getMockedGraph()

	edges, err := graph.InnerWhereVertex(func(v *Vertex) bool {
		return v.GetID() == "finish"
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(edges) != 2 {
		t.Errorf("Unexpected number of edges %d != %d\n", len(edges), 2)
	}

	for _, edge := range edges {
		if edge.From.GetID() != "A" && edge.From.GetID() != "B" {
			t.Fail()
		}
	}
}