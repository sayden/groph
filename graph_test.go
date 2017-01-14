package groph

import (
	"testing"
	"fmt"
)

func getMockedGraph() *Graph {
	graph := New()

	a := graph.NewVertex(&AnyData{Data: "A", ID: "A"})
	b := graph.NewVertex(&AnyData{Data: "B", ID: "B"})

	start := graph.NewVertex(&AnyData{Data: "start", ID: "start"})

	finish := graph.NewVertexWithUpdate(&AnyData{Data: "finish", ID: "finish"})
	finish = graph.NewVertexWithUpdate(&AnyData{Data: "finish", ID: "finish"})

	graph.AddConnection(start, a, graph.NewEdge(&AnyData{Data: "Start to A", ID: "Start to A"}, 6))
	graph.AddConnection(start, b, graph.NewEdge(&AnyData{Data: "Start to B", ID: "Start to B"}, 2))

	graph.AddConnection(a, finish, graph.NewEdge(&AnyData{Data: "A to Finish", ID: "A to Finish"}, 1))
	graph.AddConnection(a, b, graph.NewEdge(&AnyData{Data: "A to B", ID: "A to B"}, 3))

	graph.AddConnection(b, finish, graph.NewEdge(&AnyData{Data: "B to Finish", ID: "B to Finish"}, 5))
	graph.AddConnection(b, a, graph.NewEdge(&AnyData{Data: "B to A", ID: "B to A"}, 3))

	graph.AddConnection(finish, start, graph.NewEdge(&AnyData{Data: "Finish to start", ID: "Finish to start"}, 100))

	graph.SetRootVertex(start)

	return graph
}

func TestGraph_Find(t *testing.T) {
	graph := getMockedGraph()

	_, err := graph.Find("start")
	if err != nil {
		t.Fatal(err)
	}

	notFound, err := graph.Find("not exists")
	if err == nil {
		t.Fatal("Expected error, found nil")
	}

	if notFound != nil {
		t.Fatalf("Expected nil vertex, found %#v\n", notFound)
	}
}

func TestVertex_Out(t *testing.T) {
	graph := getMockedGraph()

	edges := graph.Out(graph.StartVertex)

	for _, e := range edges {
		if e.PointsTo.GetID() != "A" && e.PointsTo.GetID() != "B" {
			t.Fail()
		}
	}
}

func TestVertex_OutWhereEdge(t *testing.T) {
	graph := getMockedGraph()

	edges := graph.OutWhereEdge(func(e *Edge) bool {
		return e.PointsTo.GetID() == "A"
	})

	if len(edges) != 1 {
		t.Fatal()
	}

	if edges[0].PointsTo.GetID() != "A" {
		t.Fatal()
	}
}

func TestGraph_OutWhereVertex(t *testing.T) {
	graph := getMockedGraph()

	edges := graph.OutWhereVertex(func(v *Vertex) bool {
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

func TestVertex_In(t *testing.T) {
	graph := getMockedGraph()

	edges := graph.In(graph.StartVertex)

	if len(edges) != 1 {
		t.Errorf("Unexpected number of edges %d != %d\n", len(edges), 1)
	}

	for _, e := range edges {
		if e.PointsTo.GetID() != "start" {
			t.Errorf("%s != %s", e.PointsTo.GetID(), "start")
		}
	}
}

func TestGraph_InWhereEdge(t *testing.T) {
	graph := getMockedGraph()

	edges := graph.InWhereEdge(func(e *Edge) bool {
		return e.PointsTo.GetID() == "start"
	})

	if edges[0].From.GetID() != "finish" || edges[0].PointsTo.GetID() != "start" {
		fmt.Printf("%#v\n", edges[0])
		t.Fail()
	}
}

func TestGraph_InWhereVertex(t *testing.T) {
	graph := getMockedGraph()

	edges := graph.InWhereVertex(func(v *Vertex) bool {
		return v.GetID() == "finish"
	})

	if len(edges) != 2 {
		t.Errorf("Unexpected number of edges %d != %d\n", len(edges), 2)
	}

	for _, edge := range edges {
		if edge.From.GetID() != "A" && edge.From.GetID() != "B" {
			t.Fail()
		}
	}
}