package groph

import "testing"

type VertexMockData struct {
	Data string
}

func (v *VertexMockData) GetData() interface{} {
	return v.Data
}

func (v *VertexMockData) GetID() interface{} {
	return v.Data
}

func getMockedGraph() *Graph {
	graph := NewGraph()

	a, b := graph.NewVertex(&VertexMockData{"A"}), graph.NewVertex(&VertexMockData{"B"})

	start, finish := graph.NewVertex(&VertexMockData{"start"}), graph.NewVertex(&VertexMockData{"finish"})

	graph.StartVertex = start

	graph.AddConnection(start, a, &Edge{Data: "Start to A", Weight: 6})
	graph.AddConnection(start, b, &Edge{Data: "Start to B", Weight: 2})

	graph.AddConnection(a, finish, &Edge{Data: "A to Finish", Weight: 1})
	graph.AddConnection(b, a, &Edge{Data: "A to B", Weight: 3})

	graph.AddConnection(b, finish, &Edge{Data: "B to Finish", Weight: 5})
	graph.AddConnection(b, a, &Edge{Data: "B to A", Weight: 3})

	graph.AddConnection(finish, start, &Edge{Data: "Finish to start", Weight: 100})

	graph.SetRootVertex(start)

	return graph
}

func TestGraph_LoadAndSaveToDisk(t *testing.T) {
	graph := getMockedGraph()

	t.Run("Save to disk", func(t *testing.T) {
		err := graph.SaveToDisk("/tmp/graph")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Load from disk", func(t *testing.T) {
		graph, err := LoadGraphFromDisk("/tmp/graph")
		if err != nil {
			t.Fatal(err)
		}

		if graph.StartVertex.GetID() != "start" {
			t.Fail()
		}
	})
}
