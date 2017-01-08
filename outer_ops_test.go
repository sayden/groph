package groph

import "testing"

func TestVertex_Outer(t *testing.T) {
	graph := getMockedGraph()

	edges, err := graph.Outer(graph.StartVertex)
	if err != nil {
		t.Fatal()
	}

	for _, e := range edges {
		if e.PointsTo.GetID() != "A" && e.PointsTo.GetID() != "B" {
			t.Fail()
		}
	}
}

func TestVertex_OuterWhereEdge(t *testing.T) {
	graph := getMockedGraph()

	edges, err := graph.OuterWhereEdge(func(e *Edge) bool {
		return e.PointsTo.GetID() == "A"
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(edges) != 1 {
		t.Fatal()
	}

	if edges[0].PointsTo.GetID() != "A" {
		t.Fatal()
	}
}

func TestGraph_OuterWhereVertex(t *testing.T) {
	graph := getMockedGraph()

	edges ,err := graph.OuterWhereVertex(func(v *Vertex) bool {
		return v.GetID() == "A"
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(edges) != 1 {
		t.Fail()
	}

	for _, edge := range edges {
		if edge.PointsTo.GetID() != "finish" {
			t.Fail()
		}
	}
}