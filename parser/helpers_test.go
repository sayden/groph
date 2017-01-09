package parser

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseExtendedFormatBytes(t *testing.T) {
	file, err := os.Open("parser/test.json")
	if err != nil {
		t.Fatal(err)
	}

	byt, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	graph, err := ParseExtendedFormatBytes(byt)
	if err != nil {
		t.Fatal(err)
	}

	if graph.StartVertex.GetID() != "start" {
		t.Fatalf("%s != A\n%s\n", graph.StartVertex.GetID(), graph.StartVertex)
	}

	if graph.StartVertex.OuterEdges[0].PointsTo.GetID() != "A" {
		t.Fatalf("%s != A\n%s\n", graph.StartVertex.OuterEdges[0].PointsTo.GetID(), graph.StartVertex)
	}

	v, err := graph.Find("B")
	if err != nil {
		t.Fatal(err)
	}

	if v.OuterEdges[0].PointsTo.GetID() != "A" {
		t.Fatalf("%s != A\n%s\n", v.OuterEdges[0].PointsTo.GetID(), graph.StartVertex)
	}
}

func TestParseSimpleFormatBytes(t *testing.T) {
	file, err := os.Open("parser/simple.json")
	if err != nil {
		t.Fatal(err)
	}

	byt, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	graph, err := ParseSimpleFormatBytes(byt)
	if err != nil {
		t.Fatal(err)
	}

	if graph.StartVertex.GetID() != "start" {
		t.Fatal()
	}

	if graph.StartVertex.OuterEdges[0].PointsTo.GetID() != "B" {
		t.Fatalf("%s != B", graph.StartVertex.OuterEdges[0].PointsTo.GetID())
	}

	v, err := graph.Find("B")
	if err != nil {
		t.Fatal(err)
	}

	if v.OuterEdges[1].PointsTo.GetID() != "A" {
		t.Fatalf("%s != A\n", v.OuterEdges[0].PointsTo.GetID())
	}
}
