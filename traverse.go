package groph

func (g *Graph) Traverse(f func(*Vertex)) {
	g.traverse(g.StartVertex, f, make(map[*Vertex]bool))
}

func (g *Graph) traverse(s *Vertex, f func(*Vertex), seen map[*Vertex]bool) {
	if seen[s] {
		return
	}

	f(s)
	seen[s] = true

	for _, v := range s.OuterEdges {
		g.traverse(v.PointsTo, f, seen)
	}

	for _, v := range s.InnerEdges {
		g.traverse(v.From, f, seen)
	}
}
