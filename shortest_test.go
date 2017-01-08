package groph

import "testing"

func TestShortest(t *testing.T) {
	graph := getMockedGraph()

	finishVertex := graph.StartVertex.OuterEdges[0].PointsTo.OuterEdges[0].PointsTo
	chain, totalCost, err := graph.ShortestPathWithVertex(finishVertex)

	if chain[0] != graph.StartVertex {
		t.Fatal(chain[0])
	}

	if chain[1] != graph.StartVertex.OuterEdges[1].PointsTo {
		t.Fatal(chain[1])
	}


	if chain[2] != graph.StartVertex.OuterEdges[0].PointsTo {
		t.Fatal(chain[2])
	}


	if chain[3] != finishVertex {
		t.Fatal(chain[3])
	}


	if err != nil {
		t.Fatal()
	}

	if totalCost != 6.0 {
		t.Errorf("%f != %f", totalCost, 6.0)
	}

	_, totalCost, err = graph.ShortestPath("B")

	if err != nil {
		t.Fatal()
	}

	if totalCost != 2.0 {
		t.Errorf("%f != %f", totalCost, 6.0)
	}
}
