package groph

import "errors"

func (g *Graph) Outer(v *Vertex) ([]*Edge, error) {
	return v.Outer()
}

//OuterWhereEdge returns edges that matches specified filter
func (g *Graph) OuterWhereEdge(filterFunc func(*Edge) bool) (edges []*Edge, err error) {
	edges = make([]*Edge, 0)

	for _, e := range g.StartVertex.OuterEdges {
		if filterFunc(e) {
			edges = append(edges, e)
		}
	}

	if len(edges) == 0 {
		err = errors.New("No edges found")
	}

	return
}

//OuterWhereVertex returns edges that passes the provided filter
func (g *Graph) OuterWhereVertex(f func(*Vertex) bool) (edges []*Edge, err error) {
	edges = make([]*Edge, 0)

	g.Traverse(func(v *Vertex) {
		if f(v) {
			edges = append(edges, v.OuterEdges...)
		}
	})

	if len(edges) == 0 {
		err = noEdgesError
	}

	return
}
