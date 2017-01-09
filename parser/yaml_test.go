package parser

import "testing"

func TestParseSimpleYAMLFile(t *testing.T) {
	graph, err := ReadYAMLSimpleFileFormat("parser/simple.yaml")
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

func TestParseYAMLFile(t *testing.T) {
	graph, err := ReadYAMLExtendedFileFormat("parser/test.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if graph.StartVertex.GetID() != "start" {
		t.Fatalf("%s != start\n%s\n", graph.StartVertex.GetID(), graph.StartVertex)
	}

	if graph.StartVertex.OuterEdges[0].PointsTo.GetID() != "A" {
		t.Fatalf("%s != A\n", graph.StartVertex.OuterEdges[0].PointsTo.GetID())
	}

	v, err := graph.Find("B")
	if err != nil {
		t.Fatal(err)
	}

	if v.OuterEdges[0].PointsTo.GetID() != "A" {
		t.Fatalf("%s != A\n", v.OuterEdges[0].PointsTo.GetID())
	}
}
