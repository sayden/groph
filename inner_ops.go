package groph

//Inner returns the edges that points to current vertex
func (g *Graph) Inner(v *Vertex) []*Edge {
	return v.Inner()
}

//InnerWhereEdge traverses the entire graph and returns all inner edges that passes the filter function
func (g *Graph) InnerWhereEdge(f func(*Edge) bool) (res []*Edge) {
	res = make([]*Edge, 0)

	g.Traverse(func(v *Vertex) {
		for _, edge := range v.InnerEdges {
			if f(edge) {
				res = append(res, edge)
			}
		}
	})

	return res
}


//InnerWhereVertex returns all inner edges that matches the filter function
func (g *Graph) InnerWhereVertex(f func(*Vertex) bool) (res []*Edge) {
	res = make([]*Edge, 0)

	g.Traverse(func(c *Vertex) {
		if f(c) {
			res = append(res, c.InnerEdges...)
		}
	})

	return
}
