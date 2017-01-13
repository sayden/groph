package groph

import "testing"

func TestVertex_Outer(t *testing.T) {
	graph := getMockedGraph()

	edges := graph.Outer(graph.StartVertex)

	for _, e := range edges {
		if e.PointsTo.GetID() != "A" && e.PointsTo.GetID() != "B" {
			t.Fail()
		}
	}
}

func TestVertex_OuterWhereEdge(t *testing.T) {
	graph := getMockedGraph()

	edges := graph.OuterWhereEdge(func(e *Edge) bool {
		return e.PointsTo.GetID() == "A"
	})

	if len(edges) != 1 {
		t.Fatal()
	}

	if edges[0].PointsTo.GetID() != "A" {
		t.Fatal()
	}
}

func TestGraph_OuterWhereVertex(t *testing.T) {
	graph := getMockedGraph()

	edges := graph.OuterWhereVertex(func(v *Vertex) bool {
		return v.GetID() == "A"
	})

	if len(edges) != 2 {
		t.Errorf("%d != 1\n", len(edges))
	}

	for _, edge := range edges {
		if edge.PointsTo.GetID() != "finish" && edge.PointsTo.GetID() != "B"{
			t.Fail()
		}
	}
}