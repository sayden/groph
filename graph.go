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

	s.OuterEdges = append(s.OuterEdges, e)
	t.InnerEdges = append(t.InnerEdges, e)

	if g.StartVertex == nil {
		g.StartVertex = s
	}
}

// NewVertexWithUpdate returns an initialized vertex with the provided data if it doesn't exists already or a pointer to the
// already existing one with the contents updated with the incoming data
func (g *Graph) NewVertexWithUpdate(d Data) *Vertex {
	if g.IndexMap[d.GetID()] == nil {
		newV := &Vertex{
			InnerEdges: make([]*Edge, 0),
			OuterEdges: make([]*Edge, 0),
			Data:       d,
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
			InnerEdges: make([]*Edge, 0),
			OuterEdges: make([]*Edge, 0),
			Data:       d,
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

//TODO SaveToDisk
func (g *Graph) SaveToDisk(filePath string) error {
	return errors.New("Not implemented")
}

// NewGraph just returns an initialized graph
func NewGraph() *Graph {
	return &Graph{
		IndexMap: make(map[interface{}]*Vertex),
	}
}

