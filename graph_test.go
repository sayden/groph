package groph

import "testing"

func getMockedGraph() *Graph {
	graph := NewGraph()

	a := graph.NewVertex(&AnyData{Data: "A", ID: "A"})
	b := graph.NewVertex(&AnyData{Data: "B", ID: "B"})

	start := graph.NewVertex(&AnyData{Data: "start", ID: "start"})
	finish := graph.NewVertexWithUpdate(&AnyData{Data: "finish", ID: "finish"})
	finish = graph.NewVertexWithUpdate(&AnyData{Data: "finish", ID: "finish"})

	graph.AddConnection(start, a, graph.NewEdge(&AnyData{Data: "Start to A", ID: "Start to A"}, 6))
	graph.AddConnection(start, b, graph.NewEdge(&AnyData{Data: "Start to B", ID: "Start to B"}, 2))

	graph.AddConnection(a, finish, graph.NewEdge(&AnyData{Data: "A to Finish", ID: "A to Finish"}, 1))
	graph.AddConnection(b, a, graph.NewEdge(&AnyData{Data: "A to B", ID: "A to B"}, 3))

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