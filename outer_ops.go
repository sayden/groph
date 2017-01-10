package groph

func (g *Graph) Outer(v *Vertex) Edges {
	return v.Outer()
}

//OuterWhereEdge returns edges that matches specified filter
func (g *Graph) OuterWhereEdge(filterFunc func(*Edge) bool) (edges Edges) {
	edges = make(Edges, 0)

	for _, e := range g.StartVertex.OuterEdges {
		if filterFunc(e) {
			edges = append(edges, e)
		}
	}

	return
}

//OuterWhereVertex returns edges that passes the provided filter
func (g *Graph) OuterWhereVertex(f func(*Vertex) bool) (edges Edges) {
	edges = make(Edges, 0)

	g.Traverse(func(v *Vertex) {
		if f(v) {
			edges = append(edges, v.OuterEdges...)
		}
	})

	return
}
