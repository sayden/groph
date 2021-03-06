package groph

import (
	"strings"
	"testing"
)

func TestEdges_Map(t *testing.T) {
	graph := getMockedGraph()

	modifiedData := "Modified data"
	amount := len(graph.StartVertex.outEdges)
	for _, v := range graph.Out(graph.StartVertex).Map(func(e *Edge) *Edge {
		e.Data = &AnyData{
			Data: modifiedData,
		}

		return e
	}) {
		amount--
		if v.Data.GetData() != modifiedData {
			t.Fail()
		}
	}

	if amount != 0 {
		t.Fatalf("%d != 1\n", amount)
	}
}

func TestEdges_Each(t *testing.T) {
	graph := getMockedGraph()

	modifiedData := "Modified data"
	amount := len(graph.StartVertex.outEdges)
	graph.Out(graph.StartVertex).Each(func(e *Edge) {
		e.Data = &AnyData{
			Data: modifiedData,
		}
		amount--
	})

	for _, v := range graph.StartVertex.outEdges {
		if v.Data.GetData() != modifiedData {
			t.Fail()
		}
	}

	if amount != 0 {
		t.Fatalf("%d != 1\n", amount)
	}
}

func TestEdges_Filter(t *testing.T) {
	graph := getMockedGraph()

	if len(graph.Out(graph.StartVertex).Filter(func(e *Edge) bool {
		return e.GetID() == "nonExistent ID"
	})) != 0 {
		t.Fail()
	}

	results := len(graph.Out(graph.StartVertex).Filter(func(e *Edge) bool {
		s := e.GetID().(string)
		return strings.Contains(s, " to ")
	}))

	if results != 2 {
		t.Fatalf("%d != 2\n", results)
	}
}

func TestEdges_Fold(t *testing.T) {
	graph := getMockedGraph()

	res := graph.Out(graph.StartVertex).Fold(float64(0),func(a interface{}, b *Edge) interface{}{
		f := a.(float64)
		return b.Weight+f
	})

	if res != float64(8) {
		t.Fatalf("%f != 8\n", res)
	}
}

func TestEdges_MapV(t *testing.T) {
	graph := getMockedGraph()

	vertices := graph.StartVertex.InEdges().MapV(func(e *Edge) *Vertex{
		return e.PointsTo
	})

	if vertices == nil {
		t.Fatal("vertices are nil")
	}

	if vertices.Size() != 1 {
		t.Errorf("%d != 2\n", vertices.Size())
	}
}

func TestEdges_MapT(t *testing.T) {
	graph := getMockedGraph()

	vertices := graph.StartVertex.InEdges().MapT(func(e *Edge) Data {
		return e.PointsTo
	}).Vertices()

	if vertices == nil {
		t.Fatal("vertices are nil")
	}

	if vertices.Size() != 1 {
		t.Errorf("%d != 2\n", vertices.Size())
	}
}


func TestEdges_MapT2(t *testing.T) {
	graph := getMockedGraph()

	edges := graph.StartVertex.InEdges().MapT(func(e *Edge) Data {
		return e.PointsTo.InEdges()[0]
	}).Edges()

	if edges == nil {
		t.Fatal("edges are nil")
	}

	if edges.Size() != 1 {
		t.Errorf("%d != 2\n", edges.Size())
	}
}