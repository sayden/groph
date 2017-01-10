package groph

import (
	"strings"
	"testing"
)

func TestEdges_Map(t *testing.T) {
	graph := getMockedGraph()

	modifiedData := "Modified data"
	amount := len(graph.StartVertex.OuterEdges)
	for _, v := range graph.Outer(graph.StartVertex).Map(func(e *Edge) *Edge {
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
	amount := len(graph.StartVertex.OuterEdges)
	graph.Outer(graph.StartVertex).Each(func(e *Edge) {
		e.Data = &AnyData{
			Data: modifiedData,
		}
		amount--
	})

	for _, v := range graph.StartVertex.OuterEdges {
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

	if len(graph.Outer(graph.StartVertex).Filter(func(e *Edge) bool {
		return e.GetID() == "nonExistent ID"
	})) != 0 {
		t.Fail()
	}

	results := len(graph.Outer(graph.StartVertex).Filter(func(e *Edge) bool {
		s := e.GetID().(string)
		return strings.Contains(s, " to ")
	}))

	if results != 2 {
		t.Fatalf("%d != 2\n", results)
	}
}

func TestEdges_Fold(t *testing.T) {
	graph := getMockedGraph()

	res := graph.Outer(graph.StartVertex).Fold(float64(0),func(a interface{}, b *Edge) interface{}{
		f := a.(float64)
		return b.Weight+f
	})

	if res != float64(8) {
		t.Fatalf("%f != 8\n", res)
	}
}