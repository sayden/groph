package groph

import (
	"strings"
	"testing"
)

func TestVertices_Map(t *testing.T) {
	graph := getMockedGraph()

	amount := len(graph.StartVertex.OuterEdges)

	var vertices Vertices = make([]*Vertex, amount)
	for k, e := range graph.StartVertex.OuterEdges {
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
	amount := len(graph.StartVertex.OuterEdges)

	var vertices Vertices = make([]*Vertex, amount)
	for k, e := range graph.StartVertex.OuterEdges {
		vertices[k] = e.PointsTo
	}

	vertices.Each(func(e *Vertex) {
		e.Data = &AnyData{
			Data: modifiedData,
			ID:   modifiedData,
		}
		amount--
	})

	for _, v := range graph.StartVertex.OuterEdges {
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

	amount := len(graph.StartVertex.OuterEdges)
	var vertices Vertices = make([]*Vertex, amount)
	for k, e := range graph.StartVertex.OuterEdges {
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

	amount := len(graph.StartVertex.OuterEdges)
	var vertices Vertices = make([]*Vertex, amount)
	for k, e := range graph.StartVertex.OuterEdges {
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
