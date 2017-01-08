package groph

import "testing"

func TestTraverse(t *testing.T) {
	graph := getMockedGraph()

	res := make([]*Vertex, 0)
	graph.Traverse(func(v *Vertex) {
		if v.GetID() != "start" && v.GetID() != "finish" {
			res = append(res, v)
		}
	})

	if len(res) != 2 {
		t.Errorf("Unexpected number of elements %d != %d\n", len(res), 2)
	}

	for _, v := range res {
		if v.GetID() != "A" && v.GetID() != "B" {
			t.Fail()
		}
	}
}
