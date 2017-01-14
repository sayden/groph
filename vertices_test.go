package groph

import (
	"strings"
	"testing"
)

func TestVertices_Map(t *testing.T) {
	graph := getMockedGraph()

	amount := len(graph.StartVertex.outEdges)

	var vertices Vertices = make([]*Vertex, amount)
	for k, e := range graph.StartVertex.outEdges {
		vertices[k] = e.PointsTo
	}

	modifiedData := "Modified data"
	for _, v := range vertices.Map(func(e *Vertex) *Vertex {
		e.Data = &AnyData{
			Data: modifiedData,
			ID:   modifiedData,
		}

		return e
	}) {
		amount--
		if v.Data.GetID() != modifiedData {
			t.Fail()
		}
	}
}

func TestVertices_Each(t *testing.T) {
	graph := getMockedGraph()

	modifiedData := "Modified data"
	amount := len(graph.StartVertex.outEdges)

	var vertices Vertices = make([]*Vertex, amount)
	for k, e := range graph.StartVertex.outEdges {
		vertices[k] = e.PointsTo
	}

	vertices.Each(func(e *Vertex) {
		e.Data = &AnyData{
			Data: modifiedData,
			ID:   modifiedData,
		}
		amount--
	})

	for _, v := range graph.StartVertex.outEdges {
		if v.PointsTo.GetData() != modifiedData {
			t.Fatalf("%s != %s\n", v.Data.GetData(), modifiedData)
		}
	}

	if amount != 0 {
		t.Fatalf("%d != 1\n", amount)
	}
}

func TestVertices_Filter(t *testing.T) {
	graph := getMockedGraph()

	amount := len(graph.StartVertex.outEdges)
	var vertices Vertices = make([]*Vertex, amount)
	for k, e := range graph.StartVertex.outEdges {
		vertices[k] = e.PointsTo
	}

	if len(vertices.Filter(func(e *Vertex) bool {
		return e.GetID() == "nonExistent ID"
	})) != 0 {
		t.Fail()
	}

	results := len(vertices.Filter(func(e *Vertex) bool {
		s := e.GetID().(string)
		return strings.Contains(s, "A")
	}))

	if results != 1 {
		t.Fatalf("%d != 1\n", results)
	}
}

func TestVertices_Fold(t *testing.T) {
	graph := getMockedGraph()

	amount := len(graph.StartVertex.outEdges)
	var vertices Vertices = make([]*Vertex, amount)
	for k, e := range graph.StartVertex.outEdges {
		vertices[k] = e.PointsTo
		vertices[k].SetData(5)
		vertices[k].SetID(5)
	}

	res := vertices.Fold(0, func(a interface{}, b *Vertex) interface{} {
		f1 := a.(int)
		f2 := b.GetData().(int)

		return f1 + f2
	})

	if res != 10 {
		t.Fatalf("%d != %d\n", res, 10)
	}
}

func TestVertices_MapT(t *testing.T) {
	graph := getMockedGraph()

	vertices := graph.StartVertex.OutEdges().MapV(func(e *Edge) *Vertex {
		return e.PointsTo
	}).MapT(func(e *Vertex) Data {
		return e.inEdges[0]
	}).Edges()

	if vertices == nil {
		t.Fatal("vertices are nil")
	}

	if vertices.Size() != 2 {
		t.Errorf("%d != 2\n", vertices.Size())
	}
}

func TestVertices_MapT2(t *testing.T) {
	graph := getMockedGraph()

	vertices := graph.StartVertex.OutEdges().MapV(func(e *Edge) *Vertex {
		return e.PointsTo
	}).MapT(func(e *Vertex) Data {
		return e.inEdges[0]
	}).Edges()

	if vertices == nil {
		t.Fatal("vertices are nil")
	}

	if vertices.Size() != 2 {
		t.Errorf("%d != 2\n", vertices.Size())
	}
}

func TestVertices_MapE(t *testing.T) {
	graph := getMockedGraph()

	edges := graph.StartVertex.OutEdges().MapV(func(e *Edge) *Vertex {
		return e.PointsTo
	}).MapE(func(v *Vertex) *Edge {
		return v.inEdges[0]
	})

	if edges.Size() == 0 {
		t.Fail()
	}

	if edges[0] == nil {
		t.Fail()
	}
}


func TestVertices_FlatMap(t *testing.T) {
	graph := getMockedGraph()

	vertices := graph.StartVertex.OutEdges().MapV(func(e *Edge) *Vertex {
		return e.PointsTo
	}).FlatMap(func(v *Vertex)Vertices{
		return v.InEdges().PointsToVertices().OutEdges().FromVertices().InEdges().FromVertices()
	})

	if vertices.Size() == 0 {
		t.Fail()
	}

	if vertices[0] == nil {
		t.Fail()
	}
}
