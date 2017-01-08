package groph

import (
	"errors"
	"os"
)

type Graph struct {
	StartVertex *Vertex
	IndexMap    map[interface{}]*Vertex
}

func (g *Graph) Find(id interface{}) (*Vertex, error){
	if g.IndexMap[id] != nil {
		return g.IndexMap[id], nil
	}

	return nil, vertexNotFoundError
}

func (g *Graph) AddConnection(s, t *Vertex, e *Edge) {
	e.From = s
	e.PointsTo = t

	s.OuterEdges = append(s.OuterEdges, e)
	t.InnerEdges = append(t.InnerEdges, e)
}

func (g *Graph) NewVertex(d VertexData) *Vertex {
	newV := &Vertex{
		InnerEdges: make([]*Edge, 0),
		OuterEdges: make([]*Edge, 0),
		VertexData: d,
	}

	g.IndexMap[d.GetID()] = newV

	return newV
}

func (g *Graph) SetRootVertex(r *Vertex) {
	g.StartVertex = r
}

func (g *Graph) SaveToDisk(filePath string) error {
	return errors.New("Not implemented")
}

func LoadGraphFromDisk(filePath string) (*Graph, error) {
	_, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return nil, errors.New("Not implemented")

}

func NewGraph() *Graph {
	return &Graph{
		IndexMap: make(map[interface{}]*Vertex),
	}
}
