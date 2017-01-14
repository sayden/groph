package groph

import "testing"

func TestEdge_Points(t *testing.T) {
	graph := getMockedGraph()

	v, err := graph.Find("A")
	if err != nil {
		t.Fatal(err)
	}

	v.String()

	if !graph.StartVertex.outEdges[0].Points(v){
		t.Fail()
	}
}
