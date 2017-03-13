package groph

import "errors"

// Graph is the main type to do queries and stores a vertex to be used as the starting point of any query. At the same
// time it maintains a map of known vertices
type Graph struct {
	StartVertex *Vertex
	IndexMap    map[interface{}]*Vertex
}

// Find returns the vertex with the provided ID or a not found error
func (g *Graph) Find(id interface{}) (*Vertex, error) {
	if g.IndexMap[id] != nil {
		return g.IndexMap[id], nil
	}

	return nil, vertexNotFoundError
}

func (g *Graph) NewEdge(d Data, weight float64) *Edge {
	return &Edge{
		Data:   d,
		Weight: weight,
	}
}

// AddConnection connects two existing vertices with the provided edge
func (g *Graph) AddConnection(s, t *Vertex, e *Edge) {
	e.From = s
	e.PointsTo = t

	s.outEdges = append(s.outEdges, e)
	t.inEdges = append(t.inEdges, e)

	if g.StartVertex == nil {
		g.StartVertex = s
	}
}

// NewVertexWithUpdate returns an initialized vertex with the provided data if it doesn't exists already or a pointer to the
// already existing one with the contents updated with the incoming data
func (g *Graph) NewVertexWithUpdate(d Data) *Vertex {
	if g.IndexMap[d.GetID()] == nil {
		newV := &Vertex{
			inEdges:  make([]*Edge, 0),
			outEdges: make([]*Edge, 0),
			Data:     d,
		}

		g.IndexMap[d.GetID()] = newV

		return newV
	}

	g.IndexMap[d.GetID()].Data = d

	return g.IndexMap[d.GetID()]
}

// NewVertex returns an initialized vertex with the provided data if it doesn't exists already or a pointer to the
// already existing one. Any contents incoming in 'd' are lost if the Vertex is found. If you want to update use
// 'NewVertexWithUpdate' instead.
func (g *Graph) NewVertex(d Data) *Vertex {
	if g.IndexMap[d.GetID()] == nil {
		newV := &Vertex{
			inEdges:  make([]*Edge, 0),
			outEdges: make([]*Edge, 0),
			Data:     d,
		}

		g.IndexMap[d.GetID()] = newV

		return newV
	}

	return g.IndexMap[d.GetID()]
}

// SetRootVertex changes the current root vertex. This is useful to initiate some specific searches from a particular vertex
func (g *Graph) SetRootVertex(r *Vertex) {
	g.StartVertex = r
}

func (g *Graph) RootVertex() (v *Vertex, err error) {
	if g.StartVertex != nil {
		v = g.StartVertex
		return
	}

	err = errors.New("Root vertex not set")
	return
}

func (g *Graph) Out(v *Vertex) Edges {
	return v.OutEdges()
}

//OutWhereEdge returns edges that matches specified filter
func (g *Graph) OutWhereEdge(filterFunc func(*Edge) bool) (edges Edges) {
	edges = make(Edges, 0)

	for _, e := range g.StartVertex.outEdges {
		if filterFunc(e) {
			edges = append(edges, e)
		}
	}

	return
}

//OutWhereVertex returns edges that passes the provided filter
func (g *Graph) OutWhereVertex(f func(*Vertex) bool) (edges Edges) {
	edges = make(Edges, 0)

	g.Traverse(func(v *Vertex) {
		if f(v) {
			edges = append(edges, v.outEdges...)
		}
	})

	return
}

//In returns the edges that points to current vertex
func (g *Graph) In(v *Vertex) []*Edge {
	return v.InEdges()
}

//InWhereEdge traverses the entire graph and returns all inner edges that passes the filter function
func (g *Graph) InWhereEdge(f func(*Edge) bool) (res []*Edge) {
	res = make([]*Edge, 0)

	g.Traverse(func(v *Vertex) {
		for _, edge := range v.inEdges {
			if f(edge) {
				res = append(res, edge)
			}
		}
	})

	return res
}

func (g *Graph) Filter(f func(*Vertex) bool) (vs Vertices) {
	vs = make([]*Vertex, 0)

	for _, v := range g.IndexMap {
		if f(v) {
			vs = append(vs, v)
		}
	}

	return
}

//InWhereVertex returns all inner edges that matches the filter function
func (g *Graph) InWhereVertex(f func(*Vertex) bool) (res []*Edge) {
	res = make([]*Edge, 0)

	g.Traverse(func(c *Vertex) {
		if f(c) {
			res = append(res, c.inEdges...)
		}
	})

	return
}

//All is a map-like function for each vertex in a graph
func (g *Graph) All(f func(*Vertex)) {
	for _, v := range g.IndexMap {
		f(v)
	}

	return
}

//TODO SaveToDisk
func (g *Graph) SaveToDisk(filePath string) error {
	return errors.New("Not implemented")
}

// New just returns an initialized graph
func New() *Graph {
	return &Graph{
		IndexMap: make(map[interface{}]*Vertex),
	}
}
