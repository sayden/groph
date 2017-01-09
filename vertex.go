package groph

import "fmt"

type Vertex struct {
	Data `json:"data, omitempty"`
	OuterEdges []*Edge `json:"outer_edges, omitempty"`
	InnerEdges []*Edge `json:"inner_edges, omitempty"`
}

type Data interface {
	GetData() interface{}
	GetID() interface{}
}

// Outer returns the edges that the current vertex has
func (v *Vertex) Outer() ([]*Edge, error) {
	if len(v.OuterEdges) == 0 {
		return nil, noEdgesError
	}

	return v.OuterEdges, nil
}

func (v *Vertex) Inner() ([]*Edge, error) {
	if len(v.InnerEdges) == 0 {
		return nil, noEdgesError
	}

	return v.InnerEdges, nil
}

func (v *Vertex) String() string {
	res := fmt.Sprintf("VERTEX ID: '%s'\nVertex Data: %v\nEdges pointing this vertex:\n[START]", v.GetID(), v.GetData())
	for _, edge := range v.InnerEdges {
		res += edge.String()
	}

	res += "\n[END]\nEdges pointing out:\n[START]"
	for _, edge := range v.OuterEdges {
		res += edge.String()
	}

	res += "[END]"

	return res
}
