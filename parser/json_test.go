package parser

import (
	"testing"
	"github.com/sayden/groph"
	"fmt"
)

func TestParseSimpleJSONFile(t *testing.T) {
	graph, err := ReadJSONSimpleFileFormat("parser/simple.json")
	if err != nil {
		t.Fatal(err)
	}

	if graph.StartVertex.GetID() != "start" {
		t.Fatal()
	}

	if graph.StartVertex.OutEdges()[0].PointsTo.GetID() != "B" {
		t.Fatalf("%s != B", graph.StartVertex.OutEdges()[0].PointsTo.GetID())
	}

	v, err := graph.Find("B")
	if err != nil {
		t.Fatal(err)
	}

	if v.OutEdges()[1].PointsTo.GetID() != "A" {
		t.Fatalf("%s != A\n", v.OutEdges()[0].PointsTo.GetID())
	}
}

func TestParseJSONFile(t *testing.T) {
	graph, err := ReadJSONExtendedFileFormat("parser/test.json")
	if err != nil {
		t.Fatal(err)
	}

	if graph.StartVertex.GetID() != "start" {
		t.Fatalf("%s != A\n%s\n", graph.StartVertex.GetID(), graph.StartVertex)
	}

	if graph.StartVertex.OutEdges()[0].PointsTo.GetID() != "A" {
		t.Fatalf("%s != A\n%s\n", graph.StartVertex.OutEdges()[0].PointsTo.GetID(), graph.StartVertex)
	}

	v, err := graph.Find("B")
	if err != nil {
		t.Fatal(err)
	}

	if v.OutEdges()[0].PointsTo.GetID() != "A" {
		t.Fatalf("%s != A\n%s\n", v.OutEdges()[0].PointsTo.GetID(), graph.StartVertex)
	}
}

func TestReadJSONSimpleFileFormat(t *testing.T) {
	graph, err := ReadJSONSimpleFileFormat("simple_big_graph.json")
	if err != nil {
		t.Fatal(err)
	}

	v, err := graph.Find("240")
	if err != nil {
		t.Fatal(err)
	}

	v.InEdges().Each(func(e *groph.Edge){
		fmt.Println(e)
	})
}