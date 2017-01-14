package groph

func (g *Graph) BreadthFirst(id interface{}, f func(*Vertex)) {
	g.breadthFirst(g.StartVertex, id, f, make(map[*Vertex]bool))
}

func (g *Graph) breadthFirst(s *Vertex, id interface{}, f func(*Vertex), seen map[*Vertex]bool) {
	if seen[s] {
		return
	}

	seen[s] = true

	if s.GetID() == id {
		f(s)
		return
	} else {
		for _, v := range s.outEdges {
			g.breadthFirst(v.PointsTo, id, f, seen)
		}

		for _,v := range s.inEdges {
			g.breadthFirst(v.From, id, f, seen)
		}
	}
}
